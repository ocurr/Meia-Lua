package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
	"reflect"
	"testing"
)

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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: IntC{N: 5}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: FloatT{}}, Exp: FloatC{N: 5.5}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("SingleStringAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("string animal = \"horse\"")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "animal", TypeId: StringT{}}, Exp: StringC{S: "horse"}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
				DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: IntC{N: 5}},
				DefC{Id: IdC{Id: "y", TypeId: IntT{}}, Exp: IntC{N: 6}},
				DefC{Id: IdC{Id: "z", TypeId: IntT{}}, Exp: IntC{N: 7}},
			}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected:i %#v ; got: %#v", want, got)
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
				DefC{Id: IdC{Id: "x", TypeId: FloatT{}}, Exp: FloatC{N: 5.5}},
				DefC{Id: IdC{Id: "y", TypeId: FloatT{}}, Exp: FloatC{N: 10.789}},
				DefC{Id: IdC{Id: "z", TypeId: FloatT{}}, Exp: FloatC{N: 100.23}},
			}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("MultipleStringAssign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("string animal, dog, word = \"horse\", \"fido\", \"bar\"")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{DefLst{List: []DefC{
				DefC{Id: IdC{Id: "animal", TypeId: StringT{}}, Exp: StringC{S: "horse"}},
				DefC{Id: IdC{Id: "dog", TypeId: StringT{}}, Exp: StringC{S: "fido"}},
				DefC{Id: IdC{Id: "word", TypeId: StringT{}}, Exp: StringC{S: "bar"}},
			}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "+",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "-",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "*",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "%",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "/",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BoolC{True: true}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "<",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "<=",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  ">",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  ">=",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "==",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: BoolT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "~=",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
				Lhs: IntC{N: 5},
				Rhs: IntC{N: 2},
				Op:  "//",
			}}}}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("Reassign", func(t *testing.T) {
		inputStream := antlr.NewInputStream("int x = 5 // 2\nx = 1 + 1")
		lexer := parser.NewLuaLexer(inputStream)
		tokenStream := antlr.NewCommonTokenStream(lexer, 0)
		p := parser.NewLuaParser(tokenStream)
		tree := p.Chunk()
		p.BuildParseTrees = true
		builder := NewLuaASTBuilder()

		got := tree.Accept(builder)
		want := ChunkC{Block: BlockC{
			StatLst: []Stat{
				DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: BinaryOpC{
					Lhs: IntC{N: 5},
					Rhs: IntC{N: 2},
					Op:  "//",
				}}}},
				DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: nil}, Exp: BinaryOpC{
					Lhs: IntC{N: 1},
					Rhs: IntC{N: 1},
					Op:  "+",
				}}}},
			}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("IfOnly", func(t *testing.T) {
		inputStream := antlr.NewInputStream("if true then\nx = 5\nend")
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
						DefLst{List: []DefC{DefC{Id: IdC{Id: "x"}, Exp: IntC{N: 5}}}}},
					},
				}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("WhileOnly", func(t *testing.T) {
		inputStream := antlr.NewInputStream("while true do\nx = 5\nend")
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
						DefLst{List: []DefC{DefC{Id: IdC{Id: "x"}, Exp: IntC{N: 5}}}}},
					},
				}}}}
		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("For", func(t *testing.T) {
		inputStream := antlr.NewInputStream("for int x = 5, x < 5, 1 do\ny = 5\nend")
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
						TypeId: IntT{},
					},
					Exp: IntC{N: 5},
				},
				Cnd: BinaryOpC{
					Lhs: IdC{Id: "x"},
					Rhs: IntC{N: 5},
					Op:  "<",
				},
				Step:  IntC{N: 1},
				Block: BlockC{StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "y"}, Exp: IntC{N: 5}}}}}},
			},
		}}}

		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
	t.Run("ForNoStep", func(t *testing.T) {
		inputStream := antlr.NewInputStream("for int x = 5, x < 5 do\ny = 5\nend")
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
						TypeId: IntT{},
					},
					Exp: IntC{N: 5},
				},
				Cnd: BinaryOpC{
					Lhs: IdC{Id: "x"},
					Rhs: IntC{N: 5},
					Op:  "<",
				},
				Step:  IntC{N: 1},
				Block: BlockC{StatLst: []Stat{DefLst{List: []DefC{DefC{Id: IdC{Id: "y"}, Exp: IntC{N: 5}}}}}},
			},
		}}}

		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
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
			DefLst{List: []DefC{DefC{Id: IdC{Id: "x", TypeId: IntT{}}, Exp: NilC{}}}},
		}}}

		res := astMatch(got.(ChunkC), want)
		if !res {
			t.Errorf("expected: %#v ; got: %#v", want, got)
		}
	})
}

func TestASTMatch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		A := IdC{Id: "x", TypeId: IntT{}}
		B := IdC{Id: "x", TypeId: IntT{}}
		res := astMatch(A, B)
		if !res {
			t.Errorf("%#v and %#v should be equal", A, B)
		}
	})
	t.Run("2", func(t *testing.T) {
		A := IdC{Id: "y", TypeId: IntT{}}
		B := IdC{Id: "x", TypeId: IntT{}}
		res := astMatch(A, B)
		if res {
			t.Errorf("%#v and %#v should NOT be equal", A, B)
		}
	})
	t.Run("3", func(t *testing.T) {
		A := IdC{Id: "y", TypeId: IntT{}}
		B := IntC{N: 5}
		res := astMatch(A, B)
		if res {
			t.Errorf("%#v and %#v should NOT be equal", A, B)
		}
	})
	t.Run("4", func(t *testing.T) {
		A := ExpLst{List: []Exp{IntC{N: 5}, IntC{N: 10}, FloatC{N: 13.5}}}
		B := ExpLst{List: []Exp{IntC{N: 5}, IntC{N: 10}, FloatC{N: 13.5}}}
		res := astMatch(A, B)
		if !res {
			t.Errorf("%#v and %#v should be equal", A, B)
		}
	})
	t.Run("5", func(t *testing.T) {
		A := BlockC{StatLst: []Stat{DefC{Id: IdC{Id: "y", TypeId: IntT{}}, Exp: IntC{N: 5}}}}
		B := BlockC{StatLst: []Stat{DefC{Id: IdC{Id: "y", TypeId: IntT{}}, Exp: IntC{N: 5}}}}
		res := astMatch(A, B)
		if !res {
			t.Errorf("%#v and %#v should be equal", A, B)
		}
	})
}

func astMatch(node1, node2 Node) bool {
	if reflect.TypeOf(node1) != reflect.TypeOf(node2) {
		return false
	}

	if node1 == nil {
		if node2 == nil {
			return true
		}
		return false
	}

	switch n1 := node1.(type) {
	case ChunkC:
		n2 := node2.(ChunkC)
		return astMatch(n1.Block, n2.Block)
	case BlockC:
		n2 := node2.(BlockC)
		if len(n1.StatLst) != len(n2.StatLst) {
			return false
		}

		for i := 0; i < len(n1.StatLst); i++ {
			if !astMatch(n1.StatLst[i], n2.StatLst[i]) {
				return false
			}
		}

		return true
	case DefC:
		n2 := node2.(DefC)
		return astMatch(n1.Id, n2.Id) && astMatch(n1.Exp, n2.Exp)
	case IdC:
		n2 := node2.(IdC)
		return n1.Id == n2.Id && n1.TypeId == n2.TypeId
	case ExpLst:
		n2 := node2.(ExpLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case DefLst:
		n2 := node2.(DefLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case IdLst:
		n2 := node2.(IdLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case BinaryOpC:
		n2 := node2.(BinaryOpC)
		return astMatch(n1.Lhs, n2.Lhs) &&
			astMatch(n1.Rhs, n2.Rhs) &&
			n1.Op == n2.Op
	case CondC:
		n2 := node2.(CondC)
		match := true
		if len(n1.Elseifs) != len(n2.Elseifs) {
			return false
		} else {
			for i := 0; i < len(n1.Elseifs); i++ {
				match = astMatch(n1.Elseifs[i], n2.Elseifs[i])
			}
		}
		return astMatch(n1.Cnd, n2.Cnd) &&
			astMatch(n1.Block, n2.Block) && match && astMatch(n1.Else, n2.Else)
	case WhileC:
		n2 := node2.(WhileC)
		return astMatch(n1.Cnd, n2.Cnd) && astMatch(n1.Block, n2.Block)
	case ForC:
		n2 := node2.(ForC)

		return astMatch(n1.Assign, n2.Assign) && astMatch(n1.Cnd, n2.Cnd) &&
			astMatch(n1.Step, n2.Step) && astMatch(n1.Block, n2.Block)
	case IntC:
		n2 := node2.(IntC)
		return n1.N == n2.N
	case FloatC:
		n2 := node2.(FloatC)
		return n1.N == n2.N
	case StringC:
		n2 := node2.(StringC)
		return n1.S == n2.S
	case BoolC:
		n2 := node2.(BoolC)
		return n1.True == n2.True
	case NilC:
		return true
	default:
		panic("AST Node in ASTMatch not implemented")
	}
}
