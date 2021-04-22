package interpreter

import (
	"errors"
	"fmt"
)

func __def(params []Expression, session *Session) (Expression, error) {
	if len(params) != 2 {
		return nil, errors.New("Wrong number of args '2' passed to def")
	}

	firstParam := params[0]
	if firstParam.Type() != SymbolExpr {
		return nil, fmt.Errorf("First argument of def should be a symbol")
	}

	symbol, err := createSymbolExpression(firstParam.(*SymbolExpression).Identifier, params[1])
	if err != nil {
		return nil, err
	}

	err = session.AddObject(symbol.Identifier, symbol.Expression)
	if err != nil {
		return nil, err
	}

	return symbol, nil
}