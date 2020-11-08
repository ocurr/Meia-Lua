package main

import (
	"fmt"
	"os"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/Meia-Lua/ast"
	"github.com/ocurr/Meia-Lua/parser"
	"github.com/ocurr/Meia-Lua/typecheck"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("USAGE: meia-lua <file>")
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
	builder := ast.NewBuilder()
	mLuaAST := tree.Accept(builder)
	_, errs := typecheck.Check(mLuaAST.(ast.Node))
	if len(errs) != 0 {
		for _, err := range errs {
			fmt.Println("ERROR: ", err)
		}
		return
	}
	fmt.Printf("%s", PrintLua(mLuaAST.(ast.Node)))
}
