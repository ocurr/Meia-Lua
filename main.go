package main

import (
	//"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
)

type TreeShapeListener struct {
	*parser.BaseLuaListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream("test.lua")
	lexer := parser.NewLuaLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.BuildParseTrees = true
	tree := p.Chunk()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
