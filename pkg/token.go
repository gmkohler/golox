package golox

import "fmt"

type Token struct {
	tType   tokenType
	lexeme  string
	literal any
	line    int
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

type tokenType int

func (t tokenType) String() string {
	switch t {
	case tokenTypeLeftParen:
		return "LEFT_PAREN"
	case tokenTypeRightParen:
		return "RIGHT_PAREN"
	case tokenTypeLeftBrace:
		return "LEFT_BRACE"
	case tokenTypeRightBrace:
		return "RIGHT_BRACE"
	case tokenTypeComma:
		return "COMMA"
	case tokenTypeDot:
		return "DOT"
	case tokenTypeMinus:
		return "MINUS"
	case tokenTypePlus:
		return "PLUS"
	case tokenTypeSemicolon:
		return "SEMICOLON"
	case tokenTypeSlash:
		return "SLASH"
	case tokenTypeStar:
		return "STAR"
	case tokenTypeBang:
		return "BANG"
	case tokenTypeBangEqual:
		return "BANG_EQUAL"
	case tokenTypeEqual:
		return "EQUAL"
	case tokenTypeEqualEqual:
		return "EQUAL_EQUAL"
	case tokenTypeGreater:
		return "GREATER"
	case tokenTypeGreaterEqual:
		return "GREATER_EQUAL"
	case tokenTypeLess:
		return "LESS"
	case tokenTypeLessEqual:
		return "LESS_EQUAL"
	case tokenTypeIdentifier:
		return "IDENTIFIER"
	case tokenTypeString:
		return "STRING"
	case tokenTypeNumber:
		return "NUMBER"
	case tokenTypeAnd:
		return "AND"
	case tokenTypeClass:
		return "CLASS"
	case tokenTypeElse:
		return "ELSE"
	case tokenTypeFalse:
		return "FALSE"
	case tokenTypeFun:
		return "FUN"
	case tokenTypeFor:
		return "FOR"
	case tokenTypeIf:
		return "IF"
	case tokenTypeNil:
		return "NIL"
	case tokenTypeOr:
		return "OR"
	case tokenTypePrint:
		return "PRINT"
	case tokenTypeReturn:
		return "RETURN"
	case tokenTypeSuper:
		return "SUPER"
	case tokenTypeThis:
		return "THIS"
	case tokenTypeTrue:
		return "TRUE"
	case tokenTypeVar:
		return "VAR"
	case tokenTypeWhile:
		return "WHILE"
	case tokenTypeEof:
		return "EOF"
	default:
		return "unrecognized"
	}
}

const (
	// literals
	tokenTypeUnrecognized tokenType = iota
	// keywords
	tokenTypeLeftParen
	tokenTypeRightParen
	tokenTypeLeftBrace
	tokenTypeRightBrace
	tokenTypeComma
	tokenTypeDot
	tokenTypeMinus
	tokenTypePlus
	tokenTypeSemicolon
	tokenTypeSlash
	tokenTypeStar
	// one or two character tokens
	tokenTypeBang
	tokenTypeBangEqual
	tokenTypeEqual
	tokenTypeEqualEqual
	tokenTypeGreater
	tokenTypeGreaterEqual
	tokenTypeLess
	tokenTypeLessEqual
	// literals
	tokenTypeIdentifier
	tokenTypeString
	tokenTypeNumber
	// keywords
	tokenTypeAnd
	tokenTypeClass
	tokenTypeElse
	tokenTypeFalse
	tokenTypeFun
	tokenTypeFor
	tokenTypeIf
	tokenTypeNil
	tokenTypeOr
	tokenTypePrint
	tokenTypeReturn
	tokenTypeSuper
	tokenTypeThis
	tokenTypeTrue
	tokenTypeVar
	tokenTypeWhile

	tokenTypeEof
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

var keywords = map[keyword]tokenType{
	kwAnd:    tokenTypeAnd,
	kwClass:  tokenTypeClass,
	kwElse:   tokenTypeElse,
	kwFalse:  tokenTypeFalse,
	kwFor:    tokenTypeFor,
	kwFun:    tokenTypeFun,
	kwIf:     tokenTypeIf,
	kwNil:    tokenTypeNil,
	kwOr:     tokenTypeOr,
	kwPrint:  tokenTypePrint,
	kwReturn: tokenTypeReturn,
	kwSuper:  tokenTypeSuper,
	kwThis:   tokenTypeThis,
	kwTrue:   tokenTypeTrue,
	kwVar:    tokenTypeVar,
	kwWhile:  tokenTypeWhile,
}

func tokenTypeFromKeyword(kw string) tokenType {
	tt, ok := keywords[keyword(kw)]
	if !ok {
		tt = tokenTypeUnrecognized // is a custom error better?
	}
	return tt
}
