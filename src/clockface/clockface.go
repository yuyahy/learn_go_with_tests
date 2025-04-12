package clockface

import (
	"math"
	"time"
)

// Point is a 2D Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// 時刻から秒針の先端の座標を返す
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi
}
