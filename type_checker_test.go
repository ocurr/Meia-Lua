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

		got, errs := LuaTypeCheck(input)

		want := IntT{}

		if len(errs) != 0 {
			t.Errorf("expected no errors got: %v", errs)
		}
		if want != got {
			t.Errorf("wanted: %#v ; got: %#v", want, got)
		}
	})

	t.Run("If", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}}},
			}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})

	t.Run("IfElseIf", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}},
					Elseifs: []CondC{
						CondC{
							Cnd: BoolC{True: false},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 5}}}}}}},
					},
				}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("IfElseIfElse", func(t *testing.T) {
		input := CondC{
			Cnd: BoolC{True: true},
			Block: BlockC{StatLst: []Stat{
				DefLst{List: []DefC{
					DefC{Id: IdC{Id: "x", TypeId: IntT{}},
						Exp: IntC{N: 5}}}},
				CondC{
					Cnd: BoolC{True: true},
					Block: BlockC{
						StatLst: []Stat{
							DefLst{List: []DefC{
								DefC{Id: IdC{Id: "x", TypeId: IntT{}},
									Exp: IntC{N: 6}}}}}},
					Elseifs: []CondC{
						CondC{
							Cnd: BoolC{True: false},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 5}}}}}}},
					},
					Else: BlockC{StatLst: []Stat{
						CondC{
							Cnd: BoolC{True: true},
							Block: BlockC{
								StatLst: []Stat{
									DefLst{List: []DefC{
										DefC{Id: IdC{Id: "x", TypeId: IntT{}},
											Exp: IntC{N: 6}}}}}},
						}},
					}}}}}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
	t.Run("While", func(t *testing.T) {
		input := WhileC{
			Cnd: BoolC{True: true},
			Block: BlockC{
				StatLst: []Stat{
					DefLst{List: []DefC{
						DefC{Id: IdC{Id: "x", TypeId: IntT{}},
							Exp: IntC{N: 6}}}}}},
		}

		_, err := LuaTypeCheck(input)

		if len(err) != 0 {
			t.Errorf("expected no errors got: %v", err)
		}
	})
}
