package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    counts := make(map[string] int)
	input := bufio.NewScanner(os.Stdin)
	// each call to input.Scan() reads the next line and removes the newline char from the end
	for input.Scan() {
		// result can be retrieved by calling input.Text()
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}