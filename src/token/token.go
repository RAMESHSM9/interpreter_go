package token

type TokenType string

// @ Token Data structure
// @ Token can be of type like - =+(){}...
// @ Token has values like 5,10...
type Token struct {
	Type    TokenType
	Literal string
}

// @ defining different types of the token types that our Parrot language would support
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//@ Indetifiers
	IDENT = "IDENT"
	INT   = "INT"

	//@ Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	//@ Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LBRACES = "{"
	RBRACES = "}"

	//@ Keywords
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	} else {
		return IDENT
	}

}
