package main

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/Meia-Lua/types"
)

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

// LuaTypeCheck type checks root and returns the final type and a list of errors.
func LuaTypeCheck(root Node) (types.Type, []error) {
	return TypeCheck(root, NewTypeEnv())
}

// TypeCheck type checks root using tenv as the types.Type Environment.
func TypeCheck(root Node, tenv TypeEnv) (types.Type, []error) {

	ec := new(errorCollector)

	switch r := root.(type) {
	case ChunkC:
		return TypeCheck(r.Block, tenv)
	case BlockC:
		for _, s := range r.StatLst {
			_, err := TypeCheck(s, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case DefC:
		idT, err := TypeCheck(r.Id, tenv)
		ec.add(err...)

		expT, err := TypeCheck(r.Exp, tenv)
		ec.add(err...)

		if idT == nil {
			ec.add(formatError(r.GetCtx(), "no type defined for %s", r.Id.Id))
			return types.Error{}, ec.errors
		} else if idT != expT {
			ec.add(formatError(r.GetCtx(), "type of identifier %q does not match type of expression %s", r.Id.Id, expT.Name()))
			return types.Error{}, ec.errors
		}

		return idT, ec.errors
	case DefLst:
		for _, d := range r.List {
			_, err := TypeCheck(d, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case IdC:
		if r.TypeId == nil {
			if t, ok := tenv[r.Id]; ok {
				return t, ec.errors
			}

			ec.add(formatError(r.GetCtx(), "variable %s does is used without being declared", r.Id))
			return types.Nil{}, ec.errors
		}
		tenv[r.Id] = r.TypeId
		return r.TypeId, ec.errors
	case IdLst:
		for _, i := range r.List {
			_, err := TypeCheck(i, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case ExpLst:
		for _, e := range r.List {
			_, err := TypeCheck(e, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case CondC:
		cndT, err := TypeCheck(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsBool(cndT) {
			ec.add(formatError(r.GetCtx(), "a conditional statement must evaulate to a boolean value"))
			return types.Error{}, ec.errors
		}

		_, err = TypeCheck(r.Block, tenv.Copy())
		ec.add(err...)

		for _, els := range r.Elseifs {
			_, err := TypeCheck(els, tenv)
			ec.add(err...)
		}

		_, err = TypeCheck(r.Else, tenv)
		ec.add(err...)

		return types.Default{}, ec.errors
	case WhileC:
		cndT, err := TypeCheck(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsBool(cndT) {
			ec.add(formatError(r.GetCtx(), "a conditional statement must evaulate to a boolean value"))
			return types.Error{}, ec.errors
		}

		_, err = TypeCheck(r.Block, tenv.Copy())
		ec.add(err...)

		return types.Default{}, ec.errors
	case ForC:
		var ntenv TypeEnv
		if r.Assign.Id.TypeId == nil {
			_, err := TypeCheck(r.Assign, tenv)
			ec.add(err...)
		} else {
			ntenv = tenv.Extend(r.Assign.Id.Id, r.Assign.Id.TypeId)
			_, err := TypeCheck(r.Assign.Exp, tenv)
			ec.add(err...)
		}

		cndT, err := TypeCheck(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsNumber(cndT) {
			ec.add(formatError(r.GetCtx(), "a for statement's limit must be a number"))
			return types.Error{}, ec.errors
		}

		stpT, err := TypeCheck(r.Step, tenv)
		ec.add(err...)
		if !types.IsNumber(stpT) {
			ec.add(formatError(r.GetCtx(), "a for statement's step must be a number"))
			return types.Error{}, ec.errors
		}

		_, err = TypeCheck(r.Block, ntenv)
		ec.add(err...)

		return types.Default{}, ec.errors
	case BinaryOpC:

		switch r.Op {
		case "+":
			fallthrough
		case "*":
			fallthrough
		case "%":
			fallthrough
		case "-":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			if types.IsFloat(lhs) || types.IsFloat(rhs) {
				return types.Float{}, ec.errors
			}
			return types.Int{}, ec.errors
		case "/":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Float{}, ec.errors
		case "//":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Int{}, ec.errors
		case "==", "~=":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if lhs != rhs {
				ec.add(formatError(r.GetCtx(), "operator: %q requires that both operands have the same type", r.Op))
				return types.Error{}, ec.errors
			}
			return types.Bool{}, ec.errors
		case "<", ">", ">=", "<=":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Bool{}, ec.errors
		}
	case IntC:
		return types.Int{}, ec.errors
	case FloatC:
		return types.Float{}, ec.errors
	case StringC:
		return types.String{}, ec.errors
	case BoolC:
		return types.Bool{}, ec.errors
	case NilC:
		return types.Nil{}, ec.errors
	}

	panic(fmt.Sprintf("AST Node not implemented or you forgot a return %#v", root))
}

func formatError(ctx antlr.ParserRuleContext, errf string, vals ...interface{}) error {

	start := ctx.GetStart()
	return fmt.Errorf("%s:%d:%d\n\t%s", start.GetInputStream().GetSourceName(), start.GetLine(), start.GetColumn(), fmt.Errorf(errf, vals...))
}

type errorCollector struct {
	errors []error
}

func (ec *errorCollector) add(errs ...error) {
	for _, e := range errs {
		if e != nil {
			ec.errors = append(ec.errors, e)
		}
	}
}
