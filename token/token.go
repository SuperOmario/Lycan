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

	// Artihmetic Operators
	ASSIGN   = "="
	ADD      = "+"
	SUBTRACT = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"

	//Relation Operators
	EQUAL   = "=="
	N_EQUAL = "!="
	NEGATE  = "!"
	LT      = "<"
	GT      = ">"

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
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"int":    INTEGER,
	"string": STRING,
	"bool":   BOOLEAN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

//Searches the keywords map for the identifier
//If it exists the token literal is returned as tok and token is returned as ok
//If it doesn't exist the identifier is returned as an IDENT token
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
