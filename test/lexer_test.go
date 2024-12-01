package test

import (
	"testing"

	"github.com/RAMESHSM9/interpreter_go/src/token"

	"github.com/RAMESHSM9/interpreter_go/src/lexer"
)

func TestNextToken(t *testing.T) {

	input := `let five = 5;
	let ten = 10;
	let add = fn(x, y){
		x+y;
	}
	let result = add(five, ten);
	`

	tests := []struct {
		expectedTokenType token.TokenType
		expectedLiteral   string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACES, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACES, "}"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := lexer.New(input)

	for i, tokenType := range tests {

		tok := l.NextToken()

		if tok.Type != tokenType.expectedTokenType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tokenType.expectedTokenType, tok.Type)
		}

		if tok.Literal != tokenType.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tokenType.expectedLiteral, tok.Literal)
		}

	}

}
