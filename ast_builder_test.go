package main

import (
	"reflect"
	"testing"
)

func TestASTBuilder(t *testing.T) {
}

func TestASTMatch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		res := astMatch(IdC{Id: "x", TypeId: IntT{}}, IdC{Id: "x", TypeId: IntT{}})
		if !res {
			t.Errorf("IdC{Id: \"x\", TypeId: IntT} != IdC{Id: \"x\", TypeId: IntT}")
		}
	})
	t.Run("2", func(t *testing.T) {
		res := astMatch(
			IdC{Id: "y", TypeId: IntT{}},
			IdC{Id: "x", TypeId: IntT{}})
		if res {
			t.Errorf("IdC{Id: \"y\", TypeId: IntT} == IdC{Id: \"x\", TypeId: IntT}")
		}
	})
}

func astMatch(node1, node2 Node) bool {
	if reflect.TypeOf(node1) != reflect.TypeOf(node2) {
		return false
	}

	switch n1 := node1.(type) {
	case ChunkC:
		n2 := node2.(ChunkC)
		return astMatch(n1.Block, n2.Block)
	case BlockC:
		n2 := node2.(BlockC)
		if len(n1.StatLst) != len(n2.StatLst) {
			return false
		}

		for i := 0; i < len(n1.StatLst); i++ {
			if !astMatch(n1.StatLst[i], n2.StatLst[i]) {
				return false
			}
		}

		return true
	case DefC:
		n2 := node2.(DefC)
		return astMatch(n1.Id, n2.Id) && astMatch(n1.Exp, n2.Exp)
	case IdC:
		n2 := node2.(IdC)
		return n1.Id == n2.Id && n1.TypeId == n2.TypeId
	case ExpLst:
		n2 := node2.(ExpLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case DefLst:
		n2 := node2.(DefLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case IdLst:
		n2 := node2.(IdLst)
		if len(n1.List) != len(n2.List) {
			return false
		}

		for i := 0; i < len(n1.List); i++ {
			if !astMatch(n1.List[i], n2.List[i]) {
				return false
			}
		}

		return true
	case IntC:
		n2 := node2.(IntC)
		return n1.N == n2.N
	case FloatC:
		n2 := node2.(FloatC)
		return n1.N == n2.N
	case StringC:
		n2 := node2.(StringC)
		return n1.S == n2.S
	default:
		panic("Type not implemented")
	}

	return false
}
