package main

import "fmt"

func main() {
	// var const 👍 = true?
	// var var 1️⃣ = 1!
	// name = "New"!
	//  if (1 == 1) { print("Valid")! }
	// const const name = "Luke"!
	code := `
    var const 👍 = true?
	`

	lexer := NewLexer(code)
	for tok := lexer.NextToken(); tok.Tok != EOF; tok = lexer.NextToken() {
		fmt.Printf("Token: { Tok:%s, Lit:`%s` }\n", tok.Tok, tok.Lit)
	}
}
