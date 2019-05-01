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

	if a := ctx.Assign(); a != nil {
		return a.Accept(v)
	}

	if i := ctx.Ifstat(); i != nil {
		return i.Accept(v)
	}

	return DefC{}
}

func (v *LuaASTBuilder) VisitRetstat(ctx *parser.RetstatContext) interface{} {
	panic("VisitRetstat not implemented")
}

func (v *LuaASTBuilder) VisitLabel(ctx *parser.LabelContext) interface{} {
	panic("VisitLabel not implemented")
}

func (v *LuaASTBuilder) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	panic("VisitFuncname not implemented")
}

func (v *LuaASTBuilder) VisitIfstat(ctx *parser.IfstatContext) interface{} {
	cond := CondC{}
	cond.Cnd = ctx.Exp().Accept(v).(Exp)
	cond.Block = ctx.Block().Accept(v).(BlockC)

	if eifs := ctx.AllElseifstat(); len(eifs) != 0 {
		cond.Elseifs = make([]CondC, len(eifs))
		for i := 0; i < len(eifs); i++ {
			cond.Elseifs[i] = eifs[i].Accept(v).(CondC)
		}
	}

	if els := ctx.Elsestat(); els != nil {
		cond.Else = els.Accept(v).(BlockC)
	}

	return cond
}

func (v *LuaASTBuilder) VisitElseifstat(ctx *parser.IfstatContext) interface{} {
	cond := CondC{}
	cond.Cnd = ctx.Exp().Accept(v).(Exp)
	cond.Block = ctx.Block().Accept(v).(BlockC)
	return cond
}

func (v *LuaASTBuilder) VisitElsestat(ctx *parser.IfstatContext) interface{} {
	return ctx.Block().Accept(v)
}

func (v *LuaASTBuilder) VisitAssign(ctx *parser.AssignContext) interface{} {

	e := ctx.Explist()
	allExp := e.Accept(v).(ExpLst)

	var allTVar IdLst

	t := ctx.Typedvarlist()
	if t == nil {
		t := ctx.Varlist()
		allTVar = t.Accept(v).(IdLst)
	} else {
		allTVar = t.Accept(v).(IdLst)
	}

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

func (v *LuaASTBuilder) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	if allVars := ctx.AllVarId(); len(allVars) != 0 {
		allIdC := make([]IdC, len(allVars))

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = IdC{
				Id:     allVars[i].Accept(v).(string),
				TypeId: nil,
			}
		}

		return IdLst{List: allIdC}
	}

	return IdLst{}
}

func (v *LuaASTBuilder) VisitTypedvarlist(ctx *parser.TypedvarlistContext) interface{} {

	if allVars := ctx.AllTypedvar(); len(allVars) != 0 {
		allIdC := make([]IdC, len(allVars))

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = allVars[i].Accept(v).(IdC)
		}

		return IdLst{List: allIdC}

	} else if allVars := ctx.AllVarId(); len(allVars) != 0 {
		allIdC := make([]IdC, len(allVars))

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = IdC{
				Id:     allVars[i].Accept(v).(string),
				TypeId: ctx.TypeLiteral().Accept(v).(TypeT),
			}
		}

		return IdLst{List: allIdC}
	}

	return IdLst{}
}

func (v *LuaASTBuilder) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	panic("VisitNamelist not implemented")
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
	} else if sl := ctx.StringLiteral(); sl != nil {
		return sl.Accept(v)
	} else if bop := ctx.OperatorAddSub(); bop != nil {
		return BinaryOpC{
			Lhs: ctx.Exp(0).Accept(v).(Exp),
			Rhs: ctx.Exp(1).Accept(v).(Exp),
			Op:  bop.Accept(v).(string),
		}
	} else if bop := ctx.OperatorMulDivMod(); bop != nil {
		return BinaryOpC{
			Lhs: ctx.Exp(0).Accept(v).(Exp),
			Rhs: ctx.Exp(1).Accept(v).(Exp),
			Op:  bop.Accept(v).(string),
		}
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
	case "string":
		return StringT{}
	default:
		fmt.Printf("ERROR type %s is not supported\n", ctx.GetText())
		return ErrorT{}
	}
}

func (v *LuaASTBuilder) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	panic("VisitPrefixexp not implemented")
}

func (v *LuaASTBuilder) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	panic("VisitFunctioncall not implemented")
}

func (v *LuaASTBuilder) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	panic("VisitVarOrExp not implemented")
}

func (v *LuaASTBuilder) VisitVarId(ctx *parser.VarIdContext) interface{} {
	return ctx.GetText()
}

func (v *LuaASTBuilder) VisitTypedvar(ctx *parser.TypedvarContext) interface{} {

	return IdC{
		Id:     ctx.VarId().Accept(v).(string),
		TypeId: ctx.TypeLiteral().Accept(v).(TypeT),
	}
}

func (v *LuaASTBuilder) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	panic("VisitVarSuffix not implemented")
}

func (v *LuaASTBuilder) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	panic("VisitNameAndArgs not implemented")
}

func (v *LuaASTBuilder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	panic("VisitArgs not implemented")
}

func (v *LuaASTBuilder) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	panic("VisitFunctiondef not implemented")
}

func (v *LuaASTBuilder) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	panic("VisitFuncbody not implemented")
}

func (v *LuaASTBuilder) VisitParlist(ctx *parser.ParlistContext) interface{} {
	panic("VisitParlist not implemented")
}

func (v *LuaASTBuilder) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	panic("VisitTableconstructor not implemented")
}

func (v *LuaASTBuilder) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	panic("VisitFieldlist not implemented")
}

func (v *LuaASTBuilder) VisitField(ctx *parser.FieldContext) interface{} {
	panic("VisitField not implemented")
}

func (v *LuaASTBuilder) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	panic("VisitFieldsep not implemented")
}

func (v *LuaASTBuilder) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	panic("VisitOperatorOr not implemented")
}

func (v *LuaASTBuilder) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	panic("VisitOperatorAnd not implemented")
}

func (v *LuaASTBuilder) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	panic("VisitOperatorComparison not implemented")
}

func (v *LuaASTBuilder) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	panic("VisitOperatorStrcat not implemented")
}

func (v *LuaASTBuilder) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	return ctx.GetText()
}

func (v *LuaASTBuilder) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	return ctx.GetText()
}

func (v *LuaASTBuilder) VisitOperatorBitwise(ctx *parser.OperatorBitwiseContext) interface{} {
	panic("VisitOperatorBitwise not implemented")
}

func (v *LuaASTBuilder) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	panic("VisitOperatorUnary not implemented")
}

func (v *LuaASTBuilder) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	panic("VisitOperatorPower not implemented")
}

func (v *LuaASTBuilder) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	b := BoolC{}
	switch ctx.GetText() {
	case "true":
		b.True = true
	case "false":
		b.True = false
	}

	return b
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
	if normStr := ctx.NORMALSTRING(); normStr != nil {
		str, err := strconv.Unquote(normStr.GetText())
		if err != nil {
			fmt.Println("ERROR: unable to parse string")
		}
		return StringC{S: str}
	} else if charStr := ctx.CHARSTRING(); charStr != nil {
		str, err := strconv.Unquote(charStr.GetText())
		if err != nil {
			fmt.Println("ERROR: unable to parse string")
		}
		return StringC{S: str}
	} else if longStr := ctx.LONGSTRING(); longStr != nil {
		str, err := strconv.Unquote(longStr.GetText())
		if err != nil {
			fmt.Println("ERROR: unable to parse string")
		}
		return StringC{S: str}
	}

	fmt.Println("ERROR: unable to parse string literal")
	return StringC{S: "ERROR"}
}
