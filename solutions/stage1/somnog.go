package stage1

// ConferenceName returns the name of the conference.
func ConferenceName() string {
	return "SomNOG9"
}

// Year returns the year of the conference.
func Year() int {
	return 2026
}

// MaxAttendees returns the maximum number of attendees.
func MaxAttendees() int {
	return 200
}

// IsRegistrationOpen returns whether registration is currently open.
func IsRegistrationOpen() bool {
	return true
}

// WelcomeMessage returns a personalised welcome string.
func WelcomeMessage(name string) string {
	return "Welcome to SomNOG9, " + name + "!"
}
