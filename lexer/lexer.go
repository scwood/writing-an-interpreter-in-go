package lexer

import "github.com/scwood/writing-an-interpreter-in-go/token"

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	switch l.ch {
	case '+':
		tok = newToken(token.Plus, l.ch)
	case '-':
		tok = newToken(token.Minus, l.ch)
	case '/':
		tok = newToken(token.Slash, l.ch)
	case '*':
		tok = newToken(token.Asterisk, l.ch)
	case '<':
		tok = newToken(token.LessThan, l.ch)
	case '>':
		tok = newToken(token.GreaterThan, l.ch)
	case ';':
		tok = newToken(token.Semicolon, l.ch)
	case ',':
		tok = newToken(token.Comma, l.ch)
	case '(':
		tok = newToken(token.LeftParen, l.ch)
	case ')':
		tok = newToken(token.RightParen, l.ch)
	case '{':
		tok = newToken(token.LeftBrace, l.ch)
	case '}':
		tok = newToken(token.RightBrace, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.Equal, Literal: "=="}
		} else {
			tok = newToken(token.Assign, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			tok = token.Token{Type: token.NotEqual, Literal: "!="}
		} else {
			tok = newToken(token.Bang, l.ch)
		}
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readSection(isLetter)
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readSection(isDigit)
			tok.Type = token.Int
			return tok
		} else {
			tok = newToken(token.Illegal, l.ch)
		}
	}
	l.readChar()
	return tok
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' ||
		ch == '_' || ch == '?' || ch == '!'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readSection(predicate func(byte) bool) string {
	start := l.position
	for predicate(l.ch) {
		l.readChar()
	}
	return l.input[start:l.position]
}
