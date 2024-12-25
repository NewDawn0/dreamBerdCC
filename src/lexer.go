package main

import (
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	pos          uint
	rpos         uint
	ch           rune
	nextToken    Token
	currentToken Token
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	l.nextToken = l.tokenize()
	return l
}

func (l *Lexer) readChar() {
	if l.rpos >= uint(len(l.input)) {
		l.ch = 0 // EOF
	} else {
		r, size := utf8.DecodeRuneInString(l.input[l.rpos:])
		l.ch = r
		l.pos = l.rpos
		l.rpos += uint(size)
	}
}

func (l *Lexer) tokenize() Token {
	var tok Token
	// Skip over various useless chars like `(`, `)`, etc. because they don't matter
	for strings.Contains("() \t\n\r", string(l.ch)) {
		l.readChar()
	}
	switch {
	case l.ch == 0:
		tok = newToken(EOF, "")
	case isLetter(l.ch):
		tok = l.readLetter()
	case inTable(string(l.ch)):
		if strings.Contains("!?", string(l.ch)) {
			tok = l.readEndStmt(l.ch)
		} else {
			tok = newToken(getItem(string(l.ch)), string(l.ch))
			l.readChar()
		}
	default:
		tok = newToken(Illegal, string(l.ch))
		l.readChar() // Consume illegal char
	}
	return tok
}

func (l *Lexer) NextToken() Token {
	l.currentToken = l.nextToken
	l.nextToken = l.tokenize()
	return l.currentToken
}

func (l *Lexer) readLetter() Token {
	var accum string
	var tok Token
	for isLetter(l.ch) {
		accum += string(l.ch)
		l.readChar()
	}
	if inTable(accum) {
		tok = newToken(getItem(accum), accum)
	} else {
		tok = newToken(Illegal, accum)
	}
	return tok
}

func (l *Lexer) readEndStmt(ch rune) Token {
	var accum string
	var tok Token
	for l.ch == ch {
		accum += string(l.ch)
		l.readChar()
	}
	if inTable(string(ch)) {
		tok = newToken(getItem(string(ch)), accum)
	} else {
		tok = newToken(Illegal, accum)
	}
	return tok
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
