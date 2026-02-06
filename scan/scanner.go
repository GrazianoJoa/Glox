package scan

import (
	"fmt"
	"strconv"
)

type Scanner struct {
	source string
	tokenList []*Token
	start, current, line int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		tokenList: []*Token{},
		start: 0,
		current: 0,
		line: 1,
	}
}

func (scan Scanner) isAtEnd() bool {
	return scan.current >= len(scan.source)
}

func (scan *Scanner) advance() byte {
	c := scan.source[scan.current]
	scan.current++
	return c
}

func (scan *Scanner) match(expected byte) bool {
	if scan.isAtEnd() { 
		return false
	}

	if scan.source[scan.current] != expected { 
		return false
	}

	scan.current++
	return true
}

func (scan Scanner) peek() byte {
	if scan.isAtEnd() {
		return '\000'
	}

	return scan.source[scan.current]
}

func (scan Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (scan Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (scan *Scanner) peekNext() byte {
	if scan.current + 1 >= len(scan.source) {
		return '\000'
	}
	return scan.source[scan.current+1]
}

func (scan *Scanner) number() error {
	for scan.isDigit(scan.peek()) {
		scan.advance()
	}

	if scan.peek() == '.' && scan.isDigit(scan.peekNext()) {
		scan.advance()
	}

	for scan.isDigit(scan.peek()) {
		scan.advance()
	}

	value, err := strconv.ParseFloat(scan.source[scan.start:scan.current], 64)
	if err != nil {
		return fmt.Errorf("ERROR")
	}

	scan.addToken(TypeNumber, value)
	return nil
}

func (scan Scanner) isAlphaNumeric(c byte) bool {
	return scan.isDigit(c) || scan.isAlpha(c)
}

func (scan *Scanner) identifier() error {
	for scan.isAlphaNumeric(scan.peek()) {
		scan.advance()
	}

	scan.addToken(TypeIdentifier, "")
	return nil
}

func (scan *Scanner) addToken(tokentype TokenType, literal any) {
	text := scan.source[scan.start:scan.current]
	scan.tokenList = append(scan.tokenList, NewToken(tokentype, text, literal, scan.line))
}

func (scan *Scanner) string() error {
	for scan.peek() != '"' && !scan.isAtEnd() {
		if scan.peek() == '\n' {
			scan.line++
		}
		scan.advance()
	}

	if scan.isAtEnd() {
		return fmt.Errorf("Unclosed string in line %d", scan.line)
	}

	scan.advance()

	scan.addToken(TypeString, scan.source[scan.start+1:scan.current-1])
	return nil
}

func (scan *Scanner) ScanTokens() ([]*Token, error) {
	for !scan.isAtEnd() {
		scan.start = scan.current
		err := scan.scanToken()
		if err != nil {
			return nil, err
		}
	}

	scan.tokenList = append(scan.tokenList, NewToken(TypeEOF, "", nil, scan.line))
	return scan.tokenList, nil
}

func (scan *Scanner) scanToken() error {
	c := scan.advance()
	switch c {
	// Single-character 
	case '(':
		scan.addToken(TypeLeftParen, nil)
	case ')':
		scan.addToken(TypeRightParen, nil)
	case '{':
		scan.addToken(TypeLeftBrace, nil)
	case '}':
		scan.addToken(TypeRightBrace, nil)
	case '+':
		scan.addToken(TypePlus, nil)
	case '-':
		scan.addToken(TypeMinus, nil)
	case '*':
		scan.addToken(TypeStar, nil)
	case '/':
		if scan.match('/') {
			for !scan.isAtEnd() && scan.peek() != '\n' {
				scan.advance()
			}
		} else if scan.match('*') {
			depth := 1
			for !scan.isAtEnd() && depth != 0 {
				c := scan.advance()
				if c == '/' && scan.peek() == '*' {
					depth++
					scan.advance()
				}
				if c == '*' && scan.peek() == '/' {
					depth--
					scan.advance()
				}
			}
			if depth != 0 {
				return fmt.Errorf("Unclosed comment in line: %d", scan.line)
			}
		} else {
			scan.addToken(TypeSlash, nil)
		}
	case ',':
		scan.addToken(TypeComma, nil)
	case '.':
		scan.addToken(TypeDot, nil)
	case ';':
		scan.addToken(TypeSemicolon, nil)

	case '"': 
		return scan.string()

	case ' ':
		break
	case '\t':
		break
	case '\r':
		break
	case '\n':
		scan.line++

	default:
		if scan.isDigit(c) {
			return scan.number()
		} else if scan.isAlpha(c) {
			return scan.identifier()
		} else {
			return fmt.Errorf("Unrecognized token %c in line %d\n", c, scan.line)
		}
	}
	return nil
}
