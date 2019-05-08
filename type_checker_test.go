package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
	"testing"
)

func TestTypeChecker(t *testing.T) {
	t.Run("BasicDefinition", func(t *testing.T) {
		input := DefC{
			Id:  IdC{Id: "x", TypeId: IntT{}},
			Exp: IntC{N: 5},
		}

		got, errs := LuaTypeCheck(input)

		want := IntT{}

		if len(errs) != 0 {
			t.Errorf("expected no errors got: %v", errs)
		}
		if want != got {
			t.Errorf("wanted: %#v ; got: %#v", want, got)
		}
	})

	t.Run("If", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}}},
			}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})

	t.Run("IfElseIf", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}},
					Elseifs: []CondC{
						CondC{
							Cnd: BoolC{True: false},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 5}}}}}}},
					},
				}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("IfElseIfElse", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}},
					Elseifs: []CondC{
						CondC{
							Cnd: BoolC{True: false},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 5}}}}}}},
					},
					Else: BlockC{StatLst: []Stat{
						CondC{
							Cnd: BoolC{True: true},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 6}}}}}},
						}},
					}}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("While", func(t *testing.T) {
		input := WhileC{
			Cnd: BoolC{True: true},
			Block: BlockC{
				StatLst: []Stat{
					DefLst{List: []DefC{
						DefC{Id: IdC{Id: "x", TypeId: IntT{}},
							Exp: IntC{N: 6}}}}}},
		}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("Closure", func(t *testing.T) {
		input := "int x = 5\n" +
			"string y = \"horse\"\n" +
			"while x < 10 do\nint y = 10\nx = x + y\nend\n" +
			"if y == \"horse\" then\n" +
			"int z = 7\n" +
			"end"

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("For", func(t *testing.T) {
		input := "for int i=0, 10 do\n" +
			"int x = 5\n" +
			"end"

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("ForNoBody", func(t *testing.T) {

		_, err := LuaTypeCheck(ForC{
			Assign: DefC{
				Id:  IdC{Id: "x", TypeId: IntT{}},
				Exp: IntC{N: 5},
			},
			Cnd:   IntC{N: 10},
			Step:  IntC{N: 1},
			Block: BlockC{StatLst: []Stat{}},
		})
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("ForWithStep", func(t *testing.T) {
		input := "for int i=0, 10, 1 do\n" +
			"int x = 5\n" +
			"end"

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("LocalVar", func(t *testing.T) {
		input := "local int x = 5\n" +
			"x = x + 5"
		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
}

func buildInputTree(input string) ChunkC {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewLuaLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(tokenStream)
	tree := p.Chunk()
	p.BuildParseTrees = true
	builder := NewLuaASTBuilder()

	return tree.Accept(builder).(ChunkC)
}
