package runtime

import (
	"fmt"
	"io"
	"strconv"

	"github.com/rasulov-emirlan/kurut/lexer"
)

func Run(file io.Reader) error {
	l := lexer.NewLexer(file)
	var (
		positions []lexer.Position
		tokens    []lexer.Token
		values    []string
	)

	for {
		position, token, value := l.Lex()
		if token == lexer.EOF {
			break
		}
		positions = append(positions, position)
		tokens = append(tokens, token)
		values = append(values, value)
	}

	for i := 0; i < len(positions); i++ {
		fmt.Printf("%d:%d\t%s\t%s\n", positions[i].Line, positions[i].Column, tokens[i], values[i])
	}

	var (
		integers = map[string]int{}

		operationsBuf = 0

		isAsigning bool
		isPlus     bool
		isMul 	   bool
		lastIDENT  string
	)

	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case lexer.IDENT:
			integers[values[i]] = 1
			lastIDENT = values[i]
			continue
		case lexer.INT:
			if isAsigning {
				n, _ := strconv.Atoi(values[i])
				integers[lastIDENT] = n
				continue
			}
			if isPlus {
				n, _ := strconv.Atoi(values[i])
				operationsBuf += n
				continue
			}
			if isMul {
				n, _ := strconv.Atoi(values[i])
				operationsBuf *= n
				continue
			}
			continue
		case lexer.ASSIGN:
			isAsigning = true
			continue
		case lexer.PLUS:
			isPlus = true
			continue
		case lexer.MUL:
			isMul = true
			continue
		case lexer.KEYWORD:
			if values[i] == "print" {
				for i < len(tokens) {
					if tokens[i] == lexer.IDENT {
						fmt.Println(integers[values[i]])
					}
					if tokens[i] == lexer.INT {
						fmt.Println(values[i])
					}
					i++
				}
			}
		}
	}
	return nil
}