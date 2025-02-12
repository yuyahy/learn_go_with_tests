package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	testStr := "b"
	repeatCnt := 10
	got := Repeat(testStr, repeatCnt)
	expected := strings.Repeat(testStr, repeatCnt)

	if got != expected {
		t.Errorf("expected '%q' but got '%q'", expected, got)
	}
}

// b.Nはtestingライブラリが自動的に決定してくれる
func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	str := Repeat("str", 5)
	fmt.Println(str)
	// Output: strstrstrstrstr
}
