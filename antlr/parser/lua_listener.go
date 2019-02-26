// Code generated from Lua.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Lua

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LuaListener is a complete listener for a parse tree produced by LuaParser.
type LuaListener interface {
	antlr.ParseTreeListener

	// EnterChunk is called when entering the chunk production.
	EnterChunk(c *ChunkContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// EnterStat is called when entering the stat production.
	EnterStat(c *StatContext)

	// EnterRetstat is called when entering the retstat production.
	EnterRetstat(c *RetstatContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterFuncname is called when entering the funcname production.
	EnterFuncname(c *FuncnameContext)

	// EnterVarlist is called when entering the varlist production.
	EnterVarlist(c *VarlistContext)

	// EnterTypedvarlist is called when entering the typedvarlist production.
	EnterTypedvarlist(c *TypedvarlistContext)

	// EnterNamelist is called when entering the namelist production.
	EnterNamelist(c *NamelistContext)

	// EnterExplist is called when entering the explist production.
	EnterExplist(c *ExplistContext)

	// EnterExp is called when entering the exp production.
	EnterExp(c *ExpContext)

	// EnterTypeLiteral is called when entering the typeLiteral production.
	EnterTypeLiteral(c *TypeLiteralContext)

	// EnterPrefixexp is called when entering the prefixexp production.
	EnterPrefixexp(c *PrefixexpContext)

	// EnterFunctioncall is called when entering the functioncall production.
	EnterFunctioncall(c *FunctioncallContext)

	// EnterVarOrExp is called when entering the varOrExp production.
	EnterVarOrExp(c *VarOrExpContext)

	// EnterVarId is called when entering the varId production.
	EnterVarId(c *VarIdContext)

	// EnterTypedvar is called when entering the typedvar production.
	EnterTypedvar(c *TypedvarContext)

	// EnterVarSuffix is called when entering the varSuffix production.
	EnterVarSuffix(c *VarSuffixContext)

	// EnterNameAndArgs is called when entering the nameAndArgs production.
	EnterNameAndArgs(c *NameAndArgsContext)

	// EnterArgs is called when entering the args production.
	EnterArgs(c *ArgsContext)

	// EnterFunctiondef is called when entering the functiondef production.
	EnterFunctiondef(c *FunctiondefContext)

	// EnterFuncbody is called when entering the funcbody production.
	EnterFuncbody(c *FuncbodyContext)

	// EnterParlist is called when entering the parlist production.
	EnterParlist(c *ParlistContext)

	// EnterTableconstructor is called when entering the tableconstructor production.
	EnterTableconstructor(c *TableconstructorContext)

	// EnterFieldlist is called when entering the fieldlist production.
	EnterFieldlist(c *FieldlistContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterFieldsep is called when entering the fieldsep production.
	EnterFieldsep(c *FieldsepContext)

	// EnterOperatorOr is called when entering the operatorOr production.
	EnterOperatorOr(c *OperatorOrContext)

	// EnterOperatorAnd is called when entering the operatorAnd production.
	EnterOperatorAnd(c *OperatorAndContext)

	// EnterOperatorComparison is called when entering the operatorComparison production.
	EnterOperatorComparison(c *OperatorComparisonContext)

	// EnterOperatorStrcat is called when entering the operatorStrcat production.
	EnterOperatorStrcat(c *OperatorStrcatContext)

	// EnterOperatorAddSub is called when entering the operatorAddSub production.
	EnterOperatorAddSub(c *OperatorAddSubContext)

	// EnterOperatorMulDivMod is called when entering the operatorMulDivMod production.
	EnterOperatorMulDivMod(c *OperatorMulDivModContext)

	// EnterOperatorBitwise is called when entering the operatorBitwise production.
	EnterOperatorBitwise(c *OperatorBitwiseContext)

	// EnterOperatorUnary is called when entering the operatorUnary production.
	EnterOperatorUnary(c *OperatorUnaryContext)

	// EnterOperatorPower is called when entering the operatorPower production.
	EnterOperatorPower(c *OperatorPowerContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterStringLiteral is called when entering the stringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// ExitChunk is called when exiting the chunk production.
	ExitChunk(c *ChunkContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)

	// ExitStat is called when exiting the stat production.
	ExitStat(c *StatContext)

	// ExitRetstat is called when exiting the retstat production.
	ExitRetstat(c *RetstatContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitFuncname is called when exiting the funcname production.
	ExitFuncname(c *FuncnameContext)

	// ExitVarlist is called when exiting the varlist production.
	ExitVarlist(c *VarlistContext)

	// ExitTypedvarlist is called when exiting the typedvarlist production.
	ExitTypedvarlist(c *TypedvarlistContext)

	// ExitNamelist is called when exiting the namelist production.
	ExitNamelist(c *NamelistContext)

	// ExitExplist is called when exiting the explist production.
	ExitExplist(c *ExplistContext)

	// ExitExp is called when exiting the exp production.
	ExitExp(c *ExpContext)

	// ExitTypeLiteral is called when exiting the typeLiteral production.
	ExitTypeLiteral(c *TypeLiteralContext)

	// ExitPrefixexp is called when exiting the prefixexp production.
	ExitPrefixexp(c *PrefixexpContext)

	// ExitFunctioncall is called when exiting the functioncall production.
	ExitFunctioncall(c *FunctioncallContext)

	// ExitVarOrExp is called when exiting the varOrExp production.
	ExitVarOrExp(c *VarOrExpContext)

	// ExitVarId is called when exiting the varId production.
	ExitVarId(c *VarIdContext)

	// ExitTypedvar is called when exiting the typedvar production.
	ExitTypedvar(c *TypedvarContext)

	// ExitVarSuffix is called when exiting the varSuffix production.
	ExitVarSuffix(c *VarSuffixContext)

	// ExitNameAndArgs is called when exiting the nameAndArgs production.
	ExitNameAndArgs(c *NameAndArgsContext)

	// ExitArgs is called when exiting the args production.
	ExitArgs(c *ArgsContext)

	// ExitFunctiondef is called when exiting the functiondef production.
	ExitFunctiondef(c *FunctiondefContext)

	// ExitFuncbody is called when exiting the funcbody production.
	ExitFuncbody(c *FuncbodyContext)

	// ExitParlist is called when exiting the parlist production.
	ExitParlist(c *ParlistContext)

	// ExitTableconstructor is called when exiting the tableconstructor production.
	ExitTableconstructor(c *TableconstructorContext)

	// ExitFieldlist is called when exiting the fieldlist production.
	ExitFieldlist(c *FieldlistContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitFieldsep is called when exiting the fieldsep production.
	ExitFieldsep(c *FieldsepContext)

	// ExitOperatorOr is called when exiting the operatorOr production.
	ExitOperatorOr(c *OperatorOrContext)

	// ExitOperatorAnd is called when exiting the operatorAnd production.
	ExitOperatorAnd(c *OperatorAndContext)

	// ExitOperatorComparison is called when exiting the operatorComparison production.
	ExitOperatorComparison(c *OperatorComparisonContext)

	// ExitOperatorStrcat is called when exiting the operatorStrcat production.
	ExitOperatorStrcat(c *OperatorStrcatContext)

	// ExitOperatorAddSub is called when exiting the operatorAddSub production.
	ExitOperatorAddSub(c *OperatorAddSubContext)

	// ExitOperatorMulDivMod is called when exiting the operatorMulDivMod production.
	ExitOperatorMulDivMod(c *OperatorMulDivModContext)

	// ExitOperatorBitwise is called when exiting the operatorBitwise production.
	ExitOperatorBitwise(c *OperatorBitwiseContext)

	// ExitOperatorUnary is called when exiting the operatorUnary production.
	ExitOperatorUnary(c *OperatorUnaryContext)

	// ExitOperatorPower is called when exiting the operatorPower production.
	ExitOperatorPower(c *OperatorPowerContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitStringLiteral is called when exiting the stringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)
}
