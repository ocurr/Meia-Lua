package ast

import (
	"testing"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/ocurr/Meia-Lua/parser"
	"github.com/ocurr/Meia-Lua/types"
)

var astCompOpts = []cmp.Option{
	cmpopts.IgnoreFields(Chunk{}, "Ctx"),
	cmpopts.IgnoreFields(Block{}, "Ctx"),
	cmpopts.IgnoreFields(Cond{}, "Ctx"),
	cmpopts.IgnoreFields(While{}, "Ctx"),
	cmpopts.IgnoreFields(For{}, "Ctx"),
	cmpopts.IgnoreFields(Def{}, "Ctx"),
	cmpopts.IgnoreFields(DefLst{}, "Ctx"),
	cmpopts.IgnoreFields(Id{}, "Ctx"),
	cmpopts.IgnoreFields(IdLst{}, "Ctx"),
	cmpopts.IgnoreFields(ExpLst{}, "Ctx"),
	cmpopts.IgnoreFields(BinaryOp{}, "Ctx"),
	cmpopts.IgnoreFields(Int{}, "Ctx"),
	cmpopts.IgnoreFields(Float{}, "Ctx"),
	cmpopts.IgnoreFields(String{}, "Ctx"),
	cmpopts.IgnoreFields(Bool{}, "Ctx"),
	cmpopts.IgnoreFields(Nil{}, "Ctx"),
}

func TestASTBuilder(t *testing.T) {
	t.Run("SingleIntAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: Int{N: 5}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("SingleFloatAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("float x = 5.5")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Float{}}, Exp: Float{N: 5.5}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("SingleStringAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`string animal = "horse"`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "animal", TypeId: types.String{}}, Exp: String{S: "horse"}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("MultipleIntAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x, y, z = 5, 6, 7")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{
				{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: Int{N: 5}, Scope: GLOBAL},
				{Id: Id{Id: "y", TypeId: types.Int{}}, Exp: Int{N: 6}, Scope: GLOBAL},
				{Id: Id{Id: "z", TypeId: types.Int{}}, Exp: Int{N: 7}, Scope: GLOBAL},
			}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("MultipleFloatAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("float x, y, z = 5.5, 10.789, 100.23")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{
				{Id: Id{Id: "x", TypeId: types.Float{}}, Exp: Float{N: 5.5}, Scope: GLOBAL},
				{Id: Id{Id: "y", TypeId: types.Float{}}, Exp: Float{N: 10.789}, Scope: GLOBAL},
				{Id: Id{Id: "z", TypeId: types.Float{}}, Exp: Float{N: 100.23}, Scope: GLOBAL},
			}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("MultipleStringAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`string animal, dog, word = "horse", "fido", "bar"`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{
				{Id: Id{Id: "animal", TypeId: types.String{}}, Exp: String{S: "horse"}, Scope: GLOBAL},
				{Id: Id{Id: "dog", TypeId: types.String{}}, Exp: String{S: "fido"}, Scope: GLOBAL},
				{Id: Id{Id: "word", TypeId: types.String{}}, Exp: String{S: "bar"}, Scope: GLOBAL},
			}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("AdditionAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 + 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "+",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("SubtractionAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 - 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "-",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("MultiplyAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 * 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "*",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("ModAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 % 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "%",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("DivideAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 / 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "/",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("BoolAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = true")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: Bool{True: true}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("LessThanAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 < 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "<",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("LessThanEqualAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 <= 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "<=",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("GreaterThanAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 > 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  ">",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("GreaterThanEqualAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 >= 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  ">=",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("EqualAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 == 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "==",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("NotEqualAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("bool x = 5 ~= 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "~=",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("IntDivideAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 // 2")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
				Lhs: Int{N: 5},
				Rhs: Int{N: 2},
				Op:  "//",
			}, Scope: GLOBAL}}}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("Reassign", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`
		int x = 5 // 2
		x = 1 + 1
		`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{
			StatLst: []Stat{
				DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: BinaryOp{
					Lhs: Int{N: 5},
					Rhs: Int{N: 2},
					Op:  "//",
				}, Scope: GLOBAL}}},
				DefLst{List: []Def{{Id: Id{Id: "x", TypeId: nil}, Exp: BinaryOp{
					Lhs: Int{N: 1},
					Rhs: Int{N: 1},
					Op:  "+",
				}, Scope: GLOBAL}}},
			}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("IfOnly", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`
		if true then
			x = 5
		end
		`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)

		want := Chunk{Block: Block{
			StatLst: []Stat{
				Cond{
					Cnd: Bool{True: true},
					Block: Block{StatLst: []Stat{
						DefLst{List: []Def{{Id: Id{Id: "x"}, Exp: Int{N: 5}, Scope: GLOBAL}}}},
					},
				}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("WhileOnly", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`
		while true do
			x = 5
		end
		`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)

		want := Chunk{Block: Block{
			StatLst: []Stat{
				While{
					Cnd: Bool{True: true},
					Block: Block{StatLst: []Stat{
						DefLst{List: []Def{{Id: Id{Id: "x"}, Exp: Int{N: 5}, Scope: GLOBAL}}}},
					},
				}}}}
		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("For", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`
		for int x = 5, x < 5, 1 do
			y = 5
		end
		`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)

		want := Chunk{Block: Block{StatLst: []Stat{
			For{
				Assign: Def{
					Id: Id{
						Id:     "x",
						TypeId: types.Int{},
					},
					Exp:   Int{N: 5},
					Scope: GLOBAL,
				},
				Cnd: BinaryOp{
					Lhs: Id{Id: "x"},
					Rhs: Int{N: 5},
					Op:  "<",
				},
				Step:  Int{N: 1},
				Block: Block{StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "y"}, Exp: Int{N: 5}, Scope: GLOBAL}}}}},
			},
		}}}

		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("ForNoStep", func(t *testing.T) {
		inputStream := antlr.NewInputStream(`
		for int x = 5, x < 5 do
			y = 5
		end
		`)
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)

		want := Chunk{Block: Block{StatLst: []Stat{
			For{
				Assign: Def{
					Id: Id{
						Id:     "x",
						TypeId: types.Int{},
					},
					Exp:   Int{N: 5},
					Scope: GLOBAL,
				},
				Cnd: BinaryOp{
					Lhs: Id{Id: "x"},
					Rhs: Int{N: 5},
					Op:  "<",
				},
				Step:  Int{N: 1},
				Block: Block{StatLst: []Stat{DefLst{List: []Def{{Id: Id{Id: "y"}, Exp: Int{N: 5}, Scope: GLOBAL}}}}},
			},
		}}}

		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("Nil", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = nil")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{StatLst: []Stat{
			DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: Nil{}, Scope: GLOBAL}}},
		}}}

		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
	t.Run("LocalVar", func(t *testing.T) {
		inputStream := antlr.NewInputStream("local int x")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewBuilder()

		got := tree.Accept(builder)
		want := Chunk{Block: Block{StatLst: []Stat{
			DefLst{List: []Def{{Id: Id{Id: "x", TypeId: types.Int{}}, Exp: Nil{}, Scope: LOCAL}}},
		}}}

		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
}
