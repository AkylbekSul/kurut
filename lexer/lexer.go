package lexer

import (
	"bufio"
	"io"
	"unicode"
)

type Position struct {
	Line   int
	Column int
}

type Lexer struct {
	position Position
	reader   *bufio.Reader
}

func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		position: Position{Line: 1, Column: 0},
		reader:   bufio.NewReader(reader),
	}
}

func (l *Lexer) Lex() (Position, Token, string) {
	for {
		c, _, err := l.reader.ReadRune()
		if err != nil {
			return l.position, EOF, ""
		}
		l.position.Column++
		switch c {
		case '\n':
			l.position.Line++
			l.position.Column = 0
		case ' ':
			continue
		case '{', '}':
			return l.position, SCOPE, string(c)
		case '(', ')':
			return l.position, FUNC, string(c)
		case '"':
			return l.position, STRING, l.readString()
		case '+':
			return l.position, PLUS, string(c)
		case '-':
			return l.position, MINUS, string(c)
		case '*':
			return l.position, MUL, string(c)
		case '/':
			return l.position, DIV, string(c)
		case '<', '>':
			return l.position, COMPARISON, string(c)
		case '=':
			cc, _, err := l.reader.ReadRune()
			if err != nil {
				panic(err)
			}
			if cc == '=' {
				return l.position, COMPARISON, "=="
			}
			l.backup()
			return l.position, ASSIGN, string(c)
		case '\t':
			continue
		case '\r':
			continue
		case '\f':
			continue
		case '\v':
			continue
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.reader.UnreadRune()
			return l.position, INT, l.readNumber()
		default:
			if unicode.IsLetter(c) {
				l.backup()

				word := l.readIdent()
				switch word {
				case "print":
					return l.position, KEYWORD, word
				case "if":
					return l.position, KEYWORD, word
				case "for":
					return l.position, KEYWORD, word
				case "return":
					return l.position, KEYWORD, word
				default:
					return l.position, IDENT, word
				}
			} else {
				return l.position, ILLEGAL, string(c)
			}
		}
	}
}

func (l *Lexer) backup() error {
	if err := l.reader.UnreadRune(); err != nil {
		return err
	}
	l.position.Column--
	return nil
}

func (l *Lexer) readString() string {
	var buf []rune
	for {
		c, _, err := l.reader.ReadRune()
		if err != nil {
			break
		}
		if c == '"' {
			break
		}
		buf = append(buf, c)
	}
	return string(buf)
}

func (l *Lexer) readNumber() string {
	var buf []rune
	for {
		c, _, err := l.reader.ReadRune()
		if err != nil {
			break
		}
		if !unicode.IsDigit(c) {
			l.reader.UnreadRune()
			break
		}
		buf = append(buf, c)
	}
	return string(buf)
}

func (l *Lexer) readIdent() string {
	var buf []rune
	for {
		c, _, err := l.reader.ReadRune()
		if err != nil {
			break
		}
		if !unicode.IsLetter(c) {
			l.reader.UnreadRune()
			break
		}
		buf = append(buf, c)
	}
	return string(buf)
}

func (l *Lexer) readKeyword() string {
	var buf []rune
	for {
		c, _, err := l.reader.ReadRune()
		if err != nil {
			break
		}
		if !unicode.IsLetter(c) {
			l.reader.UnreadRune()
			break
		}
		buf = append(buf, c)
	}
	return string(buf)
}
