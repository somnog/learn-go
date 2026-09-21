package stage6_test

import (
	"testing"

	"github.com/somnog/learn-go/stage6"
)

// --- helpers ---

func makeSession(id, capacity int) *stage6.Session {
	return &stage6.Session{
		ID:       id,
		Title:    "Test Session",
		Capacity: capacity,
	}
}

// --- Enroll ---

func TestEnroll_Success(t *testing.T) {
	s := makeSession(1, 3)
	err := stage6.Enroll(s, "abdi@somnog.org")
	if err != nil {
		t.Fatalf("Enroll returned unexpected error: %v", err)
	}
	if len(s.Registered) != 1 {
		t.Errorf("expected 1 registrant, got %d", len(s.Registered))
	}
	if s.Registered[0] != "abdi@somnog.org" {
		t.Errorf("expected abdi@somnog.org, got %s", s.Registered[0])
	}
}

func TestEnroll_Full(t *testing.T) {
	s := makeSession(2, 1)
	_ = stage6.Enroll(s, "layla@somnog.org")
	err := stage6.Enroll(s, "hassan@somnog.org")
	if err == nil {
		t.Fatal("expected an error when session is full, got nil")
	}
}

func TestEnroll_Duplicate(t *testing.T) {
	s := makeSession(3, 5)
	_ = stage6.Enroll(s, "nimo@somnog.org")
	err := stage6.Enroll(s, "nimo@somnog.org")
	if err == nil {
		t.Fatal("expected an error for duplicate registration, got nil")
	}
}

// --- IsEnrolled ---

func TestIsEnrolled_True(t *testing.T) {
	s := makeSession(4, 5)
	_ = stage6.Enroll(s, "omar@somnog.org")
	if !stage6.IsEnrolled(s, "omar@somnog.org") {
		t.Error("expected IsEnrolled to return true for registered attendee")
	}
}

func TestIsEnrolled_False(t *testing.T) {
	s := makeSession(5, 5)
	if stage6.IsEnrolled(s, "hodan@somnog.org") {
		t.Error("expected IsEnrolled to return false for unregistered attendee")
	}
}

// --- FindSession ---

func TestFindSession_Found(t *testing.T) {
	sessions := []*stage6.Session{
		makeSession(10, 5),
		makeSession(20, 5),
		makeSession(30, 5),
	}
	got, err := stage6.FindSession(sessions, 20)
	if err != nil {
		t.Fatalf("FindSession returned unexpected error: %v", err)
	}
	if got == nil {
		t.Fatal("FindSession returned nil, expected a valid pointer")
	}
	if got.ID != 20 {
		t.Errorf("expected ID 20, got %d", got.ID)
	}
}

func TestFindSession_NotFound(t *testing.T) {
	sessions := []*stage6.Session{
		makeSession(10, 5),
	}
	got, err := stage6.FindSession(sessions, 999)
	if err == nil {
		t.Fatal("expected an error when session not found, got nil")
	}
	if got != nil {
		t.Errorf("expected nil pointer on not-found, got %+v", got)
	}
}

// --- ValidateEmail ---

func TestValidateEmail_Valid(t *testing.T) {
	if err := stage6.ValidateEmail("deeqa@somnog.org"); err != nil {
		t.Errorf("expected nil error for valid email, got %v", err)
	}
}

func TestValidateEmail_Empty(t *testing.T) {
	if err := stage6.ValidateEmail(""); err == nil {
		t.Error("expected error for empty email, got nil")
	}
}

func TestValidateEmail_NoAt(t *testing.T) {
	if err := stage6.ValidateEmail("khalidsomnog.org"); err == nil {
		t.Error("expected error for email missing @, got nil")
	}
}

func TestValidateEmail_Whitespace(t *testing.T) {
	if err := stage6.ValidateEmail("ayan @somnog.org"); err == nil {
		t.Error("expected error for email with whitespace, got nil")
	}
}
