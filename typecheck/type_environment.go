package typecheck

import "github.com/ocurr/Meia-Lua/types"

// TypeEnv is a type environment representation that maps from labels to types.
type TypeEnv map[string]types.Type

// NewTypeEnv returns a new empty TypeEnv.
func NewTypeEnv() TypeEnv {
	return map[string]types.Type{}
}

// Copy returns a copy of the TypeEnv.
func (tenv TypeEnv) Copy() TypeEnv {
	newTenv := NewTypeEnv()
	for k, v := range tenv {
		newTenv[k] = v
	}
	return newTenv
}

// Extend adds a new label -> type mapping to the TypeEnv.
func (tenv TypeEnv) Extend(id string, t types.Type) TypeEnv {
	ntenv := tenv.Copy()
	ntenv[id] = t

	return ntenv
}
