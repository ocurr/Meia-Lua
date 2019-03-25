package main

import (
	"fmt"
	"github.com/ocurr/senior-project/antlr/parser"
	"strconv"
)

type LuaASTBuilder struct {
	*parser.BaseLuaVisitor
}

func NewLuaASTBuilder() *LuaASTBuilder {
	return new(LuaASTBuilder)
}

func (v *LuaASTBuilder) VisitChunk(ctx *parser.ChunkContext) interface{} {
	return ChunkC{Block: ctx.Block().Accept(v).(BlockC)}
}

func (v *LuaASTBuilder) VisitBlock(ctx *parser.BlockContext) interface{} {
	stats := ctx.AllStat()
	statLst := make([]Stat, len(stats))
	for i, _ := range stats {
		statLst[i] = stats[i].Accept(v).(Stat)
	}
	return BlockC{StatLst: statLst}
}

func (v *LuaASTBuilder) VisitStat(ctx *parser.StatContext) interface{} {

	if t := ctx.Typedvarlist(); t != nil {
		e := ctx.Explist()

		allTVar := t.Accept(v).(IdLst)
		allExp := e.Accept(v).(ExpLst)

		if len(allTVar.List) != len(allExp.List) {
			fmt.Println("ERROR: AST: the var list is not the same length as the expression list")
			return DefLst{}
		}

		lst := make([]DefC, len(allTVar.List))

		for i := 0; i < len(allTVar.List); i++ {
			lst[i] = DefC{Id: allTVar.List[i], Exp: allExp.List[i]}
		}

		return DefLst{List: lst}

	}

	return DefC{}
}

func (v *LuaASTBuilder) VisitRetstat(ctx *parser.RetstatContext) interface{} {
	panic("VisitRetstat not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitLabel(ctx *parser.LabelContext) interface{} {
	panic("VisitLabel not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	panic("VisitFuncname not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	panic("VisitVarlist not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTypedvarlist(ctx *parser.TypedvarlistContext) interface{} {

	allVars := ctx.AllTypedvar()

	allIdC := make([]IdC, len(allVars))

	for i := 0; i < len(allVars); i++ {
		allIdC[i] = allVars[i].Accept(v).(IdC)
	}

	return IdLst{List: allIdC}
}

func (v *LuaASTBuilder) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	panic("VisitNamelist not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitExplist(ctx *parser.ExplistContext) interface{} {

	allExpCtx := ctx.AllExp()

	allExp := make([]Exp, len(allExpCtx))

	for i := 0; i < len(allExpCtx); i++ {
		allExp[i] = allExpCtx[i].Accept(v).(Exp)
	}

	return ExpLst{List: allExp}
}

func (v *LuaASTBuilder) VisitExp(ctx *parser.ExpContext) interface{} {
	if nl := ctx.NumberLiteral(); nl != nil {
		return nl.Accept(v)
	}
	fmt.Println("ERROR: Expression not supported")
	return StringC{S: "EXPRESSION NOT SUPPORTED"}
}

func (v *LuaASTBuilder) VisitTypeLiteral(ctx *parser.TypeLiteralContext) interface{} {
	switch ctx.GetText() {
	case "float":
		return FloatT{}
	case "int":
		return IntT{}
	default:
		fmt.Printf("ERROR type %s is not supported\n", ctx.GetText())
		return ErrorT{}
	}
}

func (v *LuaASTBuilder) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	panic("VisitPrefixexp not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	panic("VisitFunctioncall not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	panic("VisitVarOrExp not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitVarId(ctx *parser.VarIdContext) interface{} {
	return ctx.GetText()
}

func (v *LuaASTBuilder) VisitTypedvar(ctx *parser.TypedvarContext) interface{} {

	return IdC{Id: ctx.VarId().Accept(v).(string), TypeId: ctx.TypeLiteral().Accept(v).(TypeT)}
}

func (v *LuaASTBuilder) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	panic("VisitVarSuffix not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	panic("VisitNameAndArgs not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	panic("VisitArgs not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	panic("VisitFunctiondef not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	panic("VisitFuncbody not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitParlist(ctx *parser.ParlistContext) interface{} {
	panic("VisitParlist not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	panic("VisitTableconstructor not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	panic("VisitFieldlist not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	panic("VisitField not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	panic("VisitFieldsep not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	panic("VisitOperatorOr not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	panic("VisitOperatorAnd not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	panic("VisitOperatorComparison not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	panic("VisitOperatorStrcat not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	panic("VisitOperatorAddSub not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	panic("VisitOperatorMulDivMod not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorBitwise(ctx *parser.OperatorBitwiseContext) interface{} {
	panic("VisitOperatorBitwise not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	panic("VisitOperatorUnary not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	panic("VisitOperatorPower not implemented")
	return v.VisitChildren(ctx)
}

func (v *LuaASTBuilder) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	if INT := ctx.INT(); INT != nil {
		n, err := strconv.ParseInt(INT.GetText(), 10, 64)
		if err != nil {
			fmt.Println("ERROR: unable to parse int")
			return IntC{N: -1}
		}
		return IntC{N: n}
	}
	if FLOAT := ctx.FLOAT(); FLOAT != nil {
		n, err := strconv.ParseFloat(FLOAT.GetText(), 64)
		if err != nil {
			fmt.Println("ERROR: unable to parse float")
			return FloatC{N: -1}
		}
		return FloatC{N: n}
	}
	return StringC{S: "ERROR"}
}

func (v *LuaASTBuilder) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	panic("VisitStringLiteral not implemented")
	return v.VisitChildren(ctx)
}
