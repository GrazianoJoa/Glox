package scan

import "fmt"

type Token struct {
	tokentype TokenType
	lexeme string
	literal interface{}
	line int
}

func NewToken(tokentype TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{
		tokentype: tokentype,
		lexeme: lexeme,
		literal: literal,
		line: line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf("%s %s %d", t.tokentype.String(), t.lexeme, t.line)
}


