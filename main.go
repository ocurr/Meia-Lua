package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("USAGE: typed-lua <file>")
		return
	}

	inputStream, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	lexer := parser.NewLuaLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(tokenStream)
	p.AddErrorListener(NewLuaErrorListener())
	p.BuildParseTrees = true
	tree := p.Chunk()
	builder := NewLuaASTBuilder()
	ast := tree.Accept(builder)
	fmt.Printf("%s", PrintLua(ast.(Node)))
}
