package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/scwood/writing-an-interpreter-in-go/lexer"
	"github.com/scwood/writing-an-interpreter-in-go/token"
)

const prompt = "> "

func Start(in io.Reader, out io.StringWriter) {
	scanner := bufio.NewScanner(in)
	for {
		out.WriteString(fmt.Sprint(prompt))
		scanner.Scan()
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			out.WriteString(fmt.Sprintf("%+v\n", tok))
		}
	}
}
