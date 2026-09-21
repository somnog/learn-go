# SomNOG9 Go Workshop

**Track 3: Software Development | SomNOG9, 22 September 2026**

You are building the SomNOG9 Conference Management System in Go — from variables to a running HTTP API — across four sessions. Every stage gives you pre-written tests. Your job is to write code that makes them pass.

---

## Sessions

| Session | Stages | Topics |
|---------|--------|--------|
| 1 | 0 → 2 | Setup, Variables & Types, Control Flow |
| 2 | 3 → 4 | Slices, Maps |
| 3 | 5 → 7 | Structs, Pointers & Errors, Interfaces |
| 4 | 8 → 9 | Goroutines & Channels, HTTP Server |

---

## How This Works

1. Each stage has a pre-written test file inside its directory.
2. You create an implementation file in the same directory.
3. Run `go test -v ./stageN/...` — all tests must pass before you move on.
4. If you are stuck, tell the presenter. A solution will be pushed to `solutions/stageN/` — run `git pull` to get it.
5. ⏸ markers mean: stop and wait for the presenter before moving to the next stage.

---

## Installation

### Linux (Ubuntu / Debian)

```bash
sudo apt update && sudo apt install -y golang-go
```

For the latest version:

```bash
wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
source ~/.bashrc
```

### macOS

```bash
brew install go
```

Or download the `.pkg` installer from [go.dev/dl](https://go.dev/dl).

### Windows

Download the `.msi` installer from [go.dev/dl](https://go.dev/dl) and run it. Open a new Command Prompt after installation.

### Verify

```bash
go version
# Expected: go version go1.21.x or higher
```

---

## Get the Repository

```bash
git clone https://github.com/somnog/learn-go
cd learn-go
```

---

## Stage 0 — Setup & Tooling

**Session 1** | ~15 minutes

Go organises code into **modules**. A module is a directory with a `go.mod` file. Open `go.mod`:

```
module github.com/somnog/learn-go

go 1.21
```

Every package in this repo imports using that module path — for example `github.com/somnog/learn-go/stage1`.

A **package** is a directory of `.go` files. Every file in that directory must share the same `package` declaration at the top.

`go test` is built into Go. No test runner, no config file, no framework needed.

```bash
go test -v ./stage0/...
```

Stage 0 has no code for you to write. If it passes, your workspace is ready.

```
--- PASS: TestGoVersion (0.00s)
--- PASS: TestModuleName (0.00s)
ok  github.com/somnog/learn-go/stage0
```

⏸ Wait for the presenter before moving to Stage 1.

---

## Stage 1 — Variables & Types

**Session 1** | ~30 minutes

### Key syntax

**Functions**

Every piece of Go code lives inside a function. A function declares what inputs it takes and what it returns.

```go
// No parameters, no return value
func PrintWelcome() {
    fmt.Println("Welcome to SomNOG9")
}

// Returns a value — state the type after the parentheses
func ConferenceName() string {
    return "SomNOG9"
}

func Year() int {
    return 2026
}

// Takes a parameter and returns a value
func WelcomeMessage(name string) string {
    return "Welcome to SomNOG9, " + name + "!"
}
```

**Variables**

```go
// Explicit — state the type
var name string = "SomNOG9"
var year int    = 2026
var isOpen bool = true

// Short declaration — type is inferred (use this inside functions)
title    := "Routing at Scale"
capacity := 50
open     := false

// Constant — cannot be reassigned
const maxAttendees = 200
```

**Types**

| Type | Example values | Zero value |
|------|----------------|------------|
| `string` | `"SomNOG9"`, `"Abdi Hassan"` | `""` |
| `int` | `1`, `50`, `2026` | `0` |
| `bool` | `true`, `false` | `false` |

Zero value: if you declare a variable without assigning, Go sets it to that type's zero.

**Building strings**

```go
import "fmt"

name := "Abdi Hassan"
msg  := fmt.Sprintf("Welcome to SomNOG9, %s!", name)
// %s = string, %d = integer, %t = bool
```

### Your task

Create the file `stage1/somnog.go`. The first two lines must be:

```go
package stage1

import "fmt"
```

Implement these five functions:

- `ConferenceName() string` — return `"SomNOG9"`
- `Year() int` — return `2026`
- `MaxAttendees() int` — return `200`
- `IsRegistrationOpen() bool` — return `true`
- `WelcomeMessage(name string) string` — return a string that contains both `"SomNOG9"` and the `name` argument

Inside each function, declare at least one variable with `:=` before returning it. Do not return a literal directly — that defeats the purpose of this stage.

**Hint:** Use `fmt.Sprintf` for `WelcomeMessage`.

### Run the tests

```bash
go test -v ./stage1/...
```

⏸ Wait for the presenter before moving to Stage 2.

---

## Stage 2 — Control Flow

**Session 1** | ~35 minutes

### Key syntax

**if / else if / else**

```go
if enrolled < capacity {
    fmt.Println("spots available")
} else if enrolled == capacity {
    fmt.Println("just full")
} else {
    fmt.Println("over capacity")
}
```

No parentheses around the condition. Braces are always required.

**switch**

```go
switch trackNumber {
case 1:
    return "Network Infrastructure"
case 2:
    return "System & Services"
case 3:
    return "Software Development"
case 4:
    return "Cybersecurity"
default:
    return "Unknown"
}
```

No `break` needed. Cases do not fall through.

**for — three forms**

```go
// Classic
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style
for count < capacity {
    count++
}

// Range over a slice
sessions := []string{"BGP Routing", "DNS Security", "Go Workshop"}
for i, name := range sessions {
    fmt.Printf("%d: %s\n", i, name)
}

// Discard the index with _
for _, name := range sessions {
    fmt.Println(name)
}
```

### Your task

Create the file `stage2/control.go`. Declare `package stage2` at the top.

Implement these four functions:

- `TrackName(n int) string` — return the track name for 1–4, `"Unknown"` for anything else. Use `switch`.
  - 1 → `"Network Infrastructure"`
  - 2 → `"System & Services"`
  - 3 → `"Software Development"`
  - 4 → `"Cybersecurity"`

- `CanRegister(enrolled, capacity int) bool` — return `true` if `enrolled < capacity`. Use `if`.

- `CountAvailable(enrolled []int, capacity int) int` — given a slice of enrolled counts and a shared capacity, return the number of sessions that still have open spots. Use `for`.

- `Greet(hour int) string` — return a greeting string that contains **both** a time-of-day word **and** `"SomNOG9"`:
  - `hour < 12` → contains `"morning"` and `"SomNOG9"` — e.g. `"Good morning at SomNOG9"`
  - `hour >= 12` and `hour < 17` → contains `"afternoon"` and `"SomNOG9"`
  - `hour >= 17` → contains `"evening"` and `"SomNOG9"`

### Run the tests

```bash
go test -v ./stage2/...
```

⏸ Wait for the presenter before moving to Stage 3.

---

## Stage 3 — Slices

**Session 2** | ~40 minutes

An **array** has a fixed size. A **slice** is dynamic — it can grow. Use slices.

### Key syntax

```go
// Slice literal
tracks := []string{"Network Infrastructure", "System & Services", "Software Development", "Cybersecurity"}

// Length
fmt.Println(len(tracks)) // 4

// Index — zero-based
fmt.Println(tracks[0]) // "Network Infrastructure"
fmt.Println(tracks[3]) // "Cybersecurity"

// Append — always reassign the result
tracks = append(tracks, "Community")

// Slice expressions — get a subset
first2 := tracks[:2]   // elements 0 and 1
from2  := tracks[2:]   // element 2 onwards
middle := tracks[1:3]  // elements 1 and 2

// make — create a slice of length n, all zero values
slots := make([]string, 5) // ["", "", "", "", ""]

// Range — iterate with index and value
for i, track := range tracks {
    fmt.Printf("%d: %s\n", i, track)
}
```

### Your task

Create the file `stage3/slices.go`. Declare `package stage3` at the top.

Implement these six functions:

- `Tracks() []string` — return a slice of the four SomNOG9 track names in this exact order:
  `"Network Infrastructure"`, `"System & Services"`, `"Software Development"`, `"Cybersecurity"`

- `Add(items []string, item string) []string` — append `item` to `items` and return the result

- `First(items []string, n int) []string` — return the first `n` elements. If `n >= len(items)`, return all elements. If `n <= 0`, return an empty slice.

- `Contains(items []string, target string) bool` — return `true` if `target` is in `items`. Use a `for` loop. Case-sensitive.

- `MakeSchedule(n int) []string` — return a slice of length `n` where every element is `""`. Use `make`.

- `Reverse(items []string) []string` — return a **new** slice with elements in reverse order. Do not modify the original slice.

**Hint for First:** use a slice expression `items[:n]`.

**Hint for Reverse:** `make([]string, len(items))`, then fill from the last index down.

### Run the tests

```bash
go test -v ./stage3/...
```

⏸ Wait for the presenter before moving to Stage 4.

---

## Stage 4 — Maps

**Session 2** | ~40 minutes

A **map** holds key–value pairs. Keys are unique. Maps are unordered.

In the SomNOG9 system: each attendee email maps to a list of session titles — `map[string][]string`.

### Key syntax

```go
// Create with make — always use make before writing to a map
registry := make(map[string][]string)

// Write
registry["abdi@somnog.so"] = []string{"BGP Routing", "Go Workshop"}

// Read — returns zero value if key missing
sessions := registry["abdi@somnog.so"]

// Check if a key exists — comma-ok idiom
sessions, exists := registry["layla@somnog.so"]
if !exists {
    fmt.Println("not registered")
}

// Delete a key
delete(registry, "abdi@somnog.so")

// Iterate — order is not guaranteed
for email, sessions := range registry {
    fmt.Println(email, sessions)
}

// Count of keys
count := len(registry)
```

Maps are **reference types**: when you pass a map to a function, both the caller and the function share the same underlying data. Changes inside the function are visible to the caller — no pointer needed.

### Your task

Create the file `stage4/registry.go`. Declare `package stage4` at the top.

Implement these six functions:

- `NewRegistry() map[string][]string` — return an initialised, empty map

- `Register(r map[string][]string, email, session string)` — append `session` to the list for `email`. No return value.

- `Sessions(r map[string][]string, email string) []string` — return the sessions for `email`. If the email is not found, return `nil`.

- `IsRegistered(r map[string][]string, email, session string) bool` — return `true` if `email` is registered for `session`

- `AttendeeCount(r map[string][]string) int` — return the number of unique email addresses

- `Unregister(r map[string][]string, email, session string)` — remove `session` from the list for `email`. If the list becomes empty after removal, delete the key entirely. No return value.

**Hint for Unregister:** Build a new slice that skips the target session, then assign it back. If the new slice is empty, call `delete(r, email)`.

### Run the tests

```bash
go test -v ./stage4/...
```

⏸ Wait for the presenter before moving to Stage 5.

---

## Stage 5 — Structs & Methods

**Session 3** | ~35 minutes

A **struct** groups related fields into one named type.

### Key syntax

```go
// Define a struct
type Speaker struct {
    Name  string
    Title string
}

// Nested struct
type Session struct {
    ID       int
    Title    string
    Speaker  Speaker
    Track    string
    Capacity int
    Enrolled int
}

// Create a value — use field names
s := Session{
    ID:       1,
    Title:    "BGP Routing at Scale",
    Speaker:  Speaker{Name: "Abdi Hassan", Title: "CTO, AfricaNet"},
    Track:    "Network Infrastructure",
    Capacity: 50,
    Enrolled: 12,
}

// Access fields with a dot
fmt.Println(s.Title)         // "BGP Routing at Scale"
fmt.Println(s.Speaker.Name)  // "Abdi Hassan"
fmt.Println(s.Enrolled)      // 12

// Zero value — all fields default to zero
var empty Session
fmt.Println(empty.ID)    // 0
fmt.Println(empty.Title) // ""
```

**Value receiver methods**

A method is a function tied to a type. The receiver is declared before the function name.

```go
// Value receiver: s is a copy. Use for operations that only read the struct.
func (s Session) IsFull() bool {
    return s.Enrolled >= s.Capacity
}

func (s Session) SpotsLeft() int {
    return s.Capacity - s.Enrolled
}

// Call a method with a dot
fmt.Println(s.IsFull())    // true or false
fmt.Println(s.SpotsLeft()) // number
```

### Your task

Create the file `stage5/session.go`. Declare `package stage5` at the top.

Define these two types exactly:

```go
type Speaker struct {
    Name  string
    Title string
}

type Session struct {
    ID       int
    Title    string
    Speaker  Speaker
    Track    string
    Capacity int
    Enrolled int
}
```

Implement these functions and methods:

- `NewSession(id int, title, track string, capacity int) Session` — return a Session with those four fields set and `Enrolled` at `0`. The `Speaker` field is left at its zero value — the caller sets it separately.

- `(s Session) IsFull() bool` — return `true` when `Enrolled >= Capacity`

- `(s Session) SpotsLeft() int` — return `Capacity - Enrolled`. Never return a negative number.

- `(s Session) Summary() string` — return a string containing the session title, the speaker's name, and the number of spots left. The tests check that those substrings are present — exact format is up to you.

### Run the tests

```bash
go test -v ./stage5/...
```

⏸ Wait for the presenter before moving to Stage 6.

---

## Stage 6 — Pointers, Multiple Returns & Errors

**Session 3** | ~35 minutes

### Key syntax

**Pointers**

A pointer holds the memory address of a value. Pass a pointer when a function needs to modify the original.

```go
// *Session means "pointer to a Session"
//  &s     means "address of s"

s   := Session{Capacity: 30}
ptr := &s           // ptr is *Session — points to s
ptr.Capacity = 40   // modifies s directly

// A function that takes *Session can modify the original
func AddAttendee(s *Session, email string) {
    s.Registered = append(s.Registered, email)
}

session := &Session{ID: 1, Capacity: 30}
AddAttendee(session, "abdi@somnog.so")
// session.Registered now contains "abdi@somnog.so"
```

**Multiple return values**

```go
func FindSession(sessions []*Session, id int) (*Session, error) {
    for _, s := range sessions {
        if s.ID == id {
            return s, nil                           // found
        }
    }
    return nil, errors.New("session not found")    // not found
}

// Callers handle both values
s, err := FindSession(all, 42)
if err != nil {
    fmt.Println("error:", err)
    return
}
fmt.Println(s.Title)
```

**errors package**

```go
import "errors"

return errors.New("session is full")
```

**strings package**

```go
import "strings"

strings.Contains("abdi@somnog.so", "@")      // true
strings.ContainsAny("ayan @somnog.so", " ")  // true — any char in second arg
strings.TrimSpace("  hello  ")               // "hello"
```

### Your task

Create the file `stage6/functions.go`. Declare `package stage6` at the top.

Define these types (each stage is self-contained):

```go
type Speaker struct {
    Name  string
    Title string
}

type Session struct {
    ID         int
    Title      string
    Speaker    Speaker
    Track      string
    Capacity   int
    Registered []string   // email addresses of enrolled attendees
}
```

Implement these four functions:

- `Enroll(s *Session, email string) error` — append `email` to `s.Registered`. Return an error if the session is full (`len(s.Registered) >= s.Capacity`) or if `email` is already registered.

- `IsEnrolled(s *Session, email string) bool` — return `true` if `email` is in `s.Registered`

- `FindSession(sessions []*Session, id int) (*Session, error)` — return a pointer to the session with matching `id`. Return `nil` and an error if not found.

- `ValidateEmail(email string) error` — return an error if `email` is empty, contains no `"@"`, or contains whitespace. Return `nil` on success.

### Run the tests

```bash
go test -v ./stage6/...
```

⏸ Wait for the presenter before moving to Stage 7.

---

## Stage 7 — Interfaces

**Session 3** | ~25 minutes

An **interface** declares a set of method signatures. Any type that has those methods satisfies the interface automatically — no `implements` keyword.

### Key syntax

```go
// Declare an interface
type Describable interface {
    Describe() string
}

// Speaker satisfies Describable automatically
type Speaker struct{ Name, Title string }

func (sp Speaker) Describe() string {
    return fmt.Sprintf("Speaker: %s — %s", sp.Name, sp.Title)
}

// Session also satisfies Describable
type Session struct{ ID int; Title string }

func (s Session) Describe() string {
    return fmt.Sprintf("Session %d: %s", s.ID, s.Title)
}

// A function that works with ANY Describable
func PrintAll(items []Describable) {
    for _, item := range items {
        fmt.Println(item.Describe())
    }
}

// Mix different types in one slice
items := []Describable{
    Speaker{Name: "Abdi Hassan", Title: "CTO, AfricaNet"},
    Session{ID: 1, Title: "BGP Routing at Scale"},
}
PrintAll(items)
```

### Your task

Create the file `stage7/interfaces.go`. Declare `package stage7` at the top.

Define this interface:

```go
type Describable interface {
    Describe() string
}
```

Define these three types:

```go
type Speaker struct {
    Name  string
    Title string
}

type Session struct {
    ID    int
    Title string
}

type Workshop struct {
    ID              int
    Title           string
    Instructor      string
    DurationMinutes int
}
```

Implement `Describe() string` on each type:
- `Speaker.Describe()` — must include `Name` and `Title`
- `Session.Describe()` — must include `ID` (as a number) and `Title`
- `Workshop.Describe()` — must include `Title` and `DurationMinutes` (as a number)

Implement these two functions:

- `DescribeAll(items []Describable) []string` — call `Describe()` on each item and return the results as a slice

- `FindByTitle(items []Describable, title string) (Describable, bool)` — return the first item whose `Describe()` output contains `title` (case-insensitive). Return `nil, false` if nothing matches.

**Hint:** `strings.Contains(strings.ToLower(item.Describe()), strings.ToLower(title))`

### Run the tests

```bash
go test -v ./stage7/...
```

⏸ Wait for the presenter before moving to Stage 8.

---

## Stage 8 — Goroutines & Channels

**Session 4** | ~35 minutes

A **goroutine** is a lightweight concurrent function — launch one with `go`. A **channel** is a typed pipe between goroutines.

### Key syntax

```go
// Launch a goroutine
go func() {
    fmt.Println("running concurrently")
}()

// Buffered channel — can hold n values without blocking the sender
ch := make(chan string, 10)

// Send (from a goroutine)
go func() {
    ch <- "hello"
}()

// Receive — blocks until a value arrives
msg := <-ch

// WaitGroup — wait for N goroutines to finish
import "sync"

var wg sync.WaitGroup
emails := []string{"abdi@somnog.so", "layla@somnog.so"}

for _, email := range emails {
    wg.Add(1)              // one more goroutine starting
    go func(e string) {
        defer wg.Done()    // this goroutine is done
        fmt.Println("sending to", e)
    }(email)               // pass email as argument — do not close over the loop variable
}

wg.Wait()  // blocks until all Done() calls complete
```

**Fan-out pattern — launch N goroutines, collect results**

```go
ch := make(chan string, len(emails))  // buffered so goroutines never block

for _, email := range emails {
    wg.Add(1)
    go func(e string) {
        defer wg.Done()
        ch <- process(e)
    }(email)
}

wg.Wait()
close(ch)  // safe to close after all senders have called Done

for result := range ch {
    results = append(results, result)
}
```

### Your task

Create the file `stage8/notify.go`. Declare `package stage8` at the top.

Implement these three functions:

- `NotifyOne(email, message string) string` — return `"sent: " + email`. No concurrency here — this is the unit of work.

- `NotifyAll(emails []string, message string) []string` — call `NotifyOne` for every email **concurrently** using goroutines and a buffered channel. Return all results as a slice. Order does not matter — every email must appear exactly once.

- `FetchAll(urls []string) []string` — for each URL, return `"fetched: " + url`. Run all fetches concurrently. Every URL must appear exactly once.

Import `"sync"` for `sync.WaitGroup`.

### Run the tests

```bash
go test -v ./stage8/...
```

⏸ Wait for the presenter before moving to Stage 9.

---

## Stage 9 — HTTP Server

**Session 4** | ~50 minutes

This is a **code-along**. The skeleton file `stage9/server.go` is already in the repo. Fill in the `// TODO` blocks during the session.

### Key syntax

```go
import (
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

// Register handlers — trailing slash catches /sessions/1, /sessions/1/attendees
mux := http.NewServeMux()
mux.HandleFunc("/sessions", handleAll)
mux.HandleFunc("/sessions/", handleOne)
http.ListenAndServe(":8080", mux)

// Write a JSON response
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(myData)

// Read JSON from the request body
var s Session
json.NewDecoder(r.Body).Decode(&s)

// Set a status code — BEFORE writing the body; 200 OK is the default
w.WriteHeader(http.StatusCreated)  // 201
w.WriteHeader(http.StatusNotFound) // 404

// Check the HTTP method
switch r.Method {
case http.MethodGet:
    // handle GET
case http.MethodPost:
    // handle POST
default:
    http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
}

// Parse a number from a URL segment
id, err := strconv.Atoi("42")
```

**JSON tags** — control the key names in JSON output:

```go
type Session struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}
```

### Your task

Open `stage9/server.go`. Fill in every `// TODO` block:

1. **`GET /sessions`** — write `store` as JSON. Set `Content-Type: application/json`.
2. **`POST /sessions`** — decode a Session from `r.Body`. Assign `ID = len(store) + 1`. Append to `store`. Respond 201 with the created session as JSON.
3. **`GET /sessions/{id}/attendees`** — find the session by ID; if not found respond 404, otherwise write its `Attendees` slice as JSON.

`GET /sessions/{id}` is already implemented — read it before you start.

### Run the tests

```bash
go test -v ./stage9/...
```

### Run the server

```bash
go run ./stage9/
```

Test with curl:

```bash
curl http://localhost:8080/sessions
curl http://localhost:8080/sessions/1
curl http://localhost:8080/sessions/1/attendees
curl -X POST http://localhost:8080/sessions \
  -H "Content-Type: application/json" \
  -d '{"title":"IPv6 at the Edge","speaker":{"name":"Nimo Abdi","title":"Infrastructure Lead"},"track":"Network Infrastructure","capacity":40}'
```

⏸ End of workshop. The full API is running.

---

## Quick Reference

### Variables

```go
var name string = "SomNOG9"  // explicit
count := 4                    // inferred (inside functions)
const max = 200               // constant
```

### Types

```go
string   // "SomNOG9"
int      // 42
bool     // true / false
```

### Composite types

```go
[]string              // slice
map[string][]string   // map
```

### Functions

```go
func Name() string { return "SomNOG9" }
func Name(a, b int) int { return a + b }
func Name(a int) (int, error) { return a, nil }
```

### Control flow

```go
if x > 0 { } else { }
switch x { case 1: ... default: ... }
for i := 0; i < n; i++ { }
for _, v := range slice { }
for k, v := range myMap { }
```

### Structs & methods

```go
type Session struct { ID int; Title string }
func (s Session) IsFull() bool { ... }    // value receiver — read only
```

### Pointers

```go
ptr := &session         // *Session — address of session
func Enroll(s *Session, email string) error { s.Registered = append(...) }
```

### Interfaces

```go
type Describable interface { Describe() string }
// Any type with Describe() string satisfies this — no declaration needed
```

### Goroutines & channels

```go
go func() { ... }()
ch := make(chan string, n)
ch <- value
value := <-ch
var wg sync.WaitGroup
wg.Add(1); go func() { defer wg.Done(); ... }(); wg.Wait()
```

### Errors

```go
import "errors"
return errors.New("something went wrong")
if err != nil { return err }
```

### fmt

```go
fmt.Println("text")
fmt.Printf("name: %s count: %d\n", name, count)
s := fmt.Sprintf("Hello, %s!", name)
```

### strings

```go
strings.Contains(s, sub)
strings.ContainsAny(s, chars)
strings.TrimSpace(s)
strings.ToLower(s)
strings.HasPrefix(s, prefix)
```
