package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/Meia-Lua/types"
)

// Node represents a node in the abstact syntax tree.
type Node interface {
	GetCtx() antlr.ParserRuleContext
}

// Exp represents an expression node in the AST.
type Exp interface {
	Node
	expNode()
}

// Stat represents a statement node in the AST.
type Stat interface {
	Node
	statNode()
}

// ScopeC represents the scope of a statement in the AST.
type ScopeC string

const (
	// LOCAL represents a locally scoped statement.
	LOCAL ScopeC = "local"
	// GLOBAL represents a globally scoped statement.
	GLOBAL ScopeC = "global"
)

// Chunk is a Chunk node in the AST.
type Chunk struct {
	Ctx   antlr.ParserRuleContext
	Block Block
}

// GetCtx returns the context of the ChunkC.
func (c Chunk) GetCtx() antlr.ParserRuleContext {
	return c.Ctx
}

// Block is a Block node in the AST.
type Block struct {
	Ctx     antlr.ParserRuleContext
	StatLst []Stat
}

// GetCtx returns the context of the BlockC.
func (b Block) GetCtx() antlr.ParserRuleContext {
	return b.Ctx
}

// Cond is a conditional statement node in the AST.
type Cond struct {
	Ctx     antlr.ParserRuleContext
	Cnd     Exp
	Block   Block
	Elseifs []Cond
	Else    Block
}

func (c Cond) statNode() {}

// GetCtx returns the context of the CondC.
func (c Cond) GetCtx() antlr.ParserRuleContext {
	return c.Ctx
}

// While is a while loop node in the AST.
type While struct {
	Ctx   antlr.ParserRuleContext
	Cnd   Exp
	Block Block
}

func (w While) statNode() {}

// GetCtx returns the context of the WhileC.
func (w While) GetCtx() antlr.ParserRuleContext {
	return w.Ctx
}

// For is a for loop node in the AST.
type For struct {
	Ctx    antlr.ParserRuleContext
	Assign Def
	Cnd    Exp
	Step   Exp
	Block  Block
}

func (f For) statNode() {}

// GetCtx returns the context of the ForC.
func (f For) GetCtx() antlr.ParserRuleContext {
	return f.Ctx
}

// Def is a defined variable in the AST.
type Def struct {
	Ctx   antlr.ParserRuleContext
	Id    Id
	Exp   Exp
	Scope ScopeC
}

func (d Def) statNode() {}

// GetCtx returns the context of the DefC.
func (d Def) GetCtx() antlr.ParserRuleContext {
	return d.Ctx
}

// DefLst is a list of definitions in the AST.
type DefLst struct {
	Ctx  antlr.ParserRuleContext
	List []Def
}

func (dl DefLst) statNode() {}

// GetCtx returns the context of the DefLst.
func (dl DefLst) GetCtx() antlr.ParserRuleContext {
	return dl.Ctx
}

// Id is an label in the AST.
type Id struct {
	Ctx    antlr.ParserRuleContext
	Id     string
	TypeId types.Type
}

func (i Id) expNode() {}

// GetCtx returns the context of the IdC.
func (i Id) GetCtx() antlr.ParserRuleContext {
	return i.Ctx
}

// IdLst is a list of labels in the AST.
type IdLst struct {
	Ctx  antlr.ParserRuleContext
	List []Id
}

// GetCtx returns the context of the IdLst.
func (idl IdLst) GetCtx() antlr.ParserRuleContext {
	return idl.Ctx
}

// ExpLst is a list of expressions in the AST.
type ExpLst struct {
	Ctx  antlr.ParserRuleContext
	List []Exp
}

// GetCtx returns the context of the ExpLst.
func (edl ExpLst) GetCtx() antlr.ParserRuleContext {
	return edl.Ctx
}

// BinaryOp is a binary operation node in the AST.
type BinaryOp struct {
	Ctx antlr.ParserRuleContext
	Lhs Exp
	Rhs Exp
	Op  string
}

func (bop BinaryOp) expNode() {}

// GetCtx returns the context of the BinaryOpC.
func (bop BinaryOp) GetCtx() antlr.ParserRuleContext {
	return bop.Ctx
}

// Int is an integer.
type Int struct {
	Ctx antlr.ParserRuleContext
	N   int64
}

func (i Int) expNode() {}

// GetCtx returns the context of the IntC.
func (i Int) GetCtx() antlr.ParserRuleContext {
	return i.Ctx
}

// Float is a float.
type Float struct {
	Ctx antlr.ParserRuleContext
	N   float64
}

func (f Float) expNode() {}

// GetCtx returns the context of the FloatC.
func (f Float) GetCtx() antlr.ParserRuleContext {
	return f.Ctx
}

// String is a String.
type String struct {
	Ctx antlr.ParserRuleContext
	S   string
}

func (s String) expNode() {}

// GetCtx returns the context of the StringC.
func (s String) GetCtx() antlr.ParserRuleContext {
	return s.Ctx
}

// Bool is a boolean.
type Bool struct {
	Ctx  antlr.ParserRuleContext
	True bool
}

func (b Bool) expNode() {}

// GetCtx returns the context of the BoolC.
func (b Bool) GetCtx() antlr.ParserRuleContext {
	return b.Ctx
}

// Nil is a nil node.
type Nil struct {
	Ctx antlr.ParserRuleContext
}

func (n Nil) expNode() {}

// GetCtx returns the context of the NilC.
func (n Nil) GetCtx() antlr.ParserRuleContext {
	return n.Ctx
}
