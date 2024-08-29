package token

type TokenType string

type Token struct {
	Type    TokenType // Declares new TokenType type
	Literal string    // Declares new Literal type with type string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// User-defined Identifiers + Literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1, 2, 3, ...

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS = "-"
	SLASH = "/"
	ASTERISK = "*" 
	LT = "<"
	GT = ">"
	BANG = "!"
	EQ = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"
)

// Define language keywords to distinguish from variable names
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
	"true": TRUE,
	"false": FALSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

/*
	Go note:
	In lookupIdent, tok = value, ok = does value exist in the map, ident = key
*/
