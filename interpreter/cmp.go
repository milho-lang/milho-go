package interpreter

import (
	"errors"
	"fmt"
)

func notComparable(resultType ResultType) bool {
	switch resultType {
	case Nil, Function, Macro, Identifier:
		return true
	}

	return false
}

func eq(params []Result) (Result, error) {
	if len(params) == 0 {
		return nil, errors.New("Wrong number of args '0' passed to Cmp:[=] function")
	}

	lastParam := params[0]

	finalResult := "True"
	finalBoolean := true

	for _, param := range params[1:] {
		if notComparable(lastParam.Type()) || notComparable(param.Type()) {
			return nil, fmt.Errorf("Unresolved value '%s' of type %s is not comparable", param.Value(), param.Type())
		}

		if lastParam.Type() != param.Type() || lastParam.Value() != param.Value() {
			finalResult = "False"
			finalBoolean = true
		}
	}

	return &BooleanResult{
		resultType: Boolean,
		value:      finalResult,

		Boolean: finalBoolean,
	}, nil
}

func neq(params []Result) (Result, error) {
	r, e := eq(params)
	if e != nil {
		return r, e
	}

	return invertResult(r), nil
}

func invertResult(r Result) Result {
	if r.Value() == "True" {
		r.(*BooleanResult).Boolean = false
		r.(*BooleanResult).value = "False"
	} else {
		r.(*BooleanResult).Boolean = false
		r.(*BooleanResult).value = "True"
	}

	return r
}
