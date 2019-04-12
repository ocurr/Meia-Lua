package main

import (
	"errors"
	"fmt"
	"reflect"
)

func TypeCheck(root Node) (TypeT, []error) {

	ec = new(errorCollector)

	switch r := root.(type) {
	case ChunkC:
		return TypeCheck(r.Block)
	case BlockC:
		for _, s := range r.StatLst {
			_, err := TypeCheck(s)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case DefC:
		idT, err := TypeCheck(r.Id)
		ec.add(err...)

		expT, err := TypeCheck(r.Exp)
		ec.add(err...)

		if idT != expT {
			ec.add(fmt.Errorf("Type of identifier %q does not match type of expression %s", r.Id.Id, expT))
			return ErrorT, ec.errors
		}

		return idT, ec.err
	case DefLst:
		for _, d := range r.List {
			_, err := TypeCheck(d)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case IdC:
		return r.TypeId, ec.errors
	case IdLst:
		for _, i := range r.List {
			_, err := TypeCheck(i)
			ec.add(err...)
		}
		return DefaultT{}, ec.errors
	case ExpLst:
		for _, e := range r.List {
			_, err := TypeCheck(e)
			ec.add(err...)
		}
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
			lhs, err := TypeCheck(r.Lhs)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			floatType := reflect.TypeOf(FloatT{})
			if reflect.TypeOf(lhs) == floatType || reflect.TypeOf(rhs) == floatType {
				return FloatT{}, ec.errors
			} else {
				return IntT{}, ec.errors
			}
		case "/":
			lhs, err := TypeCheck(r.Lhs)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			return FloatT{}, ec.errors
		case "//":
			lhs, err := TypeCheck(r.Lhs)
			ec.add(err...)
			if !isNumber(lhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return ErrorT{}, ec.errors
			}
			rhs, err := TypeCheck(r.Rhs)
			ec.add(err...)
			if !isNumber(rhs) {
				ec.add(fmt.Errorf("operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return ErrorT{}, ec.errors
			}
			return IntT{}, ec.errors
		}
	case IntC:
		return IntT{}, ec.errors
	case FloatC:
		return FloatT{}, ec.errors
	case StringC:
		return StringT{}, ec.errors
	}

	panic("AST Node not implemented")
}

func isNumber(t TypeT) bool {
	switch t.(type) {
	case IntT:
		return true
	case FloatT:
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
