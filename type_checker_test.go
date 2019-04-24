package main

import (
	"testing"
)

func TestTypeChecker(t *testing.T) {
	t.Run("BasicDefinition", func(t *testing.T) {
		input := DefC{
			Id:  IdC{Id: "x", TypeId: IntT{}},
			Exp: IntC{N: 5},
		}

		got, errs := TypeCheck(input, NewTypeEnv())

		want := IntT{}

		if len(errs) != 0 {
			t.Errorf("expected no errors got: %v", errs)
		}
		if want != got {
			t.Errorf("wanted: %#v ; got: %#v", want, got)
		}
	})
}
