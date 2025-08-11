package main

import (
	"fmt"
	"geohash-example/geohash_"
)

func main() {
	latitude := -6.2088
	longitude := 106.8456
	precision := uint8(12)

	gh := geohash_.NewGeoHash(latitude, longitude, precision)
	hash := gh.Encode()

	fmt.Printf("Geohash Encode for lat=%.4f, long=%.4f, precision=%d:\n%s\n",
		latitude, longitude, precision, hash)
}
