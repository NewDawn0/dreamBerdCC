package main

import "fmt"

func main() {
	// fn print("Hello world")?
	// if (;false) {
	//   print("hello world")
	// }
	code := `
    fn print("Hello world")!
	  fn print("Hello world")!!!
  `
	l := NewLexer(code)
	for tok := l.NextToken(); tok.Tok != EOF; tok = l.NextToken() {
		fmt.Printf("Token: %v\n", tok)
	}
}
