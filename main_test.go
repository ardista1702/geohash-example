package main

import (
	"testing"

	"geohash-example/geohash_"

	geohash "github.com/mmcloughlin/geohash"
)

func TestEncodeMultiplePrecisions(t *testing.T) {
	type testCase struct {
		name      string
		latitude  float64
		longitude float64
	}

	testPoints := []testCase{
		{name: "San Francisco", latitude: 37.7749, longitude: -122.4194},
		{name: "Jakarta", latitude: -6.2088, longitude: 106.8456},
		{name: "Tokyo", latitude: 35.6895, longitude: 139.6917},
	}

	for _, point := range testPoints {
		for precision := uint8(1); precision <= 12; precision++ {
			t.Run(point.name+" precision "+string(precision+'0'), func(t *testing.T) {
				gh := geohash_.NewGeoHash(point.latitude, point.longitude, precision)
				got := gh.Encode()
				want := geohash.EncodeWithPrecision(point.latitude, point.longitude, uint(precision))
				if got != want {
					t.Errorf("Encode() = %v, want %v", got, want)
				}
			})
		}
	}
}
