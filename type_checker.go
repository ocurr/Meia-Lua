package main

import (
	"fmt"
	"reflect"
)

type TypeEnv map[string]TypeT

func NewTypeEnv() TypeEnv {
	return map[string]TypeT{}
}

func (tenv TypeEnv) Copy() TypeEnv {
	newTenv := NewTypeEnv()
	for k, v := range tenv {
		newTenv[k] = v
	}
	return newTenv
}

func (tenv TypeEnv) Extend(id string, t TypeT) TypeEnv {
	ntenv := tenv.Copy()
	ntenv[id] = t

	return ntenv
}

func LuaTypeCheck(root Node) (TypeT, []error) {
	return TypeCheck(root, NewTypeEnv())
}

func TypeCheck(root Node, tenv TypeEnv) (TypeT, []error) {

	ec := new(errorCollector)

	switch r := root.(type) {
	case ChunkC:
		return TypeCheck(r.Block, tenv)
	case BlockC:
		for _, s := range r.StatLst {
			_, err := TypeCheck(s, tenv)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case DefC:
		idT, err := TypeCheck(r.Id, tenv)
		ec.add(err...)

		expT, err := TypeCheck(r.Exp, tenv)
		ec.add(err...)

		if idT == nil {
			ec.add(fmt.Errorf("No type defined for %s.", r.Id.Id))
			return ErrorT{}, ec.errors
		} else if idT != expT {
			ec.add(fmt.Errorf("Type of identifier %q does not match type of expression %s.", r.Id.Id, expT.Name()))
			return ErrorT{}, ec.errors
		}

		return idT, ec.errors
	case DefLst:
		for _, d := range r.List {
			_, err := TypeCheck(d, tenv)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case IdC:
		if r.TypeId == nil {
			t := tenv[r.Id]
			return t, ec.errors
		}
		tenv[r.Id] = r.TypeId
		return r.TypeId, ec.errors
	case IdLst:
		for _, i := range r.List {
			_, err := TypeCheck(i, tenv)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case ExpLst:
		for _, e := range r.List {
			_, err := TypeCheck(e, tenv)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case CondC:
		cndT, err := TypeCheck(r.Cnd, tenv)
		ec.add(err...)
		if !isBool(cndT) {
			ec.add(fmt.Errorf("a conditional statement must evaulate to a boolean value."))
			return ErrorT{}, ec.errors
		}

		_, err = TypeCheck(r.Block, tenv.Copy())
		ec.add(err...)

		for _, els := range r.Elseifs {
			_, err := TypeCheck(els, tenv)
			ec.add(err...)
		}

		_, err = TypeCheck(r.Else, tenv)
		ec.add(err...)

		return DefaultT{}, ec.errors
	case WhileC:
		cndT, err := TypeCheck(r.Cnd, tenv)
		ec.add(err...)
		if !isBool(cndT) {
			ec.add(fmt.Errorf("a conditional statement must evaulate to a boolean value."))
			return ErrorT{}, ec.errors
		}

		_, err = TypeCheck(r.Block, tenv.Copy())
		ec.add(err...)

		return DefaultT{}, ec.errors
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
		if !isNumber(cndT) {
			ec.add(fmt.Errorf("A For statement's limit must be a number."))
			return ErrorT{}, ec.errors
		}

		stpT, err := TypeCheck(r.Step, tenv)
		ec.add(err...)
		if !isNumber(stpT) {
			ec.add(fmt.Errorf("A For statement's step must be a number."))
			return ErrorT{}, ec.errors
		}

		_, err = TypeCheck(r.Block, ntenv)
		ec.add(err...)

		return DefaultT{}, ec.errors
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
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			floatType := reflect.TypeOf(FloatT{})
			if reflect.TypeOf(lhs) == floatType || reflect.TypeOf(rhs) == floatType {
				return FloatT{}, ec.errors
			} else {
				return IntT{}, ec.errors
			}
		case "/":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			return FloatT{}, ec.errors
		case "//":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			return IntT{}, ec.errors
		case "==", "~=":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			rhs, err := TypeCheck(r.Lhs, tenv)
			if lhs != rhs {
				ec.add(fmt.Errorf("operator: %q requires that both operands have the same type.", r.Op))
				return ErrorT{}, ec.errors
			}
			return BoolT{}, ec.errors
		case "<", ">", ">=", "<=":
			lhs, err := TypeCheck(r.Lhs, tenv)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs, tenv)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s.", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			return BoolT{}, ec.errors
		}
	case IntC:
		return IntT{}, ec.errors
	case FloatC:
		return FloatT{}, ec.errors
	case StringC:
		return StringT{}, ec.errors
	case BoolC:
		return BoolT{}, ec.errors
	case NilC:
		return NilT{}, ec.errors
	}

	panic(fmt.Sprintf("AST Node not implemented or you forgot a return %#v", root))
}

func isNumber(t TypeT) bool {
	switch t.(type) {
	case IntT, FloatT:
		return true
	default:
		return false
	}
}

func isBool(t TypeT) bool {
	switch t.(type) {
	case BoolT:
		return true
	default:
		return false
	}
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
