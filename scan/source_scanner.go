package scan

import (
	"fmt"
	"github.com/gmkohler/golox/ast"
	"os"
	"strconv"
)

type sourceScanner struct {
	source  string // consider some sort of reader
	tokens  []*ast.Token
	start   int
	current int
	line    int
}

func NewSourceScanner(source string) Scanner {
	return &sourceScanner{
		source: source,
		line:   1,
	}
}

func (s *sourceScanner) ScanTokens() ([]*ast.Token, error) {
	var err error
	for !s.atEnd() {
		s.start = s.current
		if err = s.scanToken(); err != nil {
			return nil, err
		}
	}

	s.tokens = append(s.tokens, ast.NewToken(
		ast.TokenTypeEof,
		"",
		s.line,
	))

	return s.tokens, nil
}

func (s *sourceScanner) scanToken() error {
	switch r := s.advance(); r {
	case runeLeftParen:
		s.addToken(ast.TokenTypeLeftParen)
	case runeRightParen:
		s.addToken(ast.TokenTypeRightParen)
	case runeLeftBrace:
		s.addToken(ast.TokenTypeLeftBrace)
	case runeRightBrace:
		s.addToken(ast.TokenTypeRightBrace)
	case runeComma:
		s.addToken(ast.TokenTypeComma)
	case runeDot:
		s.addToken(ast.TokenTypeDot)
	case runeMinus:
		s.addToken(ast.TokenTypeMinus)
	case runePlus:
		s.addToken(ast.TokenTypePlus)
	case runeSemicolon:
		s.addToken(ast.TokenTypeSemicolon)
	case runeStar:
		s.addToken(ast.TokenTypeStar)
	case runeBang:
		token := ast.TokenTypeBang
		if s.match(runeEqual) {
			token = ast.TokenTypeBangEqual
		}
		s.addToken(token)
	case runeEqual:
		token := ast.TokenTypeEqual
		if s.match(runeEqual) {
			token = ast.TokenTypeEqualEqual
		}
		s.addToken(token)
	case runeLess:
		token := ast.TokenTypeLess
		if s.match(runeEqual) {
			token = ast.TokenTypeLessEqual
		}
		s.addToken(token)
	case runeGreater:
		token := ast.TokenTypeGreater
		if s.match(runeEqual) {
			token = ast.TokenTypeGreaterEqual
		}
		s.addToken(token)
	case runeSlash:
		if s.match(runeSlash) {
			// comments
			for s.peek() != runeNewline && !s.atEnd() {
				s.advance()
			}
		} else {
			s.addToken(ast.TokenTypeSlash)
		}
	case runeNewline:
		s.line++
	case runeSpace, runeCarriageReturn, runeTab:
	case runeDoubleQuote:
		if err := s.string(); err != nil {
			return err
		}
	default:
		if isDigit(r) {
			if err := s.number(); err != nil {
				return err
			}
		} else if isAlpha(r) {
			s.identifier()
		} else {
			_, _ = fmt.Fprintf(os.Stderr, "unexpected character %c\n", r)
		}
	}

	return nil
}

func (s *sourceScanner) addToken(tt ast.TokenType) {
	s.addTokenWithLiteral(tt, nil)
}

func (s *sourceScanner) addTokenWithLiteral(tt ast.TokenType, literal any) {
	s.tokens = append(
		s.tokens,
		ast.NewTokenWithLiteral(
			tt,
			s.source[s.start:s.current],
			literal,
			s.line,
		),
	)
}

func (s *sourceScanner) advance() rune {
	r := rune(s.source[s.current])
	s.current++
	return r
}

func (s *sourceScanner) match(expected rune) bool {
	if s.atEnd() {
		return false
	}
	if rune(s.source[s.current]) != expected {
		return false
	}
	s.current++
	return true
}

func (s *sourceScanner) string() error {
	for s.peek() != runeDoubleQuote && !s.atEnd() {
		if s.peek() == runeNewline {
			s.line++
		}
		s.advance()
	}
	if s.atEnd() {
		return fmt.Errorf("unterminated string (line %d)", s.line)
	}
	s.advance()
	value := s.source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(ast.TokenTypeString, value)
	return nil
}

func (s *sourceScanner) number() error {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == runeDot && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		return err
	}
	s.addTokenWithLiteral(ast.TokenTypeNumber, f)
	return nil
}

func (s *sourceScanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	tt := ast.TokenTypeFromKeyword(s.source[s.start:s.current])
	if tt == ast.TokenTypeUnrecognized {
		tt = ast.TokenTypeIdentifier
	}
	s.addToken(tt)
}

func (s *sourceScanner) peek() rune {
	if s.atEnd() {
		return rune(0)
	}
	return rune(s.source[s.current])
}

func (s *sourceScanner) peekNext() rune {
	if s.current+1 > len(s.source) {
		return runeNull
	}
	return rune(s.source[s.current+1])
}

func (s *sourceScanner) atEnd() bool {
	return s.current >= len(s.source)
}
