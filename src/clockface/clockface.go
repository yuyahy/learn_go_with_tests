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

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

// 時刻から秒針の先端の座標を返す
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // 針の長さにスケール
	p = Point{p.X, -p.Y}                                      // X軸に対して反転
	p = Point{p.X + clockCentreX, p.Y + clockCentreY}         // 位置調整(原点は(150, 150))
	return p
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
	angle := secondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
