package main

type TokenType string

type Token struct {
	Tok TokenType
	Lit string
}

const (
	EndStmt          TokenType = "EndStmt"         // For '!' repeated
	EndStmtDbg                 = "EndStmtDbg"      // For '?' (debug info) repeated
	Not                        = "Not"             // For ';' (false / Not)
	Illegal                    = "[*]"             // For illegal tokens
	EOF                        = "EOF"             // End of File
	Fn                         = "Fn"              // For 'fn' keyword
	VName                      = "var"             // Variable & function names
	VConstConstConst           = "ConstConstConst" // Immutable data
	VConstConst                = "ConstConst"      // Type can be reassigned
	VConstVar                  = "ConstVar"        // Var can be edited but not reassigned
	VVarConst                  = "VarConst"        // Var cannot be edited but reassigned
	If                         = "If"              // If conditional
	Else                       = "Else"            // Else conditional
	True                       = "true"            // True bool
	False                      = "false"           // False bool
	Maybe                      = "maybe"           // Maybe bool
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
	v := lookupTable[tok]
	return v
}

var lookupTable = map[string]TokenType{
	"fn": Fn,
	"!":  EndStmt,
	";":  Not,
}
