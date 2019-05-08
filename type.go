package main

type TypeT interface {
	typeDesc()
	Name() string
}

// DefaultT represents the default type
// it does not exist in the concrete grammar
// some nodes will return it during the type check phase when the type returned will not be used
type DefaultT struct {
}

func (dt DefaultT) typeDesc()    {}
func (dt DefaultT) Name() string { return "default" }

type FloatT struct {
}

func (f FloatT) typeDesc()    {}
func (f FloatT) Name() string { return "float" }

type IntT struct {
}

func (i IntT) typeDesc()    {}
func (f IntT) Name() string { return "int" }

type BoolT struct {
}

func (b BoolT) typeDesc()    {}
func (b BoolT) Name() string { return "bool" }

type StringT struct {
}

func (s StringT) typeDesc()    {}
func (f StringT) Name() string { return "string" }

type NilT struct {
}

func (s NilT) typeDesc()    {}
func (f NilT) Name() string { return "nil" }

type ErrorT struct {
}

func (e ErrorT) typeDesc()    {}
func (f ErrorT) Name() string { return "error" }
