package ast

import (
	"fmt"
	"strconv"

	"github.com/ocurr/Meia-Lua/parser"
	"github.com/ocurr/Meia-Lua/types"
)

// Builder builds an abstract syntax tree from parse contexts.
type Builder struct {
	parser.BaseLuaVisitor
}

// NewBuilder returns a new Builder.
func NewBuilder() *Builder {
	return new(Builder)
}

// VisitChunk visits a Chunk.
func (v *Builder) VisitChunk(ctx *parser.ChunkContext) interface{} {
	return Chunk{Ctx: ctx, Block: ctx.Block().Accept(v).(Block)}
}

// VisitBlock visits a Block.
func (v *Builder) VisitBlock(ctx *parser.BlockContext) interface{} {
	stats := ctx.AllStat()
	statLst := make([]Stat, len(stats))
	for i := range stats {
		statLst[i] = stats[i].Accept(v).(Stat)
	}
	return Block{Ctx: ctx, StatLst: statLst}
}

// VisitStat visits a Stat.
func (v *Builder) VisitStat(ctx *parser.StatContext) interface{} {

	if a := ctx.Assign(); a != nil {
		return a.Accept(v)
	}

	if i := ctx.Ifstat(); i != nil {
		return i.Accept(v)
	}

	if w := ctx.Whilestat(); w != nil {
		return w.Accept(v)
	}

	if f := ctx.Forstat(); f != nil {
		return f.Accept(v)
	}

	return Def{Ctx: ctx}
}

// VisitRetstat visits a Retstat.
func (v *Builder) VisitRetstat(ctx *parser.RetstatContext) interface{} {
	panic("VisitRetstat not implemented")
}

// VisitLabel visits a label.
func (v *Builder) VisitLabel(ctx *parser.LabelContext) interface{} {
	panic("VisitLabel not implemented")
}

// VisitFuncname visits a funcname.
func (v *Builder) VisitFuncname(ctx *parser.FuncnameContext) interface{} {
	panic("VisitFuncname not implemented")
}

// VisitIfstat visits an ifstat.
func (v *Builder) VisitIfstat(ctx *parser.IfstatContext) interface{} {
	cond := Cond{}
	cond.Cnd = ctx.Exp().Accept(v).(Exp)
	cond.Block = ctx.Block().Accept(v).(Block)

	if eifs := ctx.AllElseifstat(); len(eifs) != 0 {
		cond.Elseifs = make([]Cond, len(eifs))
		for i := 0; i < len(eifs); i++ {
			cond.Elseifs[i] = eifs[i].Accept(v).(Cond)
		}
	}

	if els := ctx.Elsestat(); els != nil {
		cond.Else = els.Accept(v).(Block)
	}

	return cond
}

// VisitElseifstat visits an elseifstat.
func (v *Builder) VisitElseifstat(ctx *parser.ElseifstatContext) interface{} {
	cond := Cond{Ctx: ctx}
	cond.Cnd = ctx.Exp().Accept(v).(Exp)
	cond.Block = ctx.Block().Accept(v).(Block)
	return cond
}

// VisitElsestat visits an elsestat.
func (v *Builder) VisitElsestat(ctx *parser.ElsestatContext) interface{} {
	return ctx.Block().Accept(v)
}

// VisitAssign visits an assign.
func (v *Builder) VisitAssign(ctx *parser.AssignContext) interface{} {

	var allExp ExpLst
	e := ctx.Explist()
	if e != nil {
		allExp = e.Accept(v).(ExpLst)
	} else {
		allExp = ExpLst{Ctx: ctx, List: []Exp{}}
	}

	var allTVar IdLst

	t := ctx.Typedvarlist()
	if t == nil {
		t := ctx.Varlist()
		allTVar = t.Accept(v).(IdLst)
	} else {
		allTVar = t.Accept(v).(IdLst)
	}

	if len(allExp.List) != 0 && len(allTVar.List) != len(allExp.List) {
		fmt.Println("ERROR: AST: the var list is not the same length as the expression list")
		return DefLst{}
	}

	lst := make([]Def, len(allTVar.List))

	scope := GLOBAL
	if len(allExp.List) == 0 {
		scope = LOCAL
	}

	for i := 0; i < len(allTVar.List); i++ {
		lst[i] = Def{
			Ctx: ctx,
			Id:  allTVar.List[i],
			Exp: func() Exp {
				if len(allExp.List) == 0 {
					return Nil{}
				}
				return allExp.List[i]
			}(),
			Scope: scope,
		}
	}

	return DefLst{Ctx: ctx, List: lst}
}

// VisitWhilestat visits a whilestat.
func (v *Builder) VisitWhilestat(ctx *parser.WhilestatContext) interface{} {
	return While{
		Ctx:   ctx,
		Cnd:   ctx.Exp().Accept(v).(Exp),
		Block: ctx.Block().Accept(v).(Block),
	}
}

// VisitForstat visits a forstat.
func (v *Builder) VisitForstat(ctx *parser.ForstatContext) interface{} {

	forS := For{
		Ctx:    ctx,
		Assign: Def{Id: Id{}, Scope: GLOBAL},
	}

	if typeL := ctx.TypeLiteral(); typeL != nil {
		forS.Assign.Id.TypeId = typeL.Accept(v).(types.Type)
	}

	forS.Assign.Id.Id = ctx.NAME().GetText()

	allExp := ctx.AllExp()

	forS.Assign.Exp = allExp[0].Accept(v).(Exp)

	forS.Cnd = allExp[1].Accept(v).(Exp)

	if len(allExp) > 2 {
		forS.Step = allExp[2].Accept(v).(Exp)
	} else {
		forS.Step = Int{Ctx: ctx, N: 1}
	}

	forS.Block = ctx.Block().Accept(v).(Block)

	return forS
}

// VisitVarlist visits a varlist.
func (v *Builder) VisitVarlist(ctx *parser.VarlistContext) interface{} {
	if allVars := ctx.AllVarId(); len(allVars) != 0 {
		allIdC := make([]Id, len(allVars))

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = Id{
				Ctx:    ctx,
				Id:     allVars[i].Accept(v).(string),
				TypeId: nil,
			}
		}

		return IdLst{List: allIdC}
	}

	return IdLst{}
}

// VisitTypedvarlist visits a typedvarlist.
func (v *Builder) VisitTypedvarlist(ctx *parser.TypedvarlistContext) interface{} {

	if allVars := ctx.AllTypedvar(); len(allVars) != 0 {
		allIdC := make([]Id, len(allVars))

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = allVars[i].Accept(v).(Id)
		}

		return IdLst{Ctx: ctx, List: allIdC}

	} else if allVars := ctx.AllVarId(); len(allVars) != 0 {
		allIdC := make([]Id, len(allVars))
		listType := ctx.TypeLiteral().Accept(v).(types.Type)

		for i := 0; i < len(allVars); i++ {
			allIdC[i] = Id{
				Ctx:    ctx,
				Id:     allVars[i].Accept(v).(string),
				TypeId: listType,
			}
		}

		return IdLst{Ctx: ctx, List: allIdC}
	}

	return IdLst{}
}

// VisitNamelist visits a namelist.
func (v *Builder) VisitNamelist(ctx *parser.NamelistContext) interface{} {
	panic("VisitNamelist not implemented")
}

// VisitExplist visits an explist.
func (v *Builder) VisitExplist(ctx *parser.ExplistContext) interface{} {

	allExpCtx := ctx.AllExp()

	allExp := make([]Exp, len(allExpCtx))

	for i := 0; i < len(allExpCtx); i++ {
		allExp[i] = allExpCtx[i].Accept(v).(Exp)
	}

	return ExpLst{Ctx: ctx, List: allExp}
}

// VisitExp visits an exp.
func (v *Builder) VisitExp(ctx *parser.ExpContext) interface{} {
	if nl := ctx.NumberLiteral(); nl != nil {
		return nl.Accept(v)
	} else if sl := ctx.StringLiteral(); sl != nil {
		return sl.Accept(v)
	} else if bl := ctx.BoolLiteral(); bl != nil {
		return bl.Accept(v)
	} else if bop := ctx.OperatorAddSub(); bop != nil {
		return BinaryOp{
			Ctx: ctx,
			Lhs: ctx.Exp(0).Accept(v).(Exp),
			Rhs: ctx.Exp(1).Accept(v).(Exp),
			Op:  bop.Accept(v).(string),
		}
	} else if bop := ctx.OperatorMulDivMod(); bop != nil {
		return BinaryOp{
			Ctx: ctx,
			Lhs: ctx.Exp(0).Accept(v).(Exp),
			Rhs: ctx.Exp(1).Accept(v).(Exp),
			Op:  bop.Accept(v).(string),
		}
	} else if bop := ctx.OperatorComparison(); bop != nil {
		return BinaryOp{
			Ctx: ctx,
			Lhs: ctx.Exp(0).Accept(v).(Exp),
			Rhs: ctx.Exp(1).Accept(v).(Exp),
			Op:  bop.Accept(v).(string),
		}
	} else if pref := ctx.Prefixexp(); pref != nil {
		return pref.Accept(v).(Exp)
	}

	switch ctx.GetText() {
	case "nil":
		return Nil{}
	}

	fmt.Println("ERROR: Expression not supported")
	return String{S: "EXPRESSION NOT SUPPORTED"}
}

// VisitTypeLiteral visits a typeLiteral.
func (v *Builder) VisitTypeLiteral(ctx *parser.TypeLiteralContext) interface{} {
	switch ctx.GetText() {
	case "float":
		return types.Float{}
	case "int":
		return types.Int{}
	case "string":
		return types.String{}
	case "bool":
		return types.Bool{}
	default:
		fmt.Printf("ERROR type %s is not supported\n", ctx.GetText())
		return types.Error{}
	}
}

// VisitPrefixexp visits a prefixexp.
func (v *Builder) VisitPrefixexp(ctx *parser.PrefixexpContext) interface{} {
	return ctx.VarOrExp().Accept(v)
}

// VisitFunctioncall visits a functioncall.
func (v *Builder) VisitFunctioncall(ctx *parser.FunctioncallContext) interface{} {
	panic("VisitFunctioncall not implemented")
}

// VisitVarOrExp visits a varOrExp.
func (v *Builder) VisitVarOrExp(ctx *parser.VarOrExpContext) interface{} {
	if i := ctx.VarId(); i != nil {
		return Id{Ctx: ctx, Id: i.Accept(v).(string)}
	}
	if e := ctx.Exp(); e != nil {
		return e.Accept(v)
	}

	// we can't get here since this expression will consist of either a varid or an exp
	return nil
}

// VisitVarId visits a varId.
func (v *Builder) VisitVarId(ctx *parser.VarIdContext) interface{} {
	return ctx.GetText()
}

// VisitTypedvar visits a typedvar.
func (v *Builder) VisitTypedvar(ctx *parser.TypedvarContext) interface{} {

	return Id{
		Ctx:    ctx,
		Id:     ctx.VarId().Accept(v).(string),
		TypeId: ctx.TypeLiteral().Accept(v).(types.Type),
	}
}

func (v *Builder) VisitVarSuffix(ctx *parser.VarSuffixContext) interface{} {
	panic("VisitVarSuffix not implemented")
}

func (v *Builder) VisitNameAndArgs(ctx *parser.NameAndArgsContext) interface{} {
	panic("VisitNameAndArgs not implemented")
}

func (v *Builder) VisitArgs(ctx *parser.ArgsContext) interface{} {
	panic("VisitArgs not implemented")
}

func (v *Builder) VisitFunctiondef(ctx *parser.FunctiondefContext) interface{} {
	panic("VisitFunctiondef not implemented")
}

func (v *Builder) VisitFuncbody(ctx *parser.FuncbodyContext) interface{} {
	panic("VisitFuncbody not implemented")
}

func (v *Builder) VisitParlist(ctx *parser.ParlistContext) interface{} {
	panic("VisitParlist not implemented")
}

func (v *Builder) VisitTableconstructor(ctx *parser.TableconstructorContext) interface{} {
	panic("VisitTableconstructor not implemented")
}

func (v *Builder) VisitFieldlist(ctx *parser.FieldlistContext) interface{} {
	panic("VisitFieldlist not implemented")
}

func (v *Builder) VisitField(ctx *parser.FieldContext) interface{} {
	panic("VisitField not implemented")
}

func (v *Builder) VisitFieldsep(ctx *parser.FieldsepContext) interface{} {
	panic("VisitFieldsep not implemented")
}

func (v *Builder) VisitOperatorOr(ctx *parser.OperatorOrContext) interface{} {
	panic("VisitOperatorOr not implemented")
}

func (v *Builder) VisitOperatorAnd(ctx *parser.OperatorAndContext) interface{} {
	panic("VisitOperatorAnd not implemented")
}

func (v *Builder) VisitOperatorComparison(ctx *parser.OperatorComparisonContext) interface{} {
	return ctx.GetText()
}

func (v *Builder) VisitOperatorStrcat(ctx *parser.OperatorStrcatContext) interface{} {
	panic("VisitOperatorStrcat not implemented")
}

func (v *Builder) VisitOperatorAddSub(ctx *parser.OperatorAddSubContext) interface{} {
	return ctx.GetText()
}

func (v *Builder) VisitOperatorMulDivMod(ctx *parser.OperatorMulDivModContext) interface{} {
	return ctx.GetText()
}

func (v *Builder) VisitOperatorBitwise(ctx *parser.OperatorBitwiseContext) interface{} {
	panic("VisitOperatorBitwise not implemented")
}

func (v *Builder) VisitOperatorUnary(ctx *parser.OperatorUnaryContext) interface{} {
	panic("VisitOperatorUnary not implemented")
}

func (v *Builder) VisitOperatorPower(ctx *parser.OperatorPowerContext) interface{} {
	panic("VisitOperatorPower not implemented")
}

func (v *Builder) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {
	b := Bool{Ctx: ctx}
	switch ctx.GetText() {
	case "true":
		b.True = true
	case "false":
		b.True = false
	}

	return b
}

func (v *Builder) VisitNumberLiteral(ctx *parser.NumberLiteralContext) interface{} {
	if INT := ctx.INT(); INT != nil {
		n, err := strconv.ParseInt(INT.GetText(), 10, 64)
		if err != nil {
			fmt.Println("ERROR: unable to parse int")
			return Int{N: -1}
		}
		return Int{Ctx: ctx, N: n}
	}
	if FLOAT := ctx.FLOAT(); FLOAT != nil {
		n, err := strconv.ParseFloat(FLOAT.GetText(), 64)
		if err != nil {
			fmt.Println("ERROR: unable to parse float")
			return Float{N: -1}
		}
		return Float{Ctx: ctx, N: n}
	}
	return String{S: "ERROR"}
}

func (v *Builder) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	str, err := strconv.Unquote(ctx.GetText())
	if err != nil {
		fmt.Println("ERROR: unable to parse string", ctx.GetText())
	}
	return String{Ctx: ctx, S: str}
}
