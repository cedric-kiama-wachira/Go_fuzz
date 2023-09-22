// Steps: 
// 1. Added the main package declaration.
// 2. Wrote the package declaration.
// 3. Included main function to initialize a string, reverse it, print the output, and repeat. 
// 4. The main function uses the fmt package, so I will need to import it.
// 5. Replacing the Reverse function so that we can traverse the string by runes, instead of by bytes.
// 6. Replaced the Reverse function from above step so that I can understand what is going wrong when converting the string to a slice of runes.
// 7. Replaced the Reverse function after knowing what the exact issue is; the input in reverse is not a valid UTF-8
// 8. Replaced the Main fucntion in Step 3 to discard the extra error value. Now These calls to Reverse should return a nil error, since the input string is valid UTF-8.
// 9. Modified the import statement to also include errors and the unicode/utf8 packages

// Step 1.
package main

// Step 4.
import (
        "errors"
	"fmt"
	"unicode/utf8"
)
// Step 3.
//func main() {
//    input := "The quick brown fox jumped over the lazy dog"
//    rev := Reverse(input)
//    doubleRev := Reverse(rev)
//    fmt.Printf("original: %q\n", input)
//    fmt.Printf("reversed: %q\n", rev)
//    fmt.Printf("reversed again: %q\n", doubleRev)
//}

// Step 8.
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev, revErr := Reverse(input)
    doubleRev, doubleRevErr := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q, err: %v\n", rev, revErr)
    fmt.Printf("reversed again: %q, err: %v\n", doubleRev, doubleRevErr)
}

// Step 2.
//func Reverse(s string) string {
//    b := []byte(s)
//    for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
//        b[i], b[j] = b[j], b[i]
//    }
//    return string(b)
//}

// Step 5.
//func Reverse(s string) string {
//    r := []rune(s)
//    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//        r[i], r[j] = r[j], r[i]
//    }
//    return string(r)
//}

// Step 6.
//func Reverse(s string) string {
//    fmt.Printf("input: %q\n", s)
//    r := []rune(s)
//    fmt.Printf("runes: %q\n", r)
//    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//        r[i], r[j] = r[j], r[i]
//    }
//    return string(r)
//}

// Step 7.
func Reverse(s string) (string, error) {
    if !utf8.ValidString(s) {
        return s, errors.New("input is not valid UTF-8")
    }
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r), nil
}

