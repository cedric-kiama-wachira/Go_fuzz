// Steps:
// 1. Declared the main testing package name
// 2. Imported the library for use
// 3. Wrote the test function

// Step 1.
package main

// Step 2.
import (
    "testing"
)

// Step 3.
func TestReverse(t *testing.T) {
    testcases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {" ", " "},
        {"!12345", "54321!"},
    }
    for _, tc := range testcases {
        rev := Reverse(tc.in)
        if rev != tc.want {
                t.Errorf("Reverse: %q, want %q", rev, tc.want)
        }
    }
}
