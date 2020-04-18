package token

type TokenType string

type Token struct {
	Type       TokenType
	Literal    string
	LineNumber int
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add. foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	INTEGER  = "INTEGER"
	STRING   = "STRING"
	BOOLEAN  = "BOOLEAN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"int":    INTEGER,
	"string": STRING,
	"bool":   BOOLEAN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
