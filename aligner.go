package main

import "strings"

type Aligner struct {
    alignSet string

    columns []int
}

func NewAligner(alignSet string) *Aligner {
    return &Aligner{alignSet: alignSet}
}

func (a *Aligner) Analyze(line string) {
    var lineCols []int

    lastPos := -1 // since last pos is < current pos
    for li, r := range line {
        if strings.ContainsRune(a.alignSet, r) {
            lineCols = append(lineCols, li - lastPos)
            lastPos = li
        }
    }

    ci := 0
    for ; ci < len(a.columns) && ci < len(lineCols); ci++ {
        if a.columns[ci] < lineCols[ci] {
            a.columns[ci] = lineCols[ci]
        }
    }

    if ci < len(lineCols) {
        a.columns = append(a.columns, lineCols[ci:]...)
    }
}

func (a *Aligner) Transform(line string) string {
    var str strings.Builder

    lastPos := -1
    li, ci := 0, 0
    for ; li < len(line) && ci < len(a.columns); li++ {
        if strings.ContainsRune(a.alignSet, rune(line[li])) {
            colLen := li - lastPos
            str.WriteString(strings.Repeat(" ", a.columns[ci] - colLen))
            str.WriteString(line[lastPos+1:li+1])
            ci++
            lastPos = li
        }
    }

    if li < len(line) {
        str.WriteString(line[li:])
    }

    return str.String()
}


