package main

import (
	"testing"

	"github.com/ocurr/Meia-Lua/ast"
	"github.com/ocurr/Meia-Lua/parser"
	"github.com/ocurr/Meia-Lua/types"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func TestTypeChecker(t *testing.T) {
	t.Run("BasicDefinition", func(t *testing.T) {
		input := ast.Def{
			Id:  ast.Id{Id: "x", TypeId: types.Int{}},
			Exp: ast.Int{N: 5},
		}

		got, errs := LuaTypeCheck(input)

		want := types.Int{}

		if len(errs) != 0 {
			t.Errorf("expected no errors got: %v", errs)
		}
		if want != got {
			t.Errorf("wanted: %#v ; got: %#v", want, got)
		}
	})

	t.Run("If", func(t *testing.T) {
		input := ast.Cond{
			Cnd: ast.Bool{True: true},
			Block: ast.Block{StatLst: []ast.Stat{
				ast.DefLst{List: []ast.Def{
					{Id: ast.Id{Id: "x", TypeId: types.Int{}},
						Exp: ast.Int{N: 5}}}},
				ast.Cond{
					Cnd: ast.Bool{True: true},
					Block: ast.Block{
						StatLst: []ast.Stat{
							ast.DefLst{List: []ast.Def{
								{Id: ast.Id{Id: "x", TypeId: types.Int{}},
									Exp: ast.Int{N: 6}}}}}}},
			}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})

	t.Run("IfElseIf", func(t *testing.T) {
		input := ast.Cond{
			Cnd: ast.Bool{True: true},
			Block: ast.Block{StatLst: []ast.Stat{
				ast.DefLst{List: []ast.Def{
					{Id: ast.Id{Id: "x", TypeId: types.Int{}},
						Exp: ast.Int{N: 5}}}},
				ast.Cond{
					Cnd: ast.Bool{True: true},
					Block: ast.Block{
						StatLst: []ast.Stat{
							ast.DefLst{List: []ast.Def{
								{Id: ast.Id{Id: "x", TypeId: types.Int{}},
									Exp: ast.Int{N: 6}}}}}},
					Elseifs: []ast.Cond{
						{
							Cnd: ast.Bool{True: false},
							Block: ast.Block{
								StatLst: []ast.Stat{
									ast.DefLst{List: []ast.Def{
										{Id: ast.Id{Id: "x", TypeId: types.Int{}},
											Exp: ast.Int{N: 5}}}}}}},
					},
				}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("IfElseIfElse", func(t *testing.T) {
		input := ast.Cond{
			Cnd: ast.Bool{True: true},
			Block: ast.Block{StatLst: []ast.Stat{
				ast.DefLst{List: []ast.Def{
					{Id: ast.Id{Id: "x", TypeId: types.Int{}},
						Exp: ast.Int{N: 5}}}},
				ast.Cond{
					Cnd: ast.Bool{True: true},
					Block: ast.Block{
						StatLst: []ast.Stat{
							ast.DefLst{List: []ast.Def{
								{Id: ast.Id{Id: "x", TypeId: types.Int{}},
									Exp: ast.Int{N: 6}}}}}},
					Elseifs: []ast.Cond{
						{
							Cnd: ast.Bool{True: false},
							Block: ast.Block{
								StatLst: []ast.Stat{
									ast.DefLst{List: []ast.Def{
										{Id: ast.Id{Id: "x", TypeId: types.Int{}},
											Exp: ast.Int{N: 5}}}}}}},
					},
					Else: ast.Block{StatLst: []ast.Stat{
						ast.Cond{
							Cnd: ast.Bool{True: true},
							Block: ast.Block{
								StatLst: []ast.Stat{
									ast.DefLst{List: []ast.Def{
										{Id: ast.Id{Id: "x", TypeId: types.Int{}},
											Exp: ast.Int{N: 6}}}}}},
						}},
					}}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("While", func(t *testing.T) {
		input := ast.While{
			Cnd: ast.Bool{True: true},
			Block: ast.Block{
				StatLst: []ast.Stat{
					ast.DefLst{List: []ast.Def{
						{Id: ast.Id{Id: "x", TypeId: types.Int{}},
							Exp: ast.Int{N: 6}}}}}},
		}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("Closure", func(t *testing.T) {
		input := `
		int x = 5
		string y = "horse"
		while x < 10 do
			int y = 10
			x = x + y
		end
		if y == "horse" then
			int z = 7
		end
		`

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("For", func(t *testing.T) {
		input := `
		for int i=0, 10 do
			int x = 5
		end
		`

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("ForNoBody", func(t *testing.T) {

		_, err := LuaTypeCheck(ast.For{
			Assign: ast.Def{
				Id:  ast.Id{Id: "x", TypeId: types.Int{}},
				Exp: ast.Int{N: 5},
			},
			Cnd:   ast.Int{N: 10},
			Step:  ast.Int{N: 1},
			Block: ast.Block{StatLst: []ast.Stat{}},
		})
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("ForWithStep", func(t *testing.T) {
		input := `
		for int i=0, 10, 1 do
			int x = 5
		end
		`

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("LocalVar", func(t *testing.T) {
		input := `
		local int x = 5
		x = x + 5
		`

		_, err := LuaTypeCheck(buildInputTree(input))
		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
}

func TestEquality(t *testing.T) {
	input := `
	bool y = 5 == "horse"
	`

	_, err := LuaTypeCheck(buildInputTree(input))
	if len(err) == 0 {
		t.Errorf("expected errors, got no errors instead")
	}
}

func buildInputTree(input string) ast.Chunk {
	inputStream := antlr.NewInputStream(input)
	lexer := parser.NewLuaLexer(inputStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLuaParser(tokenStream)
	tree := p.Chunk()
	p.BuildParseTrees = true
	builder := ast.NewBuilder()

	return tree.Accept(builder).(ast.Chunk)
}
