package stage1_test

import (
	"strings"
	"testing"

	"github.com/somnog/learn-go/stage1"
)

func TestConferenceName(t *testing.T) {
	got := stage1.ConferenceName()
	want := "SomNOG9"
	if got != want {
		t.Errorf("ConferenceName() = %q; want %q", got, want)
	}
}

func TestYear(t *testing.T) {
	got := stage1.Year()
	want := 2026
	if got != want {
		t.Errorf("Year() = %d; want %d", got, want)
	}
}

func TestMaxAttendees(t *testing.T) {
	got := stage1.MaxAttendees()
	if got <= 0 {
		t.Errorf("MaxAttendees() = %d; want a positive number", got)
	}
	want := 200
	if got != want {
		t.Errorf("MaxAttendees() = %d; want %d", got, want)
	}
}

func TestIsRegistrationOpen(t *testing.T) {
	got := stage1.IsRegistrationOpen()
	if !got {
		t.Errorf("IsRegistrationOpen() = %v; want true", got)
	}
}

func TestWelcomeMessage_ContainsName(t *testing.T) {
	msg := stage1.WelcomeMessage("Abdi Hassan")
	if !strings.Contains(msg, "Abdi Hassan") {
		t.Errorf("WelcomeMessage(%q) = %q; want it to contain the name", "Abdi Hassan", msg)
	}
}

func TestWelcomeMessage_ContainsSomNOG9(t *testing.T) {
	msg := stage1.WelcomeMessage("Layla Ibrahim")
	if !strings.Contains(msg, "SomNOG9") {
		t.Errorf("WelcomeMessage(%q) = %q; want it to contain 'SomNOG9'", "Layla Ibrahim", msg)
	}
}

func TestWelcomeMessage_DifferentNames(t *testing.T) {
	names := []string{"Abdi Hassan", "Layla Ibrahim", "Ayan Mohamed", "Khalid Omar"}
	for _, name := range names {
		msg := stage1.WelcomeMessage(name)
		if !strings.Contains(msg, name) {
			t.Errorf("WelcomeMessage(%q) = %q; want it to contain the name", name, msg)
		}
	}
}
