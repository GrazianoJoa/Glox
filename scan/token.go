package scan

import "fmt"

type Token struct {
	tokentype TokenType
	Lexeme string
	literal interface{}
	line int
}

func NewToken(tokentype TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		tokentype: tokentype,
		Lexeme: lexeme,
		literal: literal,
		line: line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %d", t.tokentype.String(), t.Lexeme, t.line)
}


