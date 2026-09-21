package stage7

import (
	"fmt"
	"strings"
)

// Describable is implemented by any value that can describe itself.
type Describable interface {
	Describe() string
}

// Speaker represents a session speaker.
type Speaker struct {
	Name  string
	Title string
}

// Describe satisfies the Describable interface for Speaker.
func (sp Speaker) Describe() string {
	return fmt.Sprintf("Speaker: %s (%s)", sp.Name, sp.Title)
}

// Session represents a conference session.
type Session struct {
	ID    int
	Title string
}

// Describe satisfies the Describable interface for Session.
func (s Session) Describe() string {
	return fmt.Sprintf("Session #%d: %s", s.ID, s.Title)
}

// Workshop represents a hands-on workshop.
type Workshop struct {
	ID              int
	Title           string
	Instructor      string
	DurationMinutes int
}

// Describe satisfies the Describable interface for Workshop.
func (w Workshop) Describe() string {
	return fmt.Sprintf("Workshop: %s (%d min) — %s", w.Title, w.DurationMinutes, w.Instructor)
}

// DescribeAll returns a Describe() string for every item in the slice.
func DescribeAll(items []Describable) []string {
	results := make([]string, len(items))
	for i, item := range items {
		results[i] = item.Describe()
	}
	return results
}

// FindByTitle returns the first Describable whose Describe() output contains
// title (case-insensitive). Returns nil and false if no match is found.
func FindByTitle(items []Describable, title string) (Describable, bool) {
	lower := strings.ToLower(title)
	for _, item := range items {
		if strings.Contains(strings.ToLower(item.Describe()), lower) {
			return item, true
		}
	}
	return nil, false
}
