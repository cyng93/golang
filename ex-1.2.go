// exercise 1.2
package main

import (
	"fmt"
	"os"
)

func main() {
    for index, arg := range(os.Args) {
        fmt.Printf(" %d\t%v\n", index, arg)
    }
}
