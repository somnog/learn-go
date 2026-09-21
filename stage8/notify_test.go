package stage8_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/somnog/learn-go/stage8"
)

// --- NotifyOne ---

func TestNotifyOne(t *testing.T) {
	result := stage8.NotifyOne("abdi@somnog.org", "Your session starts soon")
	if !strings.Contains(result, "abdi@somnog.org") {
		t.Errorf("NotifyOne result missing email: %q", result)
	}
	if !strings.HasPrefix(result, "sent:") {
		t.Errorf("NotifyOne result should start with 'sent:': %q", result)
	}
}

// --- NotifyAll ---

func TestNotifyAll_Length(t *testing.T) {
	emails := []string{
		"layla@somnog.org",
		"hassan@somnog.org",
		"nimo@somnog.org",
	}
	results := stage8.NotifyAll(emails, "SomNOG9 starts tomorrow")
	if len(results) != len(emails) {
		t.Errorf("NotifyAll returned %d results, want %d", len(results), len(emails))
	}
}

func TestNotifyAll_AllPresent(t *testing.T) {
	emails := []string{
		"omar@somnog.org",
		"hodan@somnog.org",
	}
	results := stage8.NotifyAll(emails, "See you at SomNOG9")
	for _, email := range emails {
		found := false
		for _, r := range results {
			if strings.Contains(r, email) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("NotifyAll result missing entry for %s", email)
		}
	}
}

func TestNotifyAll_OrderIndependent(t *testing.T) {
	emails := []string{
		"deeqa@somnog.org",
		"khalid@somnog.org",
		"ayan@somnog.org",
	}
	results := stage8.NotifyAll(emails, "Welcome")
	sort.Strings(results)
	if len(results) != 3 {
		t.Errorf("expected 3 results, got %d", len(results))
	}
}

// --- FetchAll ---

func TestFetchAll_Length(t *testing.T) {
	urls := []string{
		"https://somnog.org/schedule",
		"https://somnog.org/speakers",
		"https://somnog.org/venue",
	}
	results := stage8.FetchAll(urls)
	if len(results) != len(urls) {
		t.Errorf("FetchAll returned %d results, want %d", len(results), len(urls))
	}
}

func TestFetchAll_AllPresent(t *testing.T) {
	urls := []string{
		"https://somnog.org/registration",
		"https://somnog.org/sponsors",
	}
	results := stage8.FetchAll(urls)
	for _, url := range urls {
		found := false
		for _, r := range results {
			if strings.Contains(r, url) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("FetchAll result missing entry for %s", url)
		}
	}
}

func TestFetchAll_HasFetchedPrefix(t *testing.T) {
	urls := []string{"https://somnog.org/about"}
	results := stage8.FetchAll(urls)
	if len(results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(results))
	}
	if !strings.HasPrefix(results[0], "fetched:") {
		t.Errorf("FetchAll result should start with 'fetched:': %q", results[0])
	}
}
