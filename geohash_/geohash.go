package geohash_

import "strings"

type GeoHash_ interface {
}

type geohash struct {
	Latitude   float64
	Longitude  float64
	Pressision uint8
}

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

    latitude := gh.convertToBinary("latitude", uint8(latBits))
    longitude := gh.convertToBinary("longitude", uint8(lonBits))

    interleaved := gh.interleaved(longitude, latitude)
    chunck := gh.chunck(interleaved)
    decimals := convertBinaryToDecimal(chunck)
    return gh.hash(decimals)
}
func (gh *geohash) convertToBinary(coordinate string, numBits uint8) []byte {
    DefaultLatitude, DefaultLongitude := []float64{-90.0, 90.0}, []float64{-180.0, 180.0}

    var coordinates []float64
    var value float64
    
    result := make([]byte, numBits)

    switch coordinate {
    case "longitude":
        coordinates = DefaultLongitude
        value = gh.Longitude
    case "latitude":
        coordinates = DefaultLatitude
        value = gh.Latitude
    }

    left, right := coordinates[0], coordinates[1]
    
    // Loop sebanyak numBits, bukan gh.Pressision
    for length := uint8(0); length < numBits; length++ {
        mid := (left + right) / 2
        if value > mid {
            result[length] = 1
            left = mid
        } else {
            result[length] = 0
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
	var res [][5]byte
	for i := 0; i < len(interleaved); i += 5 {
		var currentChunk [5]byte
		for j := 0; j < 5; j++ {
			if i+j < len(interleaved) {
				currentChunk[j] = interleaved[i+j]
			} else {
				currentChunk[j] = 0 
			}
		}
		res = append(res, currentChunk)
	}
	return res
}

func (gh *geohash) hash(decimals []byte) string {
base32Table := map[uint8]string{
    0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7",
    8: "8", 9: "9", 10: "b", 11: "c", 12: "d", 13: "e", 14: "f", 15: "g",
    16: "h", 17: "j", 18: "k", 19: "m", 20: "n", 21: "p", 22: "q", 23: "r",
    24: "s", 25: "t", 26: "u", 27: "v", 28: "w", 29: "x", 30: "y", 31: "z",
}	
	var res []string
	for _, dec := range decimals {
		res = append(res, base32Table[dec])
	}
	return strings.Join(res, "")
}

func convertBinaryToDecimal(binaries [][5]byte) []byte {
	var res []byte
	for _, bin := range binaries {
		var dec byte = 0
		for j := range bin {
			if bin[j] == 1 {
				dec += pow(2, byte(4-j))
			}
		}
		res = append(res, dec)

	}
	return res
}

func pow(base, exp byte) byte {
	if exp == 0 {
		return 1
	}
	if exp%2 == 0 {
		half := pow(base, exp/2)
		return half * half
	}
	return base * pow(base,exp - 1)
}
