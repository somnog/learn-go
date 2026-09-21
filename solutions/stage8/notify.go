package stage8

import "sync"

// NotifyOne simulates sending a notification email.
// It returns a confirmation string of the form "sent: <email>".
func NotifyOne(email, message string) string {
	// In a real system this would call an SMTP server.
	_ = message
	return "sent: " + email
}

// NotifyAll sends a notification to every email in the list concurrently
// using goroutines and a channel, and returns all confirmation strings.
func NotifyAll(emails []string, message string) []string {
	ch := make(chan string, len(emails))
	var wg sync.WaitGroup

	for _, email := range emails {
		wg.Add(1)
		go func(e string) {
			defer wg.Done()
			ch <- NotifyOne(e, message)
		}(email)
	}

	wg.Wait()
	close(ch)

	results := make([]string, 0, len(emails))
	for r := range ch {
		results = append(results, r)
	}
	return results
}

// FetchAll simulates fetching multiple URLs concurrently.
// It returns a "fetched: <url>" string for each URL.
func FetchAll(urls []string) []string {
	ch := make(chan string, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			ch <- "fetched: " + u
		}(url)
	}

	wg.Wait()
	close(ch)

	results := make([]string, 0, len(urls))
	for r := range ch {
		results = append(results, r)
	}
	return results
}
