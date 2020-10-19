package main

// TypeT represents the type of a given node of the AST.
type TypeT interface {
	typeDesc()
	Name() string
}

// DefaultT represents the default type it does not exist in the concrete grammar.
// Some nodes will return it during the type check phase when the type returned will not be used.
type DefaultT struct{}

func (dt DefaultT) typeDesc() {}

// Name returns the string representation of the type.
func (dt DefaultT) Name() string { return "default" }

// FloatT represents a node with a float type.
type FloatT struct{}

func (f FloatT) typeDesc() {}

// Name returns the string representation of the type.
func (f FloatT) Name() string { return "float" }

// IntT represents a node with the integer type.
type IntT struct{}

func (i IntT) typeDesc() {}

// Name returns the string representation of the type.
func (i IntT) Name() string { return "int" }

// BoolT represents a node with a boolean type.
type BoolT struct{}

func (b BoolT) typeDesc() {}

// Name returns the string representation of the type.
func (b BoolT) Name() string { return "bool" }

// StringT represents a node with a string type.
type StringT struct{}

func (s StringT) typeDesc() {}

// Name returns the string representation of the type.
func (s StringT) Name() string { return "string" }

// NilT represents a nil type.
type NilT struct{}

func (n NilT) typeDesc() {}

// Name returns the string representation of the type.
func (n NilT) Name() string { return "nil" }

// ErrorT represents an error type. It does not exist in the concrete grammar.
// It will be returned by nodes when a type error occurs.
type ErrorT struct{}

func (e ErrorT) typeDesc() {}

// Name returns the string representation of the type.
func (e ErrorT) Name() string { return "error" }
