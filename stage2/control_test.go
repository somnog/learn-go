package stage2_test

import (
	"strings"
	"testing"

	"github.com/somnog/learn-go/stage2"
)

// ---------------------------------------------------------------------------
// TrackName
// ---------------------------------------------------------------------------

func TestTrackName_Track1(t *testing.T) {
	got := stage2.TrackName(1)
	want := "Network Infrastructure"
	if got != want {
		t.Errorf("TrackName(1) = %q; want %q", got, want)
	}
}

func TestTrackName_Track2(t *testing.T) {
	got := stage2.TrackName(2)
	want := "System & Services"
	if got != want {
		t.Errorf("TrackName(2) = %q; want %q", got, want)
	}
}

func TestTrackName_Track3(t *testing.T) {
	got := stage2.TrackName(3)
	want := "Software Development"
	if got != want {
		t.Errorf("TrackName(3) = %q; want %q", got, want)
	}
}

func TestTrackName_Track4(t *testing.T) {
	got := stage2.TrackName(4)
	want := "Cybersecurity"
	if got != want {
		t.Errorf("TrackName(4) = %q; want %q", got, want)
	}
}

func TestTrackName_Zero_IsUnknown(t *testing.T) {
	got := stage2.TrackName(0)
	want := "Unknown"
	if got != want {
		t.Errorf("TrackName(0) = %q; want %q", got, want)
	}
}

func TestTrackName_Negative_IsUnknown(t *testing.T) {
	got := stage2.TrackName(-1)
	want := "Unknown"
	if got != want {
		t.Errorf("TrackName(-1) = %q; want %q", got, want)
	}
}

func TestTrackName_OutOfRange_IsUnknown(t *testing.T) {
	got := stage2.TrackName(99)
	want := "Unknown"
	if got != want {
		t.Errorf("TrackName(99) = %q; want %q", got, want)
	}
}

// ---------------------------------------------------------------------------
// CanRegister
// ---------------------------------------------------------------------------

func TestCanRegister_SpotsAvailable(t *testing.T) {
	if !stage2.CanRegister(10, 50) {
		t.Error("CanRegister(10, 50) = false; want true (10 enrolled, 50 capacity)")
	}
}

func TestCanRegister_SessionFull(t *testing.T) {
	if stage2.CanRegister(50, 50) {
		t.Error("CanRegister(50, 50) = true; want false (session full)")
	}
}

func TestCanRegister_OverCapacity(t *testing.T) {
	if stage2.CanRegister(51, 50) {
		t.Error("CanRegister(51, 50) = true; want false (over capacity)")
	}
}

func TestCanRegister_OneSpotLeft(t *testing.T) {
	if !stage2.CanRegister(49, 50) {
		t.Error("CanRegister(49, 50) = false; want true (one spot left)")
	}
}

func TestCanRegister_EmptySession(t *testing.T) {
	if !stage2.CanRegister(0, 30) {
		t.Error("CanRegister(0, 30) = false; want true (no one enrolled yet)")
	}
}

// ---------------------------------------------------------------------------
// CountAvailable
// ---------------------------------------------------------------------------

func TestCountAvailable_AllOpen(t *testing.T) {
	enrolled := []int{5, 10, 15}
	got := stage2.CountAvailable(enrolled, 50)
	if got != 3 {
		t.Errorf("CountAvailable([5,10,15], 50) = %d; want 3", got)
	}
}

func TestCountAvailable_AllFull(t *testing.T) {
	enrolled := []int{50, 50, 50}
	got := stage2.CountAvailable(enrolled, 50)
	if got != 0 {
		t.Errorf("CountAvailable([50,50,50], 50) = %d; want 0", got)
	}
}

func TestCountAvailable_Mixed(t *testing.T) {
	enrolled := []int{10, 50, 30, 50, 5}
	got := stage2.CountAvailable(enrolled, 50)
	if got != 3 {
		t.Errorf("CountAvailable([10,50,30,50,5], 50) = %d; want 3", got)
	}
}

func TestCountAvailable_EmptySlice(t *testing.T) {
	got := stage2.CountAvailable([]int{}, 50)
	if got != 0 {
		t.Errorf("CountAvailable([], 50) = %d; want 0", got)
	}
}

// ---------------------------------------------------------------------------
// Greet
// ---------------------------------------------------------------------------

func TestGreet_Morning(t *testing.T) {
	got := stage2.Greet(9)
	if !strings.Contains(got, "morning") {
		t.Errorf("Greet(9) = %q; want it to contain 'morning'", got)
	}
	if !strings.Contains(got, "SomNOG9") {
		t.Errorf("Greet(9) = %q; want it to contain 'SomNOG9'", got)
	}
}

func TestGreet_Afternoon(t *testing.T) {
	got := stage2.Greet(14)
	if !strings.Contains(got, "afternoon") {
		t.Errorf("Greet(14) = %q; want it to contain 'afternoon'", got)
	}
	if !strings.Contains(got, "SomNOG9") {
		t.Errorf("Greet(14) = %q; want it to contain 'SomNOG9'", got)
	}
}

func TestGreet_Evening(t *testing.T) {
	got := stage2.Greet(19)
	if !strings.Contains(got, "evening") {
		t.Errorf("Greet(19) = %q; want it to contain 'evening'", got)
	}
	if !strings.Contains(got, "SomNOG9") {
		t.Errorf("Greet(19) = %q; want it to contain 'SomNOG9'", got)
	}
}

func TestGreet_Noon_IsAfternoon(t *testing.T) {
	got := stage2.Greet(12)
	if !strings.Contains(got, "afternoon") {
		t.Errorf("Greet(12) = %q; want it to contain 'afternoon' (noon is afternoon)", got)
	}
}

func TestGreet_Midnight_IsMorning(t *testing.T) {
	got := stage2.Greet(0)
	if !strings.Contains(got, "morning") {
		t.Errorf("Greet(0) = %q; want it to contain 'morning'", got)
	}
}
