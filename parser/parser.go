package parser

import "github.com/rasulov-emirlan/kurut/lexer"

type Tokens struct {
	Positions []lexer.Position
	Tokens    []lexer.Token
	Values    []string
}

func Parse(tokens Tokens) {
	// TODO: write a badass parser 😎
}
