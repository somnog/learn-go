package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ── GET /sessions ─────────────────────────────────────────────────────────────

func TestGetSessions(t *testing.T) {
	ResetStore()
	req := httptest.NewRequest(http.MethodGet, "/sessions", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET /sessions: want 200, got %d", w.Code)
	}

	var sessions []Session
	if err := json.NewDecoder(w.Body).Decode(&sessions); err != nil {
		t.Fatalf("GET /sessions: could not decode response: %v", err)
	}
	if len(sessions) == 0 {
		t.Error("GET /sessions: expected at least one session, got empty list")
	}
}

// ── POST /sessions ────────────────────────────────────────────────────────────

func TestPostSession(t *testing.T) {
	ResetStore()
	newSession := Session{
		Title:    "DNS Security",
		Speaker:  Speaker{Name: "Hassan Warsame", Title: "Network Engineer"},
		Track:    "Networking",
		Capacity: 25,
	}
	body, _ := json.Marshal(newSession)
	req := httptest.NewRequest(http.MethodPost, "/sessions", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("POST /sessions: want 201, got %d", w.Code)
	}

	var created Session
	if err := json.NewDecoder(w.Body).Decode(&created); err != nil {
		t.Fatalf("POST /sessions: could not decode response: %v", err)
	}
	if created.ID == 0 {
		t.Error("POST /sessions: created session should have a non-zero ID")
	}
	if created.Title != "DNS Security" {
		t.Errorf("POST /sessions: expected title 'DNS Security', got %q", created.Title)
	}
}

// ── GET /sessions/{id} ────────────────────────────────────────────────────────

func TestGetSessionByID_Found(t *testing.T) {
	ResetStore()
	req := httptest.NewRequest(http.MethodGet, "/sessions/1", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET /sessions/1: want 200, got %d", w.Code)
	}

	var s Session
	if err := json.NewDecoder(w.Body).Decode(&s); err != nil {
		t.Fatalf("GET /sessions/1: could not decode response: %v", err)
	}
	if s.ID != 1 {
		t.Errorf("GET /sessions/1: expected ID 1, got %d", s.ID)
	}
}

func TestGetSessionByID_NotFound(t *testing.T) {
	ResetStore()
	req := httptest.NewRequest(http.MethodGet, "/sessions/999", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("GET /sessions/999: want 404, got %d", w.Code)
	}
}

// ── GET /sessions/{id}/attendees ──────────────────────────────────────────────

func TestGetAttendees(t *testing.T) {
	ResetStore()
	req := httptest.NewRequest(http.MethodGet, "/sessions/1/attendees", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("GET /sessions/1/attendees: want 200, got %d", w.Code)
	}

	var attendees []string
	if err := json.NewDecoder(w.Body).Decode(&attendees); err != nil {
		t.Fatalf("GET /sessions/1/attendees: could not decode response: %v", err)
	}
	// Attendees may be empty for a fresh store — just ensure it's a valid array
	if attendees == nil {
		t.Error("GET /sessions/1/attendees: expected a JSON array, got null")
	}
}

func TestGetAttendees_NotFound(t *testing.T) {
	ResetStore()
	req := httptest.NewRequest(http.MethodGet, "/sessions/999/attendees", nil)
	w := httptest.NewRecorder()
	NewRouter().ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Fatalf("GET /sessions/999/attendees: want 404, got %d", w.Code)
	}
}
