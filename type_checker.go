package main

import (
	"errors"
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
	case IntC:
		return IntT{}, nil
	case FloatC:
		return FloatT{}, nil
	case StringC:
		return StringT{}, nil
	}

	panic("type not implemented")
}
