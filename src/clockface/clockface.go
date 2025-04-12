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
	// return float64(t.Second()) * (math.Pi / 30) // 掛け算割り算を行うと浮動小数点の誤差が蓄積するため、割り算のみで計算する
	return (math.Pi / (30 / (float64(t.Second()))))
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func secondHandPoint(t time.Time) Point {
	return Point{0, -1}
}
