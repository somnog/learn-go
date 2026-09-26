package stage2

// Implement these four functions:

// - `TrackName(n int) string` — return the track name for 1–4, `"Unknown"` for anything else. Use `switch`.

// - 1 → `"Network Infrastructure"`
// - 2 → `"System & Services"`
// - 3 → `"Software Development"`
// - 4 → `"Cybersecurity"`

// - `CanRegister(enrolled, capacity int) bool` — return `true` if `enrolled < capacity`. Use `if`.

// - `CountAvailable(enrolled []int, capacity int) int` — given a slice of enrolled counts and a shared capacity, return the number of sessions that still have open spots. Use `for`.

// - `Greet(hour int) string` — return a greeting string that contains **both** a time-of-day word **and** `"SomNOG9"`:
//   - `hour < 12` → contains `"morning"` and `"SomNOG9"` — e.g. `"Good morning at SomNOG9"`
//   - `hour >= 12` and `hour < 17` → contains `"afternoon"` and `"SomNOG9"`
//   - `hour >= 17` → contains `"evening"` and `"SomNOG9"`

func TrackName(number int) string {
	switch number {
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

func CanRegister(enrolled, capacity int) bool {
	return enrolled < capacity
}

func CountAvailable(enrolled []int, capacity int) int {

}
