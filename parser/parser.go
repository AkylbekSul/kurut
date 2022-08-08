package parser

import "github.com/rasulov-emirlan/kurut/lexer"

func Parse(tokens []lexer.Tokens) (SyntaxTree, error) {
	// TODO: write a badass parser 😎

	var (
		tree       SyntaxTree
		lastIDENT  string
		isVariable bool
	)

	// TODO: god help us 🍀

	for _, token := range tokens {
		switch token.Token {
		case lexer.IDENT:
			lastIDENT = token.Value
			isVariable = true
		case lexer.INT:
			if !isVariable {
				break
			}

			tree.Scopes = append(tree.Scopes, &Scope{
				Identities: []*Identity[any]{
					{
						Name:  lastIDENT,
						Value: token.Value,
					},
				},
			})
		}
	}

	return tree, nil
}
