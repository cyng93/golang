// exercise 1.3
// TODO - finish the elapsed counting part
package main

import (
	"fmt"
	"os"
    "time"
)

func main() {
    args := os.Args
    start := time.Now()
    for index, arg := range(args) {
        fmt.Printf(" %d\t%v\n", index, arg)
    }
    t := time.Now()
    elapsedSecond := t.Sub(start)
    fmt.Printf("Elapsed: %v\n", elapsedSecond)
}
