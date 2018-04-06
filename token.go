package main

const (
	TokenOpen       TokenType = "<"
	TokenClose      TokenType = ">"
	TokenSlash      TokenType = "/"
	TokenEqual      TokenType = "="
	TokenQuote      TokenType = "\""
	TokenIdentifier TokenType = "IDENT"
	TokenEOF        TokenType = "EOF"
	TokenIllegal    TokenType = "ILLEGAL"
)

type TokenType string

type Token struct {
	Type     TokenType
	Litteral string
}
