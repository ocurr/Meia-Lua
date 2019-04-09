package main

import (
	"errors"
	"reflect"
)

func TypeCheck(root Node) (TypeT, error) {

	switch r := root.(type) {
	case ChunkC:
		return TypeCheck(r.Block)
	case BlockC:
		for i := 0; i < len(r.StatLst); i++ {
		}
	case DefC:
		idT, err := TypeCheck(r.Id)
		if err != nil {
			return nil, err
		}
		expT, err := TypeCheck(r.Exp)
		if err != nil {
			return nil, err
		}

		if idT != expT {
			return nil, errors.New("Type of identifier must match type of expression")
		}

		return idT, nil
	case DefLst:
	case IdC:
		return r.TypeId, nil
	case IdLst:
	case ExpLst:
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
			if !isNumber(lhs) || err != nil {
			}
			rhs, err := TypeCheck(r.Rhs)
			if !isNumber(rhs) || err != nil {
			}
			floatType := reflect.TypeOf(FloatT{})
			if reflect.TypeOf(lhs) == floatType || reflect.TypeOf(rhs) == floatType {
				return FloatT{}, nil
			} else {
				return IntT{}, nil
			}
		case "/":
			lhs, err := TypeCheck(r.Lhs)
			if !isNumber(lhs) || err != nil {
			}
			rhs, err := TypeCheck(r.Rhs)
			if !isNumber(rhs) || err != nil {
			}
			return FloatT{}, nil
		case "//":
			lhs, err := TypeCheck(r.Lhs)
			if !isNumber(lhs) || err != nil {
			}
			rhs, err := TypeCheck(r.Rhs)
			if !isNumber(rhs) || err != nil {
			}
			return IntT{}, nil
		}
	case IntC:
		return IntT{}, nil
	case FloatC:
		return FloatT{}, nil
	case StringC:
		return StringT{}, nil
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
