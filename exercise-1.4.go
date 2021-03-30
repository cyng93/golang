// dup2 prints the count and text of lines that appear more than once in the
// input. It reads from the stdin or from a list of named files
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    counts := make(map[string]string)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts, "stdin")
    } else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts, arg)
            f.Close()
        }
    }
    for line, filename := range counts {
        if strings.Contains(filename, ";") {
            fmt.Printf("%s\t%s\n", line, filename)
        }
        //if n > 1 {
        //    fmt.Printf("%d\t%s\n", n, line)
        //}
    }
}

func countLines(f *os.File, counts map[string]string, filename string) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        if counts[input.Text()] != "" {
            counts[input.Text()] += ";"
        }
        counts[input.Text()] += filename
    }
    // NOTE: ignoring potential errors from input.Err()
}
