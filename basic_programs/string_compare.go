// String Compare function in Go language returns integer value -1, 0 or 1 after comparing two strings lexicographically. The result will be 0 if S1==S2, -1 if S1 < S2, and +1 if S1 > S2. 

package main
 
import (
    "fmt"
    "strings"
)
 
func main() {
    fmt.Println(strings.Compare("India", "Indiana"))
    fmt.Println(strings.Compare("Indiana", "India"))
    fmt.Println(strings.Compare("India", "India"))
}