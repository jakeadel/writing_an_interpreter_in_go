package lexer

import (
	"fmt"
	"monkey/token"
)

type Lexer struct {
	input        string
	position     int  // The current token
	readPosition int  // A subsequent token (to peek)
	ch           byte // current char we're looking at
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) peekChar() byte {
	if (l.readPosition >= len(l.input)) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()
	fmt.Println(l.position, string(l.ch))

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "=="
			tok.Type = token.EQ
		} else {
			tok = newToken((token.ASSIGN), l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok.Literal = "!="
			tok.Type = token.NOT_EQ
		} else {
			tok = newToken((token.BANG), l.ch)
		}
	case ';':
		tok = newToken((token.SEMICOLON), l.ch)
	case '(':
		tok = newToken((token.LPAREN), l.ch)
	case ')':
		tok = newToken((token.RPAREN), l.ch)
	case '{':
		tok = newToken((token.LBRACE), l.ch)
	case '}':
		tok = newToken((token.RBRACE), l.ch)
	case ',':
		tok = newToken((token.COMMA), l.ch)
	case '+':
		tok = newToken((token.PLUS), l.ch)
	case '-':
		tok = newToken((token.MINUS), l.ch)
	case '/':
		tok = newToken((token.SLASH), l.ch)
	case '*':
		tok = newToken((token.ASTERISK), l.ch)
	case '<':
		tok = newToken((token.LT), l.ch)
	case '>':
		tok = newToken((token.GT), l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			// Need to read until not a letter to get variable name
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// Return early because readIdentifier advances our position
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(letter byte) bool {
	if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') || letter == '_' {
		return true
	} else {
		return false
	}
}

// Advance until we find a non-letter, return string of letters advanced over
func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func isDigit(char byte) bool {
	if (char >= '0' && char <= '9') {
		return true
	} else {
		return false
	}
}

func (l *Lexer) readNumber() string {
	startPosition := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		fmt.Println(l.ch)
		l.readChar()
	}
}
