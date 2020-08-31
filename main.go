package main

import (
    "os"
    "fmt"
    "bufio"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    alignSet := ""
    if len(os.Args) > 1 {
        alignSet = os.Args[1]
    }

    aligner := NewAligner(alignSet)
    for i := range lines {
        aligner.Analyze(lines[i])
    }

    for i := range lines {
        fmt.Println(aligner.Transform(lines[i]))
    }
}
