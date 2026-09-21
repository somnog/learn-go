package stage5

import "fmt"

// Speaker holds information about a session's presenter.
type Speaker struct {
	Name  string
	Title string
}

// Session represents a single conference session.
type Session struct {
	ID       int
	Title    string
	Speaker  Speaker
	Track    string
	Capacity int
	Enrolled int
}

// NewSession creates a Session with the given id, title, track, and capacity.
// Speaker is left at its zero value and can be set separately.
// Enrolled starts at 0.
func NewSession(id int, title, track string, capacity int) Session {
	return Session{
		ID:       id,
		Title:    title,
		Track:    track,
		Capacity: capacity,
	}
}

// IsFull reports whether the session has reached or exceeded its capacity.
func (s Session) IsFull() bool {
	return s.Enrolled >= s.Capacity
}

// SpotsLeft returns how many seats are still available.
// Never returns a negative number.
func (s Session) SpotsLeft() int {
	spots := s.Capacity - s.Enrolled
	if spots < 0 {
		return 0
	}
	return spots
}

// Summary returns a one-line human-readable summary of the session.
func (s Session) Summary() string {
	return fmt.Sprintf("[%s] %s | Speaker: %s | Spots left: %d",
		s.Track, s.Title, s.Speaker.Name, s.SpotsLeft())
}
