// Steps: 
// 1. Added the main package declaration.
// 2. Wrote the package declaration.
// 3. Included main function to initialize a string, reverse it, print the output, and repeat. 
// 4. The main function uses the fmt package, so I will need to import it.
// 5. Replacing the Reverse function so that we can traverse the string by runes, instead of by bytes.


// Step 1.
package main

// Step 4.
import "fmt"

// Step 3.
func main() {
    input := "The quick brown fox jumped over the lazy dog"
    rev := Reverse(input)
    doubleRev := Reverse(rev)
    fmt.Printf("original: %q\n", input)
    fmt.Printf("reversed: %q\n", rev)
    fmt.Printf("reversed again: %q\n", doubleRev)
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
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}	

