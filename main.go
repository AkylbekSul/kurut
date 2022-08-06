package main

import (
	"fmt"
	"os"

	"github.com/rasulov-emirlan/kurut/lexer"
	"github.com/rasulov-emirlan/kurut/parser"
)

func main() {
	file, err := os.Open("testing/example.kut")
	if err != nil {
		panic(err)
	}

	l := lexer.NewLexer(file)
	tokens := parser.Tokens{}
	for {
		position, token, value := l.Lex()
		if token == lexer.EOF {
			break
		}
		tokens.Positions = append(tokens.Positions, position)
		tokens.Tokens = append(tokens.Tokens, token)
		tokens.Values = append(tokens.Values, value)
		fmt.Printf("%d:%d\t%s\t%s\n", position.Line, position.Column, token, value)
	}
}
