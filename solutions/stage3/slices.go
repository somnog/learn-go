package stage3

// Tracks returns the four SomNOG9 track names in order.
func Tracks() []string {
	return []string{
		"Network Infrastructure",
		"System & Services",
		"Software Development",
		"Cybersecurity",
	}
}

// Add returns a new slice with item appended. The original slice is not modified.
func Add(items []string, item string) []string {
	result := make([]string, len(items), len(items)+1)
	copy(result, items)
	return append(result, item)
}

// First returns the first n items from items.
// If n <= 0, an empty slice is returned.
// If n >= len(items), all items are returned.
func First(items []string, n int) []string {
	if n <= 0 {
		return []string{}
	}
	if n >= len(items) {
		return items
	}
	return items[:n]
}

// Contains reports whether target appears in items (case-sensitive).
func Contains(items []string, target string) bool {
	for _, v := range items {
		if v == target {
			return true
		}
	}
	return false
}

// MakeSchedule returns a slice of n empty strings — a blank schedule grid.
func MakeSchedule(n int) []string {
	return make([]string, n)
}

// Reverse returns a new slice with the elements in reverse order.
// The original slice is not modified.
func Reverse(items []string) []string {
	result := make([]string, len(items))
	for i, v := range items {
		result[len(items)-1-i] = v
	}
	return result
}
