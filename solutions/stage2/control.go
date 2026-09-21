package stage2

import "fmt"

// TrackName returns the name of the SomNOG9 track by number (1–4).
func TrackName(n int) string {
	switch n {
	case 1:
		return "Network Infrastructure"
	case 2:
		return "System & Services"
	case 3:
		return "Software Development"
	case 4:
		return "Cybersecurity"
	default:
		return "Unknown"
	}
}

// CanRegister reports whether a session has room for more attendees.
func CanRegister(enrolled, capacity int) bool {
	return enrolled < capacity
}

// CountAvailable counts how many sessions still have open spots.
func CountAvailable(enrolled []int, capacity int) int {
	count := 0
	for _, e := range enrolled {
		if e < capacity {
			count++
		}
	}
	return count
}

// Greet returns a time-of-day greeting that also mentions SomNOG9.
// hour 0–11  → morning
// hour 12–16 → afternoon
// hour 17+   → evening
func Greet(hour int) string {
	var timeWord string
	if hour < 12 {
		timeWord = "morning"
	} else if hour < 17 {
		timeWord = "afternoon"
	} else {
		timeWord = "evening"
	}
	return fmt.Sprintf("Good %s at SomNOG9!", timeWord)
}
