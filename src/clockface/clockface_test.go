package clockface

import (
	"math"
	"testing"
	"time"
)

// 0時0分0秒の時の秒針の位置をテスト
func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

// 0時0分30秒の時の秒針の位置をテスト
func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	expected := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)

	if got != expected {
		t.Errorf("Got %v, expected %v", got, expected)
	}
}

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}
