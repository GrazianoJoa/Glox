package scan

type TokenType int

const (
	// Single character tokens
	TypeLeftParen TokenType = iota
	TypeRightParen
	TypeLeftBrace
	TypeRightBrace
	TypePlus
	TypeMinus
	TypeSlash
	TypeStar
	TypeComma
	TypeDot
	TypeSemicolon

	// Two or more character tokens
	TypeBang
	TypeBangEqual
	TypeEqual
	TypeEqualEqual
	TypeGreater
	TypeGreaterEqual
	TypeLess
	TypeLessEqual

	// Literals
	TypeIdentifier
	TypeString
	TypeNumber

	// Keywords
	TypeIf
	TypeElse
	TypeTrue
	TypeFalse
	TypeAnd
	TypeOr
	TypeNil

	TypeVar
	TypeFor
	TypeWhile

	TypeFun
	TypeReturn
	TypePrint
	TypeClass
	TypeSuper
	TypeThis

	TypeEOF
)

var typeName = []string {
	"LEFT_PAREN",
	"RIGHT_PAREN",
	"LEFT_BRACE",
	"RIGHT_BRACE",
	"PLUS",
	"MINUS",
	"SLASH",
	"STAR",
	"COMMA",
	"DOT",
	"SEMICOLON",

	"BANG",
	"BANG_EQUAL",
	"EQUAL",
	"EQUAL_EQUAL",
	"GREATER",
	"GREATER_EQUAL",
	"LESS",
	"LESS_EQUAL",

	"IDENTIFIER",
	"STRING",
	"NUMBER",

	"IF",
	"ELSE",
	"TRUE",
	"FALSE",
	"AND",
	"OR",
	"NIL",

	"VAR",
	"FOR",
	"WHILE",

	"FUN",
	"RETURN",
	"PRINT",
	"CLASS",
	"SUPER",
	"THIS",

	"EOF",
}

func (tokentype TokenType) String() string {
	return typeName[tokentype]
}
