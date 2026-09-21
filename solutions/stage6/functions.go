package stage6

import (
	"errors"
	"strings"
)

// Speaker represents a session speaker.
type Speaker struct {
	Name  string
	Title string
}

// Session represents a conference session.
// Registered holds the email addresses of enrolled attendees.
type Session struct {
	ID         int
	Title      string
	Speaker    Speaker
	Track      string
	Capacity   int
	Registered []string
}

// Enroll adds an attendee's email to the session.
// Returns an error if the session is full or the attendee is already registered.
func Enroll(s *Session, email string) error {
	if len(s.Registered) >= s.Capacity {
		return errors.New("session is full")
	}
	for _, e := range s.Registered {
		if e == email {
			return errors.New("already registered")
		}
	}
	s.Registered = append(s.Registered, email)
	return nil
}

// IsEnrolled reports whether an attendee is enrolled in the session.
func IsEnrolled(s *Session, email string) bool {
	for _, e := range s.Registered {
		if e == email {
			return true
		}
	}
	return false
}

// FindSession searches sessions for a matching ID.
// Returns a pointer to the session and nil error if found.
// Returns nil and an error if not found.
func FindSession(sessions []*Session, id int) (*Session, error) {
	for _, s := range sessions {
		if s.ID == id {
			return s, nil
		}
	}
	return nil, errors.New("session not found")
}

// ValidateEmail returns an error if email is invalid.
// Rules: must be non-empty, contain '@', and contain no whitespace.
func ValidateEmail(email string) error {
	if email == "" {
		return errors.New("email must not be empty")
	}
	if !strings.Contains(email, "@") {
		return errors.New("email must contain @")
	}
	if strings.ContainsAny(email, " \t\n\r") {
		return errors.New("email must not contain whitespace")
	}
	return nil
}
