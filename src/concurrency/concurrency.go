package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func() {
			// チャネルへ非同期に結果を送信
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(urls); i++ {
		// チャネルから値を一個ずつ受け取る
		// ※同時アクセスを防ぐために、ここは同期的に処理する必要がある
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
