package geohash_

import "strings"

const base32Chars = "0123456789bcdefghjkmnpqrstuvwxyz"

type GeoHash_ interface {
}

type geohash struct {
	Latitude   float64
	Longitude  float64
	Pressision uint8
}

type Coordinate int

const (
	Longitude Coordinate = iota
	Latitude
)

func NewGeoHash(latitude, longitude float64, pressision uint8) *geohash {
	return &geohash{
		Latitude:   latitude,
		Longitude:  longitude,
		Pressision: pressision,
	}
}

func (gh *geohash) Encode() string {
	totalBits := int(gh.Pressision * 5)

	lonBits := totalBits / 2
	latBits := totalBits / 2
	if totalBits%2 == 1 {
		lonBits++
	}

	latitude := gh.convertToBinary(Latitude, uint8(latBits))
	longitude := gh.convertToBinary(Longitude, uint8(lonBits))

	interleaved := gh.interleaved(longitude, latitude)
	chunck := gh.chunck(interleaved)
	decimals := convertBinaryToDecimal(chunck)
	return gh.hash(decimals)
}
func (gh *geohash) convertToBinary(coordinate Coordinate, numBits uint8) []byte {
	DefaultLatitude, DefaultLongitude := []float64{-90.0, 90.0}, []float64{-180.0, 180.0}

	var bounds []float64
	var value float64

	switch coordinate {
	case Longitude:
		bounds = DefaultLongitude
		value = gh.Longitude
	case Latitude:
		bounds = DefaultLatitude
		value = gh.Latitude
	default:
		// bisa handle error atau panic kalau perlu
	}

	left, right := bounds[0], bounds[1]

	result := make([]byte, numBits)
	for i := uint8(0); i < numBits; i++ {
		mid := (left + right) / 2
		if value > mid {
			result[i] = 1
			left = mid
		} else {
			result[i] = 0
			right = mid
		}
	}
	return result
}
func (gh *geohash) interleaved(longBins, latBins []byte) []byte {
	i, j := 0, 0
	var result []byte
	for i < len(longBins) && j < len(latBins) {
		result = append(result, longBins[j])
		result = append(result, latBins[j])
		i++
		j++
	}
	return result
}
func (gh *geohash) chunck(interleaved []byte) [][5]byte {
	numChunks := (len(interleaved) + 4) / 5
	res := make([][5]byte, numChunks)

	for i := 0; i < numChunks*5; i++ {
		var bit byte = 0
		if i < len(interleaved) {
			bit = interleaved[i]
		}
		res[i/5][i%5] = bit

	}
	return res
}

func (gh *geohash) hash(decimals []byte) string {
	var builder strings.Builder
	builder.Grow(len(decimals))
	for _, dec := range decimals {
		builder.WriteByte(base32Chars[dec])
	}
	return builder.String()
}

func convertBinaryToDecimal(binaries [][5]byte) []byte {
	var res []byte
	for _, bin := range binaries {
		var dec byte = 0
		for j, bit := range bin {
			if bit == 1 {
				dec |= 1 << (4 - j)
			}
		}
		res = append(res, dec)

	}
	return res
}
