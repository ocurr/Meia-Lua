package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
)

func main() {

	/*
		fmt.Printf("> ")
		var input strings.Builder
		var c rune
		var err error
		var n int
		for n, err = fmt.Scanf("%c", &c); n != 0 && err == nil && c != 10; {
			input.WriteRune(c)
			n, err = fmt.Scanf("%c", &c)
		}
		if err != nil {
			fmt.Println(err)
		}
	*/

	inputStream := antlr.NewInputStream("int x = 5 + 2")
	lexer := parser.NewLuaLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(tokenStream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Chunk()
	builder := NewLuaASTBuilder()
	ast := tree.Accept(builder)
	fmt.Printf("%#v\n", ast)
}
