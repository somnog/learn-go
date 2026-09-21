package stage3_test

import (
	"testing"

	"github.com/somnog/learn-go/stage3"
)

// ---------------------------------------------------------------------------
// Tracks
// ---------------------------------------------------------------------------

func TestTracks_Length(t *testing.T) {
	got := stage3.Tracks()
	if len(got) != 4 {
		t.Errorf("Tracks() returned %d items; want 4", len(got))
	}
}

func TestTracks_ContainsNetworkInfrastructure(t *testing.T) {
	tracks := stage3.Tracks()
	found := false
	for _, tr := range tracks {
		if tr == "Network Infrastructure" {
			found = true
		}
	}
	if !found {
		t.Errorf("Tracks() does not contain 'Network Infrastructure'; got %v", tracks)
	}
}

func TestTracks_ContainsSoftwareDevelopment(t *testing.T) {
	tracks := stage3.Tracks()
	found := false
	for _, tr := range tracks {
		if tr == "Software Development" {
			found = true
		}
	}
	if !found {
		t.Errorf("Tracks() does not contain 'Software Development'; got %v", tracks)
	}
}

func TestTracks_Order(t *testing.T) {
	tracks := stage3.Tracks()
	if len(tracks) < 4 {
		t.Fatalf("Tracks() returned fewer than 4 elements")
	}
	if tracks[0] != "Network Infrastructure" {
		t.Errorf("Tracks()[0] = %q; want %q", tracks[0], "Network Infrastructure")
	}
	if tracks[2] != "Software Development" {
		t.Errorf("Tracks()[2] = %q; want %q", tracks[2], "Software Development")
	}
	if tracks[3] != "Cybersecurity" {
		t.Errorf("Tracks()[3] = %q; want %q", tracks[3], "Cybersecurity")
	}
}

// ---------------------------------------------------------------------------
// Add
// ---------------------------------------------------------------------------

func TestAdd_IncreasesLength(t *testing.T) {
	original := []string{"Network Infrastructure", "System & Services"}
	result := stage3.Add(original, "Software Development")
	if len(result) != 3 {
		t.Errorf("Add: result length = %d; want 3", len(result))
	}
}

func TestAdd_ContainsNewItem(t *testing.T) {
	original := []string{"Network Infrastructure"}
	result := stage3.Add(original, "Cybersecurity")
	last := result[len(result)-1]
	if last != "Cybersecurity" {
		t.Errorf("Add: last element = %q; want %q", last, "Cybersecurity")
	}
}

func TestAdd_DoesNotModifyOriginal(t *testing.T) {
	original := []string{"Network Infrastructure"}
	_ = stage3.Add(original, "Cybersecurity")
	if len(original) != 1 {
		t.Errorf("Add: original slice was modified; want length 1, got %d", len(original))
	}
}

func TestAdd_ToEmptySlice(t *testing.T) {
	result := stage3.Add([]string{}, "Network Infrastructure")
	if len(result) != 1 {
		t.Errorf("Add to empty slice: length = %d; want 1", len(result))
	}
}

// ---------------------------------------------------------------------------
// First
// ---------------------------------------------------------------------------

func TestFirst_NormalCase(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	got := stage3.First(items, 2)
	if len(got) != 2 {
		t.Fatalf("First(items, 2) length = %d; want 2", len(got))
	}
	if got[0] != "a" || got[1] != "b" {
		t.Errorf("First(items, 2) = %v; want [a b]", got)
	}
}

func TestFirst_NGreaterThanLength(t *testing.T) {
	items := []string{"a", "b"}
	got := stage3.First(items, 10)
	if len(got) != 2 {
		t.Errorf("First(items, 10) when len=2: length = %d; want 2 (return all)", len(got))
	}
}

func TestFirst_NIsZero(t *testing.T) {
	items := []string{"a", "b", "c"}
	got := stage3.First(items, 0)
	if len(got) != 0 {
		t.Errorf("First(items, 0): length = %d; want 0", len(got))
	}
}

func TestFirst_NIsNegative(t *testing.T) {
	items := []string{"a", "b", "c"}
	got := stage3.First(items, -1)
	if len(got) != 0 {
		t.Errorf("First(items, -1): length = %d; want 0", len(got))
	}
}

func TestFirst_ExactLength(t *testing.T) {
	items := []string{"a", "b", "c"}
	got := stage3.First(items, 3)
	if len(got) != 3 {
		t.Errorf("First(items, 3) when len=3: length = %d; want 3", len(got))
	}
}

// ---------------------------------------------------------------------------
// Contains
// ---------------------------------------------------------------------------

func TestContains_Found(t *testing.T) {
	items := []string{"Network Infrastructure", "System & Services", "Software Development"}
	if !stage3.Contains(items, "Software Development") {
		t.Error("Contains: expected true for 'Software Development', got false")
	}
}

func TestContains_NotFound(t *testing.T) {
	items := []string{"Network Infrastructure", "System & Services"}
	if stage3.Contains(items, "Cybersecurity") {
		t.Error("Contains: expected false for 'Cybersecurity', got true")
	}
}

func TestContains_EmptySlice(t *testing.T) {
	if stage3.Contains([]string{}, "anything") {
		t.Error("Contains on empty slice: expected false, got true")
	}
}

func TestContains_CaseSensitive(t *testing.T) {
	items := []string{"Network Infrastructure"}
	if stage3.Contains(items, "network infrastructure") {
		t.Error("Contains: expected false for lowercase (case-sensitive), got true")
	}
}

// ---------------------------------------------------------------------------
// MakeSchedule
// ---------------------------------------------------------------------------

func TestMakeSchedule_Length(t *testing.T) {
	got := stage3.MakeSchedule(5)
	if len(got) != 5 {
		t.Errorf("MakeSchedule(5): length = %d; want 5", len(got))
	}
}

func TestMakeSchedule_AllEmpty(t *testing.T) {
	got := stage3.MakeSchedule(3)
	for i, v := range got {
		if v != "" {
			t.Errorf("MakeSchedule(3)[%d] = %q; want empty string", i, v)
		}
	}
}

func TestMakeSchedule_Zero(t *testing.T) {
	got := stage3.MakeSchedule(0)
	if len(got) != 0 {
		t.Errorf("MakeSchedule(0): length = %d; want 0", len(got))
	}
}

// ---------------------------------------------------------------------------
// Reverse
// ---------------------------------------------------------------------------

func TestReverse_Order(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	got := stage3.Reverse(items)
	want := []string{"d", "c", "b", "a"}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("Reverse()[%d] = %q; want %q", i, got[i], want[i])
		}
	}
}

func TestReverse_DoesNotModifyOriginal(t *testing.T) {
	items := []string{"a", "b", "c"}
	_ = stage3.Reverse(items)
	if items[0] != "a" {
		t.Errorf("Reverse modified the original slice; items[0] = %q, want 'a'", items[0])
	}
}

func TestReverse_SingleElement(t *testing.T) {
	got := stage3.Reverse([]string{"only"})
	if len(got) != 1 || got[0] != "only" {
		t.Errorf("Reverse(single) = %v; want [only]", got)
	}
}

func TestReverse_EmptySlice(t *testing.T) {
	got := stage3.Reverse([]string{})
	if len(got) != 0 {
		t.Errorf("Reverse(empty) length = %d; want 0", len(got))
	}
}

func TestReverse_TrackNames(t *testing.T) {
	tracks := []string{"Network Infrastructure", "System & Services", "Software Development", "Cybersecurity"}
	got := stage3.Reverse(tracks)
	if got[0] != "Cybersecurity" {
		t.Errorf("Reverse(tracks)[0] = %q; want %q", got[0], "Cybersecurity")
	}
	if got[3] != "Network Infrastructure" {
		t.Errorf("Reverse(tracks)[3] = %q; want %q", got[3], "Network Infrastructure")
	}
}
