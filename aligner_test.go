package main

import (
	"testing"
    "fmt"
)

func TestAligner(t *testing.T) {
    cases := []struct{
        alignSet string
        input []string
        want []string
    }{
        {
            "{,}",
            []string{
                "{Token::ABC, 5, 50, 'are u'}",
                "{Token::Trash, -1, 0, 'no'}",
            },
            []string{
                "{  Token::ABC,  5, 50, 'are u'}",
                "{Token::Trash, -1,  0,    'no'}",
            },
        },
        {
            ".",
            []string{
                "aaa.aaaa.aaa",
                ".bbbb.bbbbb",
            },
            []string{
                "aaa.aaaa.aaa",
                "   .bbbb.bbbbb",
            },
        },
        {
            "",
            []string{
                "aaa",
                "bbbb",
                "",
            },
            []string{
                "aaa",
                "bbbb",
                "",
            },
        },
    }

    for _, test := range cases {
        t.Run(fmt.Sprint(test.input),
        func(t *testing.T) {

            aligner := NewAligner(test.alignSet)

            for _, line := range test.input {
                aligner.Analyze(line)
            }

            for i := range test.input {
                got := aligner.Transform(test.input[i])
                if got != test.want[i] {
                    t.Errorf("line %d: got %q, want %q", i+1, got, test.want[i])
                }
            }
        })
    }
}
