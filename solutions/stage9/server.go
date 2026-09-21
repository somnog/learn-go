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

// ResetStore resets the store to the seed data (used by tests).
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

func handleAll(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(store)

	case http.MethodPost:
		var s Session
		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}
		s.ID = len(store) + 1
		if s.Attendees == nil {
			s.Attendees = []string{}
		}
		store = append(store, s)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(s)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleOne(w http.ResponseWriter, r *http.Request) {
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
		s, ok := findByID(id)
		if !ok {
			http.Error(w, "session not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s)

	case "attendees":
		s, ok := findByID(id)
		if !ok {
			http.Error(w, "session not found", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.Attendees)

	default:
		http.Error(w, "unknown sub-resource", http.StatusNotFound)
	}
}

// ── Helpers ───────────────────────────────────────────────────────────────────

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
