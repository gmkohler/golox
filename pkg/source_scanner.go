package golox

import (
	"fmt"
	"os"
	"strconv"
)

type sourceScanner struct {
	source  string // consider some sort of reader
	tokens  []Token
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

func (s *sourceScanner) ScanTokens() ([]Token, error) {
	var err error
	for !s.atEnd() {
		s.start = s.current
		if err = s.scanToken(); err != nil {
			return nil, err
		}
	}

	s.tokens = append(s.tokens, Token{
		tType: tokenTypeEof,
		line:  s.line,
	})

	return s.tokens, nil
}

func (s *sourceScanner) scanToken() error {
	switch r := s.advance(); r {
	case runeLeftParen:
		s.addToken(tokenTypeLeftParen)
	case runeRightParen:
		s.addToken(tokenTypeRightParen)
	case runeLeftBrace:
		s.addToken(tokenTypeLeftBrace)
	case runeRightBrace:
		s.addToken(tokenTypeRightBrace)
	case runeComma:
		s.addToken(tokenTypeComma)
	case runeDot:
		s.addToken(tokenTypeDot)
	case runeMinus:
		s.addToken(tokenTypeMinus)
	case runePlus:
		s.addToken(tokenTypePlus)
	case runeSemicolon:
		s.addToken(tokenTypeSemicolon)
	case runeStar:
		s.addToken(tokenTypeStar)
	case runeBang:
		token := tokenTypeBang
		if s.match(runeEqual) {
			token = tokenTypeBangEqual
		}
		s.addToken(token)
	case runeEqual:
		token := tokenTypeEqual
		if s.match(runeEqual) {
			token = tokenTypeEqualEqual
		}
		s.addToken(token)
	case runeLess:
		token := tokenTypeLess
		if s.match(runeEqual) {
			token = tokenTypeLessEqual
		}
		s.addToken(token)
	case runeGreater:
		token := tokenTypeGreater
		if s.match(runeEqual) {
			token = tokenTypeGreaterEqual
		}
		s.addToken(token)
	case runeSlash:
		if s.match(runeSlash) {
			// comments
			for s.peek() != runeNewline && !s.atEnd() {
				s.advance()
			}
		} else {
			s.addToken(tokenTypeSlash)
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

func (s *sourceScanner) addToken(tt tokenType) {
	s.addTokenWithLiteral(tt, nil)
}

func (s *sourceScanner) addTokenWithLiteral(tt tokenType, literal any) {
	s.tokens = append(
		s.tokens,
		Token{
			tType:   tt,
			lexeme:  s.source[s.start:s.current],
			literal: literal,
			line:    s.line,
		},
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
	s.addTokenWithLiteral(tokenTypeString, value)
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
	s.addTokenWithLiteral(tokenTypeNumber, f)
	return nil
}

func (s *sourceScanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}

	tt := tokenTypeFromKeyword(s.source[s.start:s.current])
	if tt == tokenTypeUnrecognized {
		tt = tokenTypeIdentifier
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
