// Stage 0 — Setup Verification
// This test confirms your Go environment is correctly configured.
// Run: go test ./stage0/...
// You must see "ok" before moving to Stage 1.
package stage0_test

import (
        "os"
        "os/exec"
        "strconv"
        "strings"
        "testing"
)

// TestGoInstalled checks that the `go` command is available.
func TestGoInstalled(t *testing.T) {
        out, err := exec.Command("go", "version").Output()
        if err != nil {
                t.Fatal("Go is not installed or not on PATH.\n" +
                        "Install it from https://go.dev/dl/ and restart your terminal.")
        }
        version := string(out)
        if !strings.HasPrefix(version, "go version go") {
                t.Fatalf("Unexpected output from `go version`:\n%s", version)
        }
        t.Logf("Go found: %s", strings.TrimSpace(version))
}

// TestGoVersionRecent checks that Go 1.21 or newer is installed.
func TestGoVersionRecent(t *testing.T) {
        out, err := exec.Command("go", "version").Output()
        if err != nil {
                t.Fatal("Cannot run `go version`.")
        }
        // Output looks like: "go version go1.23.1 linux/amd64"
        fields := strings.Fields(strings.TrimSpace(string(out)))
        if len(fields) < 3 {
                t.Fatalf("Cannot parse Go version output: %s", string(out))
        }
        // fields[2] = "go1.23.1"
        numStr := strings.TrimPrefix(fields[2], "go") // "1.23.1"
        parts := strings.SplitN(numStr, ".", 3)
        if len(parts) < 2 {
                t.Fatalf("Cannot parse version number: %s", numStr)
        }
        major, _ := strconv.Atoi(parts[0])
        minor, _ := strconv.Atoi(parts[1])

        if major < 1 || (major == 1 && minor < 21) {
                t.Fatalf("Go version too old: %s\n"+
                        "This workshop requires Go 1.21 or newer.\n"+
                        "Download the latest from https://go.dev/dl/", strings.TrimSpace(string(out)))
        }
        t.Logf("Go version OK: %s", strings.TrimSpace(string(out)))
}

// TestGoModExists checks that go.mod exists in the project root.
func TestGoModExists(t *testing.T) {
        // go test runs from the package directory (stage0/),
        // so go.mod is one level up.
        _, err := os.ReadFile("../go.mod")
        if err != nil {
                t.Fatal("go.mod not found in project root.\n" +
                        "Run: go mod init github.com/somnog/learn-go")
        }
        t.Log("go.mod found.")
}

// TestModuleName checks that the module is named correctly.
func TestModuleName(t *testing.T) {
        data, err := os.ReadFile("../go.mod")
        if err != nil {
                t.Fatal("go.mod not found — run TestGoModExists to diagnose.")
        }
        content := string(data)
        expected := "module github.com/somnog/learn-go"
        if !strings.Contains(content, expected) {
                t.Fatalf("Wrong module name in go.mod.\n"+
                        "Expected:  %s\n"+
                        "Found:\n%s\n\n"+
                        "Fix: delete go.mod and run: go mod init github.com/somnog/learn-go",
                        expected, content)
        }
        t.Log("Module name is correct.")
}
