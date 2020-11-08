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
	cmpopts.IgnoreFields(ChunkC{}, "Ctx"),
	cmpopts.IgnoreFields(BlockC{}, "Ctx"),
	cmpopts.IgnoreFields(CondC{}, "Ctx"),
	cmpopts.IgnoreFields(WhileC{}, "Ctx"),
	cmpopts.IgnoreFields(ForC{}, "Ctx"),
	cmpopts.IgnoreFields(DefC{}, "Ctx"),
	cmpopts.IgnoreFields(DefLst{}, "Ctx"),
	cmpopts.IgnoreFields(IdC{}, "Ctx"),
	cmpopts.IgnoreFields(IdLst{}, "Ctx"),
	cmpopts.IgnoreFields(ExpLst{}, "Ctx"),
	cmpopts.IgnoreFields(BinaryOpC{}, "Ctx"),
	cmpopts.IgnoreFields(IntC{}, "Ctx"),
	cmpopts.IgnoreFields(FloatC{}, "Ctx"),
	cmpopts.IgnoreFields(StringC{}, "Ctx"),
	cmpopts.IgnoreFields(BoolC{}, "Ctx"),
	cmpopts.IgnoreFields(NilC{}, "Ctx"),
}

func TestASTBuilder(t *testing.T) {
	t.Run("SingleIntAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: IntC{N: 5}, Scope: GLOBAL}}}}}}
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Float{}}, Exp: FloatC{N: 5.5}, Scope: GLOBAL}}}}}}
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "animal", TypeId: types.String{}}, Exp: StringC{S: "horse"}, Scope: GLOBAL}}}}}}
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{
				{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: IntC{N: 5}, Scope: GLOBAL},
				{Id: IdC{Id: "y", TypeId: types.Int{}}, Exp: IntC{N: 6}, Scope: GLOBAL},
				{Id: IdC{Id: "z", TypeId: types.Int{}}, Exp: IntC{N: 7}, Scope: GLOBAL},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{
				{Id: IdC{Id: "x", TypeId: types.Float{}}, Exp: FloatC{N: 5.5}, Scope: GLOBAL},
				{Id: IdC{Id: "y", TypeId: types.Float{}}, Exp: FloatC{N: 10.789}, Scope: GLOBAL},
				{Id: IdC{Id: "z", TypeId: types.Float{}}, Exp: FloatC{N: 100.23}, Scope: GLOBAL},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{
				{Id: IdC{Id: "animal", TypeId: types.String{}}, Exp: StringC{S: "horse"}, Scope: GLOBAL},
				{Id: IdC{Id: "dog", TypeId: types.String{}}, Exp: StringC{S: "fido"}, Scope: GLOBAL},
				{Id: IdC{Id: "word", TypeId: types.String{}}, Exp: StringC{S: "bar"}, Scope: GLOBAL},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BoolC{True: true}, Scope: GLOBAL}}}}}}
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Bool{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{
				DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: BinaryOpC{
					Lhs: IntC{N: 5},
					Rhs: IntC{N: 2},
					Op:  "//",
				}, Scope: GLOBAL}}},
				DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: nil}, Exp: BinaryOpC{
					Lhs: IntC{N: 1},
					Rhs: IntC{N: 1},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)

		want := ChunkC{Block: BlockC{
			StatLst: []Stat{
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{StatLst: []Stat{
						DefLst{List: []DefC{{Id: IdC{Id: "x"}, Exp: IntC{N: 5}, Scope: GLOBAL}}}},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)

		want := ChunkC{Block: BlockC{
			StatLst: []Stat{
				WhileC{
					Cnd: BoolC{True: true},
					Block: BlockC{StatLst: []Stat{
						DefLst{List: []DefC{{Id: IdC{Id: "x"}, Exp: IntC{N: 5}, Scope: GLOBAL}}}},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)

		want := ChunkC{Block: BlockC{StatLst: []Stat{
			ForC{
				Assign: DefC{
					Id: IdC{
						Id:     "x",
						TypeId: types.Int{},
					},
					Exp:   IntC{N: 5},
					Scope: GLOBAL,
				},
				Cnd: BinaryOpC{
					Lhs: IdC{Id: "x"},
					Rhs: IntC{N: 5},
					Op:  "<",
				},
				Step:  IntC{N: 1},
				Block: BlockC{StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "y"}, Exp: IntC{N: 5}, Scope: GLOBAL}}}}},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)

		want := ChunkC{Block: BlockC{StatLst: []Stat{
			ForC{
				Assign: DefC{
					Id: IdC{
						Id:     "x",
						TypeId: types.Int{},
					},
					Exp:   IntC{N: 5},
					Scope: GLOBAL,
				},
				Cnd: BinaryOpC{
					Lhs: IdC{Id: "x"},
					Rhs: IntC{N: 5},
					Op:  "<",
				},
				Step:  IntC{N: 1},
				Block: BlockC{StatLst: []Stat{DefLst{List: []DefC{{Id: IdC{Id: "y"}, Exp: IntC{N: 5}, Scope: GLOBAL}}}}},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{StatLst: []Stat{
			DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: NilC{}, Scope: GLOBAL}}},
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
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{StatLst: []Stat{
			DefLst{List: []DefC{{Id: IdC{Id: "x", TypeId: types.Int{}}, Exp: NilC{}, Scope: LOCAL}}},
		}}}

		if diff := cmp.Diff(want, got, astCompOpts...); diff != "" {
			t.Errorf("AST mismatch (-want +got):\n%s", diff)
		}
	})
}
