package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/RAMESHSM9/interpreter_go/src/lexer"
	"github.com/RAMESHSM9/interpreter_go/src/token"
)

const PROMT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
