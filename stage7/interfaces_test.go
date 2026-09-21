package stage7_test

import (
	"strings"
	"testing"

	"github.com/somnog/learn-go/stage7"
)

// --- Speaker.Describe ---

func TestSpeakerDescribe(t *testing.T) {
	sp := stage7.Speaker{Name: "Abdi Hassan", Title: "CTO AfricaNet"}
	desc := sp.Describe()
	if !strings.Contains(desc, "Abdi Hassan") {
		t.Errorf("Speaker.Describe() missing name: %q", desc)
	}
	if !strings.Contains(desc, "CTO AfricaNet") {
		t.Errorf("Speaker.Describe() missing title: %q", desc)
	}
}

// --- Session.Describe ---

func TestSessionDescribe(t *testing.T) {
	s := stage7.Session{
		ID:    1,
		Title: "BGP for Operators",
	}
	desc := s.Describe()
	if !strings.Contains(desc, "BGP for Operators") {
		t.Errorf("Session.Describe() missing title: %q", desc)
	}
	if !strings.Contains(desc, "1") {
		t.Errorf("Session.Describe() missing ID: %q", desc)
	}
}

// --- Workshop.Describe ---

func TestWorkshopDescribe(t *testing.T) {
	w := stage7.Workshop{
		ID:              3,
		Title:           "Intro to Go",
		Instructor:      "Layla Ibrahim",
		DurationMinutes: 120,
	}
	desc := w.Describe()
	if !strings.Contains(desc, "Intro to Go") {
		t.Errorf("Workshop.Describe() missing title: %q", desc)
	}
	if !strings.Contains(desc, "120") {
		t.Errorf("Workshop.Describe() missing duration: %q", desc)
	}
}

// --- DescribeAll ---

func TestDescribeAll(t *testing.T) {
	items := []stage7.Describable{
		stage7.Speaker{Name: "Hassan Warsame", Title: "Engineer"},
		stage7.Session{ID: 2, Title: "DNS Security"},
		stage7.Workshop{ID: 4, Title: "IPv6 Lab", Instructor: "Nimo Abdi", DurationMinutes: 90},
	}
	results := stage7.DescribeAll(items)
	if len(results) != 3 {
		t.Fatalf("DescribeAll returned %d items, want 3", len(results))
	}
	for i, r := range results {
		if r == "" {
			t.Errorf("DescribeAll result[%d] is empty", i)
		}
	}
}

// --- FindByTitle ---

func TestFindByTitle_Found(t *testing.T) {
	items := []stage7.Describable{
		stage7.Session{ID: 1, Title: "BGP for Operators"},
		stage7.Workshop{ID: 2, Title: "Intro to Go", Instructor: "Omar Sharif", DurationMinutes: 60},
	}
	got, ok := stage7.FindByTitle(items, "intro to go")
	if !ok {
		t.Fatal("FindByTitle returned false, expected true (case-insensitive match)")
	}
	if got == nil {
		t.Fatal("FindByTitle returned nil, expected a Describable")
	}
}

func TestFindByTitle_NotFound(t *testing.T) {
	items := []stage7.Describable{
		stage7.Session{ID: 1, Title: "BGP for Operators"},
	}
	got, ok := stage7.FindByTitle(items, "nonexistent session")
	if ok {
		t.Error("FindByTitle returned true for non-existent title")
	}
	if got != nil {
		t.Errorf("FindByTitle returned non-nil on not-found: %+v", got)
	}
}
