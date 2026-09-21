package stage4

// NewRegistry creates and returns an empty attendee registry.
// The registry maps an attendee email to the list of sessions they are registered for.
func NewRegistry() map[string][]string {
        return make(map[string][]string)
}

// Register adds session to the list of sessions for email.
func Register(r map[string][]string, email, session string) {
        r[email] = append(r[email], session)
}

// Sessions returns all sessions registered for email.
// Returns an empty (nil) slice if the email is not in the registry.
func Sessions(r map[string][]string, email string) []string {
        return r[email]
}

// IsRegistered reports whether email is registered for session.
func IsRegistered(r map[string][]string, email, session string) bool {
        for _, s := range r[email] {
                if s == session {
                        return true
                }
        }
        return false
}

// AttendeeCount returns the number of unique attendees in the registry.
func AttendeeCount(r map[string][]string) int {
        return len(r)
}

// Unregister removes session from email's list.
// If the list becomes empty after removal, the email key is deleted entirely.
func Unregister(r map[string][]string, email, session string) {
        var updated []string
        for _, s := range r[email] {
                if s != session {
                        updated = append(updated, s)
                }
        }
        if len(updated) == 0 {
                delete(r, email)
        } else {
                r[email] = updated
        }
}
