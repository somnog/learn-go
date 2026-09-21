package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ── Data types ────────────────────────────────────────────────────────────────

type Speaker struct {
	Name  string `json:"name"`
	Title string `json:"title"`
}

type Session struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Speaker   Speaker  `json:"speaker"`
	Track     string   `json:"track"`
	Capacity  int      `json:"capacity"`
	Attendees []string `json:"attendees"`
}

// ── In-memory store ───────────────────────────────────────────────────────────

var store []Session

func seedStore() {
	store = []Session{
		{
			ID:        1,
			Title:     "BGP for Operators",
			Speaker:   Speaker{Name: "Abdi Hassan", Title: "CTO AfricaNet"},
			Track:     "Networking",
			Capacity:  30,
			Attendees: []string{},
		},
		{
			ID:        2,
			Title:     "Intro to Go",
			Speaker:   Speaker{Name: "Layla Ibrahim", Title: "CISO SecureNet"},
			Track:     "Software Development",
			Capacity:  20,
			Attendees: []string{},
		},
	}
}

// ResetStore resets the store to the seed data.
// Called by tests before each test case to ensure a clean state.
func ResetStore() {
	seedStore()
}

// ── Router ────────────────────────────────────────────────────────────────────

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/sessions", handleAll)
	mux.HandleFunc("/sessions/", handleOne)
	return mux
}

// ── Handlers ──────────────────────────────────────────────────────────────────

// handleAll handles:
//
//	GET  /sessions        – list all sessions
//	POST /sessions        – create a new session
func handleAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// TODO: respond with the full list of sessions as JSON (status 200)
		//
		//   Hint: json.NewEncoder(w).Encode(store)
		//   Don't forget to set Content-Type: application/json first.
		w.WriteHeader(http.StatusNotImplemented)

	case http.MethodPost:
		// TODO: decode a Session from the request body, assign it a new ID
		// (len(store)+1), append it to store, and respond with 201 Created
		// and the created session as JSON.
		//
		//   Hint: var s Session
		//         json.NewDecoder(r.Body).Decode(&s)
		w.WriteHeader(http.StatusNotImplemented)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleOne handles:
//
//	GET /sessions/{id}             – return a single session
//	GET /sessions/{id}/attendees   – return attendee list
func handleOne(w http.ResponseWriter, r *http.Request) {
	// Parse the path: /sessions/{id} or /sessions/{id}/attendees
	path := strings.TrimPrefix(r.URL.Path, "/sessions/")
	parts := strings.SplitN(path, "/", 2)

	idStr := parts[0]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid session id", http.StatusBadRequest)
		return
	}

	sub := ""
	if len(parts) == 2 {
		sub = parts[1]
	}

	switch sub {
	case "":
		// GET /sessions/{id}
		s, ok := findByID(id)
		if !ok {
			http.Error(w, "session not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)

	case "attendees":
		// TODO: look up the session by id; if not found respond 404.
		// Otherwise respond 200 with the session's Attendees slice as JSON.
		//
		//   Hint: s, ok := findByID(id)
		w.WriteHeader(http.StatusNotImplemented)

	default:
		http.Error(w, "unknown sub-resource", http.StatusNotFound)
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

// findByID searches store for a session with the given ID.
// Returns a pointer to the session and true if found; nil and false otherwise.
func findByID(id int) (*Session, bool) {
	for i := range store {
		if store[i].ID == id {
			return &store[i], true
		}
	}
	return nil, false
}

// ── Entry point ───────────────────────────────────────────────────────────────

func main() {
	seedStore()
	fmt.Println("SomNOG9 API listening on :8080")
	http.ListenAndServe(":8080", NewRouter())
}
