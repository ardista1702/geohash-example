package main

import (
	"fmt"
	"geohash-example/geohash_"
)

func main() {
	gh := geohash_.NewGeoHash(-6.1753924, 106.8271528, 12)
	res := gh.Encode()
	fmt.Println(res)
}
