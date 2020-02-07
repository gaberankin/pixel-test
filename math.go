package main

import (
	"math"

	"github.com/faiface/pixel"
)

func pointAngle(a, b pixel.Vec) float64 {
	deg := math.Atan2(b.Y-a.Y, b.X-a.X) * 180 / math.Pi

	if deg < 0 {
		return deg + 360.0
	}
	return deg
	// return (result < 0) ? (360d + result) : result;
}
