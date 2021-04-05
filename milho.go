package milho

import (
	"fmt"

	"github.com/danfragoso/milho/interpreter"
	"github.com/danfragoso/milho/parser"
	"github.com/danfragoso/milho/tokenizer"
)

func Run(src string) string {
	var ret string
	tokens, err := tokenizer.Tokenize(src)
	if err != nil {
		return fmt.Sprintf("Tokenization error: %s\n", err)
	}

	ast, err := parser.Parse(tokens)
	if err != nil {
		return fmt.Sprintf("Parsing error: %s\n", err)
	}

	results, err := interpreter.Run(ast)
	for _, result := range results {
		if err != nil {
			ret += fmt.Sprintf("Evaluation error: %s\n", err)
		} else {
			ret += "\n" + result.Value()
		}
	}

	return ret
}
