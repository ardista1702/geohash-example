package main

import (
	"fmt"
	"geohash-example/geohash_"

	"github.com/mmcloughlin/geohash"
)

func main() {
	gh := geohash_.NewGeoHash(-6.1753924, 106.8271528, 12)
	res := geohash.Encode(-6.1753924, 106.8271528)
	fmt.Println(gh.Encode())
	fmt.Println(res)
}
