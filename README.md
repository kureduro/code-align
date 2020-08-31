## Code align

Basically, turn this 
```
{ TokenType::ForLoop, 1, 0, "for"}, { TokenType::Identifier, 1, 4, "i" },
{ TokenType::InRange, 1, 6, "in"}, { TokenType::ReverseRange, 1, 9, "reverse" },
{ TokenType::IntegerLiteral, 1, 17, "1" }, { TokenType::TwoDots, 1, 18, ".." },
{ TokenType::IntegerLiteral, 1, 20, "42" }, { TokenType::LoopBegin, 1, 23, "loop" },
{ TokenType::NewLine, 1, 27, "\n" },
```

into this
```

{       TokenType::ForLoop, 1,  0, "for"}, {  TokenType::Identifier, 1,  4,       "i"},
{       TokenType::InRange, 1,  6,  "in"}, {TokenType::ReverseRange, 1,  9, "reverse"},
{TokenType::IntegerLiteral, 1, 17,   "1"}, {     TokenType::TwoDots, 1, 18,      ".."},
{TokenType::IntegerLiteral, 1, 20,  "42"}, {   TokenType::LoopBegin, 1, 23,    "loop"},
{       TokenType::NewLine, 1, 27,  "\n"},
```

### Building

```
$ go build .
```

It will produce an executable inside the directory.

### Usage

To make any use of it, you need to supply a set of alignment characters. These characters will be vertically aligned in the output. The program reads lines via `stdin`.

```
$ printf "\t{Line one, two}\n\t{three, four}" | ./code-align "{,}"
    {Line one,  two}
    {   three, four}
```

I have come with this hacky workflow: open a text file in vim, copy needed lines there, then type
```
:!cat file | ./code-align "{}," | xclip -selection clipboard
```

This will copy the transformed lines into the clipboard.
