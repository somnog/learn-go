package stage4_test

import (
	"testing"

	"github.com/somnog/learn-go/stage4"
)

// ---------------------------------------------------------------------------
// NewRegistry
// ---------------------------------------------------------------------------

func TestNewRegistry_NotNil(t *testing.T) {
	r := stage4.NewRegistry()
	if r == nil {
		t.Error("NewRegistry() returned nil; want an initialised map")
	}
}

func TestNewRegistry_Empty(t *testing.T) {
	r := stage4.NewRegistry()
	if len(r) != 0 {
		t.Errorf("NewRegistry() length = %d; want 0", len(r))
	}
}

func TestNewRegistry_Writable(t *testing.T) {
	r := stage4.NewRegistry()
	// Writing to a nil map panics — this confirms it is initialised
	r["test@somnog.so"] = []string{"BGP Routing"}
	if len(r) != 1 {
		t.Error("NewRegistry(): map is not writable after initialisation")
	}
}

// ---------------------------------------------------------------------------
// Register
// ---------------------------------------------------------------------------

func TestRegister_AddsSession(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "abdi@somnog.so", "BGP Routing at Scale")
	sessions := stage4.Sessions(r, "abdi@somnog.so")
	if len(sessions) != 1 {
		t.Fatalf("After Register: Sessions length = %d; want 1", len(sessions))
	}
	if sessions[0] != "BGP Routing at Scale" {
		t.Errorf("Sessions[0] = %q; want %q", sessions[0], "BGP Routing at Scale")
	}
}

func TestRegister_MultipleSessionsSameAttendee(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "layla@somnog.so", "BGP Routing at Scale")
	stage4.Register(r, "layla@somnog.so", "Go Workshop")
	sessions := stage4.Sessions(r, "layla@somnog.so")
	if len(sessions) != 2 {
		t.Errorf("After two Registers: Sessions length = %d; want 2", len(sessions))
	}
}

func TestRegister_MultipleAttendees(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "abdi@somnog.so", "BGP Routing at Scale")
	stage4.Register(r, "ayan@somnog.so", "Go Workshop")
	if stage4.AttendeeCount(r) != 2 {
		t.Errorf("AttendeeCount after two different attendees = %d; want 2", stage4.AttendeeCount(r))
	}
}

// ---------------------------------------------------------------------------
// Sessions
// ---------------------------------------------------------------------------

func TestSessions_ReturnsRegisteredSessions(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "nimo@somnog.so", "DNS Security")
	stage4.Register(r, "nimo@somnog.so", "IPv6 Deployment")
	got := stage4.Sessions(r, "nimo@somnog.so")
	if len(got) != 2 {
		t.Errorf("Sessions: length = %d; want 2", len(got))
	}
}

func TestSessions_MissingEmail(t *testing.T) {
	r := stage4.NewRegistry()
	got := stage4.Sessions(r, "nobody@somnog.so")
	if len(got) != 0 {
		t.Errorf("Sessions for unknown email: length = %d; want 0", len(got))
	}
}

// ---------------------------------------------------------------------------
// IsRegistered
// ---------------------------------------------------------------------------

func TestIsRegistered_True(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "hassan@somnog.so", "Go Workshop")
	if !stage4.IsRegistered(r, "hassan@somnog.so", "Go Workshop") {
		t.Error("IsRegistered: expected true, got false")
	}
}

func TestIsRegistered_False_WrongSession(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "hassan@somnog.so", "Go Workshop")
	if stage4.IsRegistered(r, "hassan@somnog.so", "BGP Routing at Scale") {
		t.Error("IsRegistered: expected false for different session, got true")
	}
}

func TestIsRegistered_False_UnknownEmail(t *testing.T) {
	r := stage4.NewRegistry()
	if stage4.IsRegistered(r, "nobody@somnog.so", "Go Workshop") {
		t.Error("IsRegistered: expected false for unknown email, got true")
	}
}

func TestIsRegistered_AfterMultipleRegistrations(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "khalid@somnog.so", "BGP Routing at Scale")
	stage4.Register(r, "khalid@somnog.so", "DNS Security")
	stage4.Register(r, "khalid@somnog.so", "Go Workshop")
	if !stage4.IsRegistered(r, "khalid@somnog.so", "DNS Security") {
		t.Error("IsRegistered: expected true for 'DNS Security', got false")
	}
}

// ---------------------------------------------------------------------------
// AttendeeCount
// ---------------------------------------------------------------------------

func TestAttendeeCount_Empty(t *testing.T) {
	r := stage4.NewRegistry()
	if stage4.AttendeeCount(r) != 0 {
		t.Errorf("AttendeeCount on empty registry = %d; want 0", stage4.AttendeeCount(r))
	}
}

func TestAttendeeCount_OneAttendee(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "abdi@somnog.so", "Go Workshop")
	stage4.Register(r, "abdi@somnog.so", "BGP Routing")  // same person, different session
	if stage4.AttendeeCount(r) != 1 {
		t.Errorf("AttendeeCount for one unique email = %d; want 1", stage4.AttendeeCount(r))
	}
}

func TestAttendeeCount_ThreeAttendees(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "abdi@somnog.so", "Go Workshop")
	stage4.Register(r, "layla@somnog.so", "BGP Routing")
	stage4.Register(r, "ayan@somnog.so", "DNS Security")
	if stage4.AttendeeCount(r) != 3 {
		t.Errorf("AttendeeCount for 3 unique emails = %d; want 3", stage4.AttendeeCount(r))
	}
}

// ---------------------------------------------------------------------------
// Unregister
// ---------------------------------------------------------------------------

func TestUnregister_RemovesSession(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "deeqa@somnog.so", "Go Workshop")
	stage4.Register(r, "deeqa@somnog.so", "BGP Routing")
	stage4.Unregister(r, "deeqa@somnog.so", "Go Workshop")
	sessions := stage4.Sessions(r, "deeqa@somnog.so")
	if len(sessions) != 1 {
		t.Errorf("After Unregister: Sessions length = %d; want 1", len(sessions))
	}
	if stage4.IsRegistered(r, "deeqa@somnog.so", "Go Workshop") {
		t.Error("After Unregister: 'Go Workshop' still registered")
	}
}

func TestUnregister_RemovesEmailWhenNoSessionsLeft(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "omar@somnog.so", "Go Workshop")
	stage4.Unregister(r, "omar@somnog.so", "Go Workshop")
	if stage4.AttendeeCount(r) != 0 {
		t.Errorf("After unregistering last session: AttendeeCount = %d; want 0", stage4.AttendeeCount(r))
	}
}

func TestUnregister_NonExistentSession_NoEffect(t *testing.T) {
	r := stage4.NewRegistry()
	stage4.Register(r, "hodan@somnog.so", "Go Workshop")
	stage4.Unregister(r, "hodan@somnog.so", "BGP Routing")  // was never registered
	sessions := stage4.Sessions(r, "hodan@somnog.so")
	if len(sessions) != 1 {
		t.Errorf("Unregister of non-existent session changed the list: length = %d; want 1", len(sessions))
	}
}
