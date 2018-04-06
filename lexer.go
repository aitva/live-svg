package main

type Lexer struct {
	input              string
	curentPos, nextPos int
	c                  byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}

func (l *Lexer) readChar() {
	l.c = l.peekChar()
	l.curentPos = l.nextPos
	l.nextPos++
}

func (l *Lexer) readWhitespace() {
	for l.c == ' ' || l.c == '\t' || l.c == '\n' {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	start := l.curentPos
	for isAlphaNum(l.c) {
		l.readChar()
	}
	return l.input[start:l.curentPos]
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.readWhitespace()

	switch l.c {
	case '<':
		tok = Token{Type: TokenOpen, Litteral: string(l.c)}
	case '>':
		tok = Token{Type: TokenClose, Litteral: string(l.c)}
	case '/':
		tok = Token{Type: TokenSlash, Litteral: string(l.c)}
	case '=':
		tok = Token{Type: TokenEqual, Litteral: string(l.c)}
	case '"':
		tok = Token{Type: TokenQuote, Litteral: string(l.c)}
	case 0:
		tok = Token{Type: TokenEOF, Litteral: string(l.c)}
	default:
		if isAlphaNum(l.c) {
			tok = Token{Type: TokenIdentifier, Litteral: l.readIdentifier()}
			return tok
		} else {
			tok = Token{Type: TokenIllegal, Litteral: string(l.c)}
		}
	}

	l.readChar()

	return tok
}

func isAlphaNum(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || '0' <= c && c <= '9' || c == '-'
}
