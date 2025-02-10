package main

import "testing"

func TestHello(t *testing.T) {
	// testing.TBはtesting.Tとtesting.Bの両方を受け入れる
	assertCorrectMessage := func(t testing.TB, got, want string) {
		// 本メソッドがヘルパーであることをテストスイートに知らせる
		// →これにより、本メソッドを呼ぶテストが失敗した場合は、テストヘルパーの中ではなく、
		//	呼び出しもとの行番号が出力される
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string defaults to World", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
}
