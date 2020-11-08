package types

// Type represents the type of a given node of the AST.
type Type interface {
	typeDesc()
	Name() string
}

// Default represents the default type it does not exist in the concrete grammar.
// Some nodes will return it during the type check phase when the type returned will not be used.
type Default struct{}

func (dt Default) typeDesc() {}

// Name returns the string representation of the type.
func (dt Default) Name() string { return "default" }

// Float represents a node with a float type.
type Float struct{}

func (f Float) typeDesc() {}

// Name returns the string representation of the type.
func (f Float) Name() string { return "float" }

// Int represents a node with the integer type.
type Int struct{}

func (i Int) typeDesc() {}

// Name returns the string representation of the type.
func (i Int) Name() string { return "int" }

// Bool represents a node with a boolean type.
type Bool struct{}

func (b Bool) typeDesc() {}

// Name returns the string representation of the type.
func (b Bool) Name() string { return "bool" }

// String represents a node with a string type.
type String struct{}

func (s String) typeDesc() {}

// Name returns the string representation of the type.
func (s String) Name() string { return "string" }

// Nil represents a nil type.
type Nil struct{}

func (n Nil) typeDesc() {}

// Name returns the string representation of the type.
func (n Nil) Name() string { return "nil" }

// Error represents an error type. It does not exist in the concrete grammar.
// It will be returned by nodes when a type error occurs.
type Error struct{}

func (e Error) typeDesc() {}

// Name returns the string representation of the type.
func (e Error) Name() string { return "error" }
