package lexer

type (
	Token   int
	Keyword int

	// TODO: create a better name for this struct
	Tokens struct {
		Position Position
		Token    Token
		Value    string
	}
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

	COMPARISON

	SCOPE // {}

	FUNC // ()

	STRING

	KEYWORD

	PRINT Keyword = iota
)

var tokens = map[Token]string{
	EOF:        "EOF",
	ILLEGAL:    "ILLEGAL",
	IDENT:      "IDENT",
	INT:        "INT",
	ASSIGN:     "=",
	PLUS:       "+",
	MINUS:      "-",
	MUL:        "*",
	DIV:        "/",
	COMPARISON: "COMPARE",
	KEYWORD:    "KEYWORD",
	SCOPE:      "{}",
	FUNC:       "()",
	STRING:     "STRING",
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
