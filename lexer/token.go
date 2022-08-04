package lexer

type (
	Token int
	Keyword int
)

const (
	EOF Token = iota
	ILLEGAL
	IDENT
	INT
	ASSIGN

	PLUS
	MINUS
	MUL
	DIV

	KEYWORD

	PRINT Keyword = iota
)

var tokens = map[Token]string{
	EOF: "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT: "IDENT",
	INT: "INT",
	ASSIGN: "=",
	PLUS: "+",
	MINUS: "-",
	MUL: "*",
	DIV: "/",
	KEYWORD: "KEYWORD",
}

var Keywords = map[Keyword]string{
	PRINT: "print",
}


func (t Token) String() string {
	v, ok := tokens[t]
	if !ok {
		return tokens[ILLEGAL]
	}
	return v
}