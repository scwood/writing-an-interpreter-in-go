package token

type TokenType string

// Possible token types
const (
	Assign      = "Assign"
	Plus        = "Plus"
	Minus       = "Minus"
	Bang        = "Bang"
	Asterisk    = "Asterisk"
	Slash       = "Slash"
	LessThan    = "LessThan"
	GreaterThan = "GreaterThan"
	Equal       = "Equal"
	NotEqual    = "NotEqual"
	Comma       = "Comma"
	Semicolon   = "Semicolon"
	LeftParen   = "LeftParen"
	RightParen  = "RightParen"
	LeftBrace   = "LeftBrace"
	RightBrace  = "RightBrace"
	Function    = "Function"
	Let         = "Let"
	True        = "True"
	False       = "False"
	If          = "If"
	Else        = "Else"
	Return      = "Return"
	Identifier  = "Identifier"
	Int         = "Int"
	Illegal     = "Illegal"
	EOF         = "EOF"
)

var keywords = map[string]TokenType{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdentifier(ident string) TokenType {
	if tokenType, exists := keywords[ident]; exists {
		return tokenType
	}
	return Identifier
}
