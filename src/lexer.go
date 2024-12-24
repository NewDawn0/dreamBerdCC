package main

import "strings"

type Lexer struct {
	input string
	pos   uint
	rpos  uint
	ch    rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.rpos >= uint(len(l.input)) {
		l.ch = 0 // EOF
	} else {
		l.ch = rune(l.input[l.rpos])
	}
	l.pos = l.rpos
	l.rpos++
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	var tok Token
	// Skip over `(` `)` Since they don't matter
	if inArray(l.ch, "()") {
		l.readChar()
	}
	switch {
	case l.ch == 0:
		return newToken(EOF, "")
	case isLetter(l.ch):
		return l.readLetter()
	case inTable(string(l.ch)):
		if inArray(l.ch, "!?") {
			return l.readEndStmt(l.ch)
		}
		tok = newToken(getItem(string(l.ch)), string(l.ch))
		l.readChar()
	default:
		tok = newToken(Illegal, string(l.ch))
		l.readChar() // Consume illegal char
	}
	return tok
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

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func inArray(ch rune, arr string) bool {
	for _, val := range strings.Split(arr, "") {
		if string(ch) == val {
			return true
		}
	}
	return false
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}
