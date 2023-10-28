package ast

import "fmt"

type Token struct {
	tType   TokenType
	lexeme  string
	literal any
	line    int
}

func NewToken(tType TokenType, lexeme string, line int) *Token {
	return NewTokenWithLiteral(tType, lexeme, nil, line)

}
func NewTokenWithLiteral(
	tType TokenType,
	lexeme string,
	literal any,
	line int,
) *Token {
	return &Token{
		tType:   tType,
		lexeme:  lexeme,
		literal: literal,
		line:    line,
	}
}

func (t Token) String() string {
	return fmt.Sprintf(
		"{ %v %v %v %v }",
		t.tType,
		t.lexeme,
		t.literal,
		t.line,
	)
}

type TokenType int

func (t TokenType) String() string {
	switch t {
	case TokenTypeLeftParen:
		return "LEFT_PAREN"
	case TokenTypeRightParen:
		return "RIGHT_PAREN"
	case TokenTypeLeftBrace:
		return "LEFT_BRACE"
	case TokenTypeRightBrace:
		return "RIGHT_BRACE"
	case TokenTypeComma:
		return "COMMA"
	case TokenTypeDot:
		return "DOT"
	case TokenTypeMinus:
		return "MINUS"
	case TokenTypePlus:
		return "PLUS"
	case TokenTypeSemicolon:
		return "SEMICOLON"
	case TokenTypeSlash:
		return "SLASH"
	case TokenTypeStar:
		return "STAR"
	case TokenTypeBang:
		return "BANG"
	case TokenTypeBangEqual:
		return "BANG_EQUAL"
	case TokenTypeEqual:
		return "EQUAL"
	case TokenTypeEqualEqual:
		return "EQUAL_EQUAL"
	case TokenTypeGreater:
		return "GREATER"
	case TokenTypeGreaterEqual:
		return "GREATER_EQUAL"
	case TokenTypeLess:
		return "LESS"
	case TokenTypeLessEqual:
		return "LESS_EQUAL"
	case TokenTypeIdentifier:
		return "IDENTIFIER"
	case TokenTypeString:
		return "STRING"
	case TokenTypeNumber:
		return "NUMBER"
	case TokenTypeAnd:
		return "AND"
	case TokenTypeClass:
		return "CLASS"
	case TokenTypeElse:
		return "ELSE"
	case TokenTypeFalse:
		return "FALSE"
	case TokenTypeFun:
		return "FUN"
	case TokenTypeFor:
		return "FOR"
	case TokenTypeIf:
		return "IF"
	case TokenTypeNil:
		return "NIL"
	case TokenTypeOr:
		return "OR"
	case TokenTypePrint:
		return "PRINT"
	case TokenTypeReturn:
		return "RETURN"
	case TokenTypeSuper:
		return "SUPER"
	case TokenTypeThis:
		return "THIS"
	case TokenTypeTrue:
		return "TRUE"
	case TokenTypeVar:
		return "VAR"
	case TokenTypeWhile:
		return "WHILE"
	case TokenTypeEof:
		return "EOF"
	default:
		return "unrecognized"
	}
}

const (
	// literals
	TokenTypeUnrecognized TokenType = iota
	// keywords
	TokenTypeLeftParen
	TokenTypeRightParen
	TokenTypeLeftBrace
	TokenTypeRightBrace
	TokenTypeComma
	TokenTypeDot
	TokenTypeMinus
	TokenTypePlus
	TokenTypeSemicolon
	TokenTypeSlash
	TokenTypeStar
	// one or two character tokens
	TokenTypeBang
	TokenTypeBangEqual
	TokenTypeEqual
	TokenTypeEqualEqual
	TokenTypeGreater
	TokenTypeGreaterEqual
	TokenTypeLess
	TokenTypeLessEqual
	// literals
	TokenTypeIdentifier
	TokenTypeString
	TokenTypeNumber
	// keywords
	TokenTypeAnd
	TokenTypeClass
	TokenTypeElse
	TokenTypeFalse
	TokenTypeFun
	TokenTypeFor
	TokenTypeIf
	TokenTypeNil
	TokenTypeOr
	TokenTypePrint
	TokenTypeReturn
	TokenTypeSuper
	TokenTypeThis
	TokenTypeTrue
	TokenTypeVar
	TokenTypeWhile

	TokenTypeEof
)

type keyword string

const (
	kwAnd    keyword = "and"
	kwClass  keyword = "class"
	kwElse   keyword = "else"
	kwFalse  keyword = "false"
	kwFor    keyword = "for"
	kwFun    keyword = "fun"
	kwIf     keyword = "if"
	kwNil    keyword = "nil"
	kwOr     keyword = "or"
	kwPrint  keyword = "print"
	kwReturn keyword = "return"
	kwSuper  keyword = "super"
	kwThis   keyword = "this"
	kwTrue   keyword = "true"
	kwVar    keyword = "var"
	kwWhile  keyword = "while"
)

var keywords = map[keyword]TokenType{
	kwAnd:    TokenTypeAnd,
	kwClass:  TokenTypeClass,
	kwElse:   TokenTypeElse,
	kwFalse:  TokenTypeFalse,
	kwFor:    TokenTypeFor,
	kwFun:    TokenTypeFun,
	kwIf:     TokenTypeIf,
	kwNil:    TokenTypeNil,
	kwOr:     TokenTypeOr,
	kwPrint:  TokenTypePrint,
	kwReturn: TokenTypeReturn,
	kwSuper:  TokenTypeSuper,
	kwThis:   TokenTypeThis,
	kwTrue:   TokenTypeTrue,
	kwVar:    TokenTypeVar,
	kwWhile:  TokenTypeWhile,
}

func TokenTypeFromKeyword(kw string) TokenType {
	tt, ok := keywords[keyword(kw)]
	if !ok {
		tt = TokenTypeUnrecognized // is a custom error better?
	}
	return tt
}
