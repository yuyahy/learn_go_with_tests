package sync

import (
	"sync"
	"testing"
)

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d,want %d", got.Value(), want)
	}
}

func TestCounter(t *testing.T) {
	t.Run("increamenting the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 100000
		counter := NewCounter()

		var wg sync.WaitGroup
		// wg.Add()でカウンタを設定する
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				// wg.Done()でカウンタを1つ減らす
				wg.Done()
			}()
		}
		// wg.Add()したカウンタが0になるまで待つ
		wg.Wait()
		assertCounter(t, counter, wantedCount)
	})
}
