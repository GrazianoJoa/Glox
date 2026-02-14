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

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current++
	return c
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() { 
		return false
	}

	if s.source[s.current] != expected { 
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}

	return s.source[s.current]
}

func (s *Scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func (s *Scanner) peekNext() byte {
	if s.current + 1 >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

func (s *Scanner) number() error {
	for s.isDigit(s.peek()) {
		s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		s.advance()
	}

	for s.isDigit(s.peek()) {
		s.advance()
	}

	value, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return fmt.Errorf("invalid number at line %d: %w", s.line, err)
	}

	s.addToken(TypeNumber, value)
	return nil
}

func (s *Scanner) isAlphaNumeric(c byte) bool {
	return s.isDigit(c) || s.isAlpha(c)
}

func (s *Scanner) identifier() error {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}

	s.addToken(TypeIdentifier, "")
	return nil
}

func (s *Scanner) addToken(tokentype TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokenList = append(s.tokenList, NewToken(tokentype, text, literal, s.line))
}

func (s *Scanner) scanString() error {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		return fmt.Errorf("unclosed string in line %d", s.line)
	}

	s.advance()

	s.addToken(TypeString, s.source[s.start+1:s.current-1])
	return nil
}

func (s *Scanner) ScanTokens() ([]*Token, error) {
	for !s.isAtEnd() {
		s.start = s.current
		err := s.scanToken()
		if err != nil {
			return nil, err
		}
	}

	s.tokenList = append(s.tokenList, NewToken(TypeEOF, "", nil, s.line))
	return s.tokenList, nil
}

func (s *Scanner) scanToken() error {
	c := s.advance()
	switch c {
	// Single-character 
	case '(':
		s.addToken(TypeLeftParen, nil)
	case ')':
		s.addToken(TypeRightParen, nil)
	case '{':
		s.addToken(TypeLeftBrace, nil)
	case '}':
		s.addToken(TypeRightBrace, nil)
	case '+':
		s.addToken(TypePlus, nil)
	case '-':
		s.addToken(TypeMinus, nil)
	case '*':
		s.addToken(TypeStar, nil)
	case '/':
		if s.match('/') {
			for !s.isAtEnd() && s.peek() != '\n' {
				s.advance()
			}
		} else if s.match('*') {
			depth := 1
			for !s.isAtEnd() && depth != 0 {
				c := s.advance()
				if c == '/' && s.peek() == '*' {
					depth++
					s.advance()
				}
				if c == '*' && s.peek() == '/' {
					depth--
					s.advance()
				}
			}
			if depth != 0 {
				return fmt.Errorf("unclosed comment in line: %d", s.line)
			}
		} else {
			s.addToken(TypeSlash, nil)
		}
	case ',':
		s.addToken(TypeComma, nil)
	case '.':
		s.addToken(TypeDot, nil)
	case ';':
		s.addToken(TypeSemicolon, nil)

	case '"': 
		return s.scanString()

	case ' ', '\t', '\r':

	case '\n':
		s.line++

	default:
		if s.isDigit(c) {
			return s.number()
		} else if s.isAlpha(c) {
			return s.identifier()
		} else {
			return fmt.Errorf("unrecognized token %c in line %d\n", c, s.line)
		}
	}
	return nil
}
