package stage5_test

import (
	"strings"
	"testing"

	"github.com/somnog/learn-go/stage5"
)

// ---------------------------------------------------------------------------
// NewSession
// ---------------------------------------------------------------------------

func TestNewSession_Fields(t *testing.T) {
	s := stage5.NewSession(1, "BGP Routing at Scale", "Network Infrastructure", 50)
	if s.ID != 1 {
		t.Errorf("NewSession: ID = %d; want 1", s.ID)
	}
	if s.Title != "BGP Routing at Scale" {
		t.Errorf("NewSession: Title = %q; want %q", s.Title, "BGP Routing at Scale")
	}
	if s.Track != "Network Infrastructure" {
		t.Errorf("NewSession: Track = %q; want %q", s.Track, "Network Infrastructure")
	}
	if s.Capacity != 50 {
		t.Errorf("NewSession: Capacity = %d; want 50", s.Capacity)
	}
}

func TestNewSession_EnrolledIsZero(t *testing.T) {
	s := stage5.NewSession(2, "Go Workshop", "Software Development", 30)
	if s.Enrolled != 0 {
		t.Errorf("NewSession: Enrolled = %d; want 0 (starts empty)", s.Enrolled)
	}
}

// ---------------------------------------------------------------------------
// IsFull
// ---------------------------------------------------------------------------

func TestIsFull_False(t *testing.T) {
	s := stage5.Session{ID: 1, Title: "BGP Routing", Capacity: 50, Enrolled: 10}
	if s.IsFull() {
		t.Error("IsFull() = true; want false (10 enrolled, 50 capacity)")
	}
}

func TestIsFull_True_AtCapacity(t *testing.T) {
	s := stage5.Session{ID: 2, Title: "Go Workshop", Capacity: 30, Enrolled: 30}
	if !s.IsFull() {
		t.Error("IsFull() = false; want true (30 enrolled, 30 capacity)")
	}
}

func TestIsFull_True_OverCapacity(t *testing.T) {
	s := stage5.Session{ID: 3, Title: "DNS Security", Capacity: 20, Enrolled: 25}
	if !s.IsFull() {
		t.Error("IsFull() = false; want true (25 enrolled, 20 capacity — over)")
	}
}

func TestIsFull_EmptySession(t *testing.T) {
	s := stage5.Session{ID: 4, Title: "IPv6", Capacity: 40, Enrolled: 0}
	if s.IsFull() {
		t.Error("IsFull() = true; want false (0 enrolled)")
	}
}

// ---------------------------------------------------------------------------
// SpotsLeft
// ---------------------------------------------------------------------------

func TestSpotsLeft_Normal(t *testing.T) {
	s := stage5.Session{Capacity: 50, Enrolled: 12}
	got := s.SpotsLeft()
	if got != 38 {
		t.Errorf("SpotsLeft() = %d; want 38 (50 - 12)", got)
	}
}

func TestSpotsLeft_Full(t *testing.T) {
	s := stage5.Session{Capacity: 30, Enrolled: 30}
	got := s.SpotsLeft()
	if got != 0 {
		t.Errorf("SpotsLeft() = %d; want 0 (session full)", got)
	}
}

func TestSpotsLeft_OverCapacity_ReturnsZero(t *testing.T) {
	s := stage5.Session{Capacity: 20, Enrolled: 25}
	got := s.SpotsLeft()
	if got < 0 {
		t.Errorf("SpotsLeft() = %d; want 0 or more (never negative)", got)
	}
	if got != 0 {
		t.Errorf("SpotsLeft() = %d; want 0 when over capacity", got)
	}
}

func TestSpotsLeft_AllSpotsOpen(t *testing.T) {
	s := stage5.Session{Capacity: 50, Enrolled: 0}
	got := s.SpotsLeft()
	if got != 50 {
		t.Errorf("SpotsLeft() = %d; want 50", got)
	}
}

// ---------------------------------------------------------------------------
// Summary
// ---------------------------------------------------------------------------

var testSession = stage5.Session{
	ID:    1,
	Title: "BGP Routing at Scale",
	Speaker: stage5.Speaker{
		Name:  "Abdi Hassan",
		Title: "CTO, AfricaNet",
	},
	Track:    "Network Infrastructure",
	Capacity: 50,
	Enrolled: 12,
}

func TestSummary_ContainsTitle(t *testing.T) {
	got := testSession.Summary()
	if !strings.Contains(got, "BGP Routing at Scale") {
		t.Errorf("Summary() = %q; want it to contain the session title", got)
	}
}

func TestSummary_ContainsSpeakerName(t *testing.T) {
	got := testSession.Summary()
	if !strings.Contains(got, "Abdi Hassan") {
		t.Errorf("Summary() = %q; want it to contain the speaker name", got)
	}
}

func TestSummary_ContainsSpotsLeft(t *testing.T) {
	got := testSession.Summary()
	// Session has 50 - 12 = 38 spots left
	if !strings.Contains(got, "38") {
		t.Errorf("Summary() = %q; want it to contain the spots left count (38)", got)
	}
}

func TestSummary_NotEmpty(t *testing.T) {
	s := stage5.Session{
		ID:    2,
		Title: "Go Workshop",
		Speaker: stage5.Speaker{Name: "Hassan Warsame", Title: "Head of Engineering, NetWatch"},
		Track:    "Software Development",
		Capacity: 30,
		Enrolled: 5,
	}
	got := s.Summary()
	if got == "" {
		t.Error("Summary() returned empty string; want non-empty summary")
	}
}

func TestSummary_DifferentSession(t *testing.T) {
	s := stage5.Session{
		ID:    3,
		Title: "DNS Security Extensions",
		Speaker: stage5.Speaker{Name: "Layla Ibrahim", Title: "CISO, SecureNet"},
		Track:    "Cybersecurity",
		Capacity: 40,
		Enrolled: 40,
	}
	got := s.Summary()
	if !strings.Contains(got, "DNS Security Extensions") {
		t.Errorf("Summary() = %q; want it to contain 'DNS Security Extensions'", got)
	}
	if !strings.Contains(got, "Layla Ibrahim") {
		t.Errorf("Summary() = %q; want it to contain 'Layla Ibrahim'", got)
	}
}
