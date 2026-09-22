package stage1

import "testing"

func TestConferenceName(t *testing.T) {
	got := ConferenceName()
	want := "SomNOG9"

	if got != want {
		t.Errorf("ConferenceName() = %q, want %q", got, want)
	}
}

func TestYear(t *testing.T) {
	got := Year()
	want := 2026

	if got != want {
		t.Errorf("Year() = %d, want %d", got, want)
	}
}

func TestMaxAttendees(t *testing.T) {
	got := MaxAttendees()
	want := 200

	if got != want {
		t.Errorf("MaxAttendees() = %d, want %d", got, want)
	}
}

func TestIsRegistrationOpen(t *testing.T) {
	got := IsRegistrationOpen()
	want := true

	if got != want {
		t.Errorf("IsRegistrationOpen() = %t, want %t", got, want)
	}
}

func TestWelcomeMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Ahmed",
			want: "Welcome to SomNOG9, Ahmed!",
		},
		{
			name: "Fatima",
			want: "Welcome to SomNOG9, Fatima!",
		},
		{
			name: "Go Developer",
			want: "Welcome to SomNOG9, Go Developer!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := WelcomeMessage(tt.name)

			if got != tt.want {
				t.Errorf("WelcomeMessage(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
