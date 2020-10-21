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
	ctx   antlr.ParserRuleContext
	Block BlockC
}

// GetCtx returns the context of the ChunkC.
func (c ChunkC) GetCtx() antlr.ParserRuleContext {
	return c.ctx
}

// BlockC is a Block node in the AST.
type BlockC struct {
	ctx     antlr.ParserRuleContext
	StatLst []Stat
}

// GetCtx returns the context of the BlockC.
func (b BlockC) GetCtx() antlr.ParserRuleContext {
	return b.ctx
}

// CondC is a conditional statement node in the AST.
type CondC struct {
	ctx     antlr.ParserRuleContext
	Cnd     Exp
	Block   BlockC
	Elseifs []CondC
	Else    BlockC
}

func (c CondC) statNode() {}

// GetCtx returns the context of the CondC.
func (c CondC) GetCtx() antlr.ParserRuleContext {
	return c.ctx
}

// WhileC is a while loop node in the AST.
type WhileC struct {
	ctx   antlr.ParserRuleContext
	Cnd   Exp
	Block BlockC
}

func (w WhileC) statNode() {}

// GetCtx returns the context of the WhileC.
func (w WhileC) GetCtx() antlr.ParserRuleContext {
	return w.ctx
}

// ForC is a for loop node in the AST.
type ForC struct {
	ctx    antlr.ParserRuleContext
	Assign DefC
	Cnd    Exp
	Step   Exp
	Block  BlockC
}

func (f ForC) statNode() {}

// GetCtx returns the context of the ForC.
func (f ForC) GetCtx() antlr.ParserRuleContext {
	return f.ctx
}

// DefC is a defined variable in the AST.
type DefC struct {
	ctx   antlr.ParserRuleContext
	Id    IdC
	Exp   Exp
	Scope ScopeC
}

func (d DefC) statNode() {}

// GetCtx returns the context of the DefC.
func (d DefC) GetCtx() antlr.ParserRuleContext {
	return d.ctx
}

// DefLst is a list of definitions in the AST.
type DefLst struct {
	ctx  antlr.ParserRuleContext
	List []DefC
}

func (dl DefLst) statNode() {}

// GetCtx returns the context of the DefLst.
func (dl DefLst) GetCtx() antlr.ParserRuleContext {
	return dl.ctx
}

// IdC is an label in the AST.
type IdC struct {
	ctx    antlr.ParserRuleContext
	Id     string
	TypeId TypeT
}

func (i IdC) expNode() {}

// GetCtx returns the context of the IdC.
func (i IdC) GetCtx() antlr.ParserRuleContext {
	return i.ctx
}

// IdLst is a list of labels in the AST.
type IdLst struct {
	ctx  antlr.ParserRuleContext
	List []IdC
}

// GetCtx returns the context of the IdLst.
func (idl IdLst) GetCtx() antlr.ParserRuleContext {
	return idl.ctx
}

// ExpLst is a list of expressions in the AST.
type ExpLst struct {
	ctx  antlr.ParserRuleContext
	List []Exp
}

// GetCtx returns the context of the ExpLst.
func (edl ExpLst) GetCtx() antlr.ParserRuleContext {
	return edl.ctx
}

// BinaryOpC is a binary operation node in the AST.
type BinaryOpC struct {
	ctx antlr.ParserRuleContext
	Lhs Exp
	Rhs Exp
	Op  string
}

func (bop BinaryOpC) expNode() {}

// GetCtx returns the context of the BinaryOpC.
func (bop BinaryOpC) GetCtx() antlr.ParserRuleContext {
	return bop.ctx
}

// IntC is an integer.
type IntC struct {
	ctx antlr.ParserRuleContext
	N   int64
}

func (i IntC) expNode() {}

// GetCtx returns the context of the IntC.
func (i IntC) GetCtx() antlr.ParserRuleContext {
	return i.ctx
}

// FloatC is a float.
type FloatC struct {
	ctx antlr.ParserRuleContext
	N   float64
}

func (f FloatC) expNode() {}

// GetCtx returns the context of the FloatC.
func (f FloatC) GetCtx() antlr.ParserRuleContext {
	return f.ctx
}

// StringC is a String.
type StringC struct {
	ctx antlr.ParserRuleContext
	S   string
}

func (s StringC) expNode() {}

// GetCtx returns the context of the StringC.
func (s StringC) GetCtx() antlr.ParserRuleContext {
	return s.ctx
}

// BoolC is a boolean.
type BoolC struct {
	ctx  antlr.ParserRuleContext
	True bool
}

func (b BoolC) expNode() {}

// GetCtx returns the context of the BoolC.
func (b BoolC) GetCtx() antlr.ParserRuleContext {
	return b.ctx
}

// NilC is a nil node.
type NilC struct {
	ctx antlr.ParserRuleContext
}

func (n NilC) expNode() {}

// GetCtx returns the context of the NilC.
func (n NilC) GetCtx() antlr.ParserRuleContext {
	return n.ctx
}
