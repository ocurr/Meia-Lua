package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/senior-project/antlr/parser"
)

type LuaASTBuilder struct {
	*antlr.BaseParseTreeVisitor
}

func NewLuaASTBuilder() *LuaASTBuilder {
	return new(LuaASTBuilder)
}

func (v *LuaASTBuilder) VisitChunk(ctx *parser.ChunkContext) interface{} {
	return ChunkC{Block: v.Visit(ctx.Block())}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	stats := ctx.AllStat()
	statc := make([]StatC, len(stats))
	for i, s := range stats {
		statc[i] = v.Visit(stats[i])
	}
	return BlockC{StatLst: statc}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitStat(ctx *parser.StatContext) interface{} {

	if t := ctx.Typedvarlist(); t != nil {
		e := ctx.Explist()
		if len(t) != len(e) {
			fmt.Println("ERROR: AST: the var list is not the same length as the expression list")
		}

		l := make([]DefC, len(t))
		for i, v := range t {
			//l[i] = DefC{Id: v.Visit(v
		}
		return StatC{}

	}
	//return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitRetstat(ctx *parser.RetstatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitLabel(ctx *parser.LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypedvarlist(ctx *parser.TypedvarlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitExplist(ctx *parser.ExplistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitExp(ctx *parser.ExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypeLiteral(ctx *parser.TypeLiteralContext) interface{} {
	return ctx.getText()
}

func (v *LuaASTBuilder) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarId(ctx *parser.VarIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypedvar(ctx *parser.TypedvarContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitParlist(ctx *parser.ParlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorBitwise(ctx *parser.OperatorBitwiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
