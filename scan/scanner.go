package scan

import "fmt"

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
	if !scan.isAtEnd() {
		return 0
	}

	return scan.source[scan.current]
}

func (scan *Scanner) addToken(tokentype TokenType, literal any) {
	text := scan.source[scan.start:scan.current]
	scan.tokenList = append(scan.tokenList, NewToken(tokentype, text, literal, scan.line))
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
				}
				if c == '*' && scan.peek() == '/' {
					depth--
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
	
	case '\n':
		scan.line++

	default:
		return fmt.Errorf("Unrecognized lexeme %c in line %d", c, scan.line)
	}
	return nil
}
