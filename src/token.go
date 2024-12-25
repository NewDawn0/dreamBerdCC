package main

type TokenType string

type Token struct {
	Tok TokenType
	Lit string
}

const (
	// Statement Endings
	EndStmt    TokenType = "EndStmt"    // For '!' repeated
	EndStmtDbg           = "EndStmtDbg" // For '?' (debug info) repeated

	// Operators and Comparisons
	Not    = "Not"    // For ';'
	Assign = "Assign" // For '='

	// Variable Types
	Const = "Const"
	Var   = "Var"

	// Other Tokens
	If      = "If"
	Else    = "Else"
	True    = "True"
	False   = "False"
	Maybe   = "Maybe"
	EOF     = "EOF"
	Illegal = "[*]"
)

func newToken(t TokenType, lit string) Token {
	return Token{
		Tok: t,
		Lit: lit,
	}
}

func inTable(tok string) bool {
	_, ok := lookupTable[tok]
	return ok
}
func getItem(tok string) TokenType {
	return lookupTable[tok]
}

var lookupTable = map[string]TokenType{
	"if":    If,
	"else":  Else,
	"true":  True,
	"false": False,
	"maybe": Maybe,
	"!":     EndStmt,
	";":     Not,
	"=":     Assign,
	"const": Const,
	"var":   Var,
}
