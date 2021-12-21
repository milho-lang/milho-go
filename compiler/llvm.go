package compiler

import (
	"fmt"

	"github.com/danfragoso/milho/interpreter"
)

var llvm_header = `
@memes = global i32 1

define i32 @main() {
	%1 = load i32, i32* @memes
	%2 = add i32 %1, 1
	store i32 %2, i32* @memes
	ret i32* @memes
}
`

type llvm_instruction interface {
	String() string
}

func TranspileLLVM(expr interpreter.Expression) string {
	fmt.Println(";Generated by the milho compiler")
	return llvm_header //+ expr.Value()
}

func llvm_transpileExpr(expr interpreter.Expression) string {
	switch expr.Type() {
	case interpreter.ListExpr:
		return llvm_transpileListExpr(expr).String()
	}

	return "call void @llvm.donothing()"
}

func llvm_transpileListExpr(expr interpreter.Expression) llvm_instruction {
	expressions := expr.(*interpreter.ListExpression).Expressions
	if len(expressions) == 0 {
		return &JSValue_Undefined{}
	}

	firstExpr := expressions[0]
	switch firstExpr.Type() {
	case interpreter.SymbolExpr:
		sym := firstExpr.(*interpreter.SymbolExpression)
		return matchListSymbolExpr(sym, expressions[1:])
	}

	return &JSValue_String{}
}
