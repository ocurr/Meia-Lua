package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
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

// ChunkC is a Chunk node in the AST.
type ChunkC struct {
	Ctx   antlr.ParserRuleContext
	Block BlockC
}

// GetCtx returns the context of the ChunkC.
func (c ChunkC) GetCtx() antlr.ParserRuleContext {
	return c.Ctx
}

// BlockC is a Block node in the AST.
type BlockC struct {
	Ctx     antlr.ParserRuleContext
	StatLst []Stat
}

// GetCtx returns the context of the BlockC.
func (b BlockC) GetCtx() antlr.ParserRuleContext {
	return b.Ctx
}

// CondC is a conditional statement node in the AST.
type CondC struct {
	Ctx     antlr.ParserRuleContext
	Cnd     Exp
	Block   BlockC
	Elseifs []CondC
	Else    BlockC
}

func (c CondC) statNode() {}

// GetCtx returns the context of the CondC.
func (c CondC) GetCtx() antlr.ParserRuleContext {
	return c.Ctx
}

// WhileC is a while loop node in the AST.
type WhileC struct {
	Ctx   antlr.ParserRuleContext
	Cnd   Exp
	Block BlockC
}

func (w WhileC) statNode() {}

// GetCtx returns the context of the WhileC.
func (w WhileC) GetCtx() antlr.ParserRuleContext {
	return w.Ctx
}

// ForC is a for loop node in the AST.
type ForC struct {
	Ctx    antlr.ParserRuleContext
	Assign DefC
	Cnd    Exp
	Step   Exp
	Block  BlockC
}

func (f ForC) statNode() {}

// GetCtx returns the context of the ForC.
func (f ForC) GetCtx() antlr.ParserRuleContext {
	return f.Ctx
}

// DefC is a defined variable in the AST.
type DefC struct {
	Ctx   antlr.ParserRuleContext
	Id    IdC
	Exp   Exp
	Scope ScopeC
}

func (d DefC) statNode() {}

// GetCtx returns the context of the DefC.
func (d DefC) GetCtx() antlr.ParserRuleContext {
	return d.Ctx
}

// DefLst is a list of definitions in the AST.
type DefLst struct {
	Ctx  antlr.ParserRuleContext
	List []DefC
}

func (dl DefLst) statNode() {}

// GetCtx returns the context of the DefLst.
func (dl DefLst) GetCtx() antlr.ParserRuleContext {
	return dl.Ctx
}

// IdC is an label in the AST.
type IdC struct {
	Ctx    antlr.ParserRuleContext
	Id     string
	TypeId TypeT
}

func (i IdC) expNode() {}

// GetCtx returns the context of the IdC.
func (i IdC) GetCtx() antlr.ParserRuleContext {
	return i.Ctx
}

// IdLst is a list of labels in the AST.
type IdLst struct {
	Ctx  antlr.ParserRuleContext
	List []IdC
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

// BinaryOpC is a binary operation node in the AST.
type BinaryOpC struct {
	Ctx antlr.ParserRuleContext
	Lhs Exp
	Rhs Exp
	Op  string
}

func (bop BinaryOpC) expNode() {}

// GetCtx returns the context of the BinaryOpC.
func (bop BinaryOpC) GetCtx() antlr.ParserRuleContext {
	return bop.Ctx
}

// IntC is an integer.
type IntC struct {
	Ctx antlr.ParserRuleContext
	N   int64
}

func (i IntC) expNode() {}

// GetCtx returns the context of the IntC.
func (i IntC) GetCtx() antlr.ParserRuleContext {
	return i.Ctx
}

// FloatC is a float.
type FloatC struct {
	Ctx antlr.ParserRuleContext
	N   float64
}

func (f FloatC) expNode() {}

// GetCtx returns the context of the FloatC.
func (f FloatC) GetCtx() antlr.ParserRuleContext {
	return f.Ctx
}

// StringC is a String.
type StringC struct {
	Ctx antlr.ParserRuleContext
	S   string
}

func (s StringC) expNode() {}

// GetCtx returns the context of the StringC.
func (s StringC) GetCtx() antlr.ParserRuleContext {
	return s.Ctx
}

// BoolC is a boolean.
type BoolC struct {
	Ctx  antlr.ParserRuleContext
	True bool
}

func (b BoolC) expNode() {}

// GetCtx returns the context of the BoolC.
func (b BoolC) GetCtx() antlr.ParserRuleContext {
	return b.Ctx
}

// NilC is a nil node.
type NilC struct {
	Ctx antlr.ParserRuleContext
}

func (n NilC) expNode() {}

// GetCtx returns the context of the NilC.
func (n NilC) GetCtx() antlr.ParserRuleContext {
	return n.Ctx
}
