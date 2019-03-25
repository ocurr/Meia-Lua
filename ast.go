package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type Node interface {
	GetCtx() *antlr.BaseParserRuleContext
}

type Exp interface {
	Node
	expNode()
}

type Stat interface {
	Node
	statNode()
}

type ChunkC struct {
	Ctx   *antlr.BaseParserRuleContext
	Block BlockC
}

func (c ChunkC) GetCtx() *antlr.BaseParserRuleContext {
	return c.Ctx
}

type BlockC struct {
	Ctx     *antlr.BaseParserRuleContext
	StatLst []Stat
}

func (b BlockC) GetCtx() *antlr.BaseParserRuleContext {
	return b.Ctx
}

type DefC struct {
	Ctx *antlr.BaseParserRuleContext
	Id  IdC
	Exp Exp
}

func (d DefC) GetCtx() *antlr.BaseParserRuleContext {
	return d.Ctx
}

func (d DefC) statNode() {}

type DefLst struct {
	List []DefC
}

func (dl DefLst) GetCtx() *antlr.BaseParserRuleContext {
	return antlr.NewBaseParserRuleContext(nil, -1)
}

func (dl DefLst) statNode() {}

type IdC struct {
	Ctx    *antlr.BaseParserRuleContext
	Id     string
	TypeId TypeT
}

func (i IdC) GetCtx() *antlr.BaseParserRuleContext {
	return i.Ctx
}

type IdLst struct {
	List []IdC
}

func (idl IdLst) GetCtx() *antlr.BaseParserRuleContext {
	return antlr.NewBaseParserRuleContext(nil, -1)
}

type ExpLst struct {
	Ctx  *antlr.BaseParserRuleContext
	List []Exp
}

func (edl ExpLst) GetCtx() *antlr.BaseParserRuleContext {
	return antlr.NewBaseParserRuleContext(nil, -1)
}

type IntC struct {
	Ctx *antlr.BaseParserRuleContext
	N   int64
}

func (i IntC) GetCtx() *antlr.BaseParserRuleContext {
	return i.Ctx
}

func (i IntC) expNode() {}

type FloatC struct {
	Ctx *antlr.BaseParserRuleContext
	N   float64
}

func (f FloatC) GetCtx() *antlr.BaseParserRuleContext {
	return f.Ctx
}

func (f FloatC) expNode() {}

type StringC struct {
	Ctx *antlr.BaseParserRuleContext
	S   string
}

func (s StringC) GetCtx() *antlr.BaseParserRuleContext {
	return s.Ctx
}

func (s StringC) expNode() {}
