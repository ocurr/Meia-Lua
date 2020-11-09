package typecheck

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/ocurr/Meia-Lua/ast"
	"github.com/ocurr/Meia-Lua/types"
)

// Check type checks root and returns the final type and a list of errors.
func Check(root ast.Node) (types.Type, []error) {
	return CheckWithEnv(root, NewTypeEnv())
}

// CheckWithEnv type checks root using tenv as the types.Type Environment.
func CheckWithEnv(root ast.Node, tenv TypeEnv) (types.Type, []error) {

	ec := new(errorCollector)

	switch r := root.(type) {
	case ast.Chunk:
		return CheckWithEnv(r.Block, tenv)
	case ast.Block:
		for _, s := range r.StatLst {
			_, err := CheckWithEnv(s, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case ast.Def:
		idT, err := CheckWithEnv(r.Id, tenv)
		ec.add(err...)

		expT, err := CheckWithEnv(r.Exp, tenv)
		ec.add(err...)

		if idT == nil {
			ec.add(formatError(r.GetCtx(), "no type defined for %s", r.Id.Id))
			return types.Error{}, ec.errors
		} else if idT != expT {
			ec.add(formatError(r.GetCtx(), "type of identifier %q does not match type of expression %s", r.Id.Id, expT.Name()))
			return types.Error{}, ec.errors
		}

		return idT, ec.errors
	case ast.DefLst:
		for _, d := range r.List {
			_, err := CheckWithEnv(d, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case ast.Id:
		if r.TypeId == nil {
			if t, ok := tenv[r.Id]; ok {
				return t, ec.errors
			}

			ec.add(formatError(r.GetCtx(), "variable %s is used without being declared", r.Id))
			return types.Nil{}, ec.errors
		}
		tenv[r.Id] = r.TypeId
		return r.TypeId, ec.errors
	case ast.IdLst:
		for _, i := range r.List {
			_, err := CheckWithEnv(i, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case ast.ExpLst:
		for _, e := range r.List {
			_, err := CheckWithEnv(e, tenv)
			ec.add(err...)
		}
		return types.Default{}, ec.errors
	case ast.Cond:
		cndT, err := CheckWithEnv(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsBool(cndT) {
			ec.add(formatError(r.GetCtx(), "a conditional statement must evaulate to a boolean value"))
			return types.Error{}, ec.errors
		}

		_, err = CheckWithEnv(r.Block, tenv.Copy())
		ec.add(err...)

		for _, els := range r.Elseifs {
			_, err := CheckWithEnv(els, tenv)
			ec.add(err...)
		}

		_, err = CheckWithEnv(r.Else, tenv)
		ec.add(err...)

		return types.Default{}, ec.errors
	case ast.While:
		cndT, err := CheckWithEnv(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsBool(cndT) {
			ec.add(formatError(r.GetCtx(), "a conditional statement must evaulate to a boolean value"))
			return types.Error{}, ec.errors
		}

		_, err = CheckWithEnv(r.Block, tenv.Copy())
		ec.add(err...)

		return types.Default{}, ec.errors
	case ast.For:
		var ntenv TypeEnv
		if r.Assign.Id.TypeId == nil {
			_, err := CheckWithEnv(r.Assign, tenv)
			ec.add(err...)
		} else {
			ntenv = tenv.Extend(r.Assign.Id.Id, r.Assign.Id.TypeId)
			_, err := CheckWithEnv(r.Assign.Exp, tenv)
			ec.add(err...)
		}

		cndT, err := CheckWithEnv(r.Cnd, tenv)
		ec.add(err...)
		if !types.IsNumber(cndT) {
			ec.add(formatError(r.GetCtx(), "a for statement's limit must be a number"))
			return types.Error{}, ec.errors
		}

		stpT, err := CheckWithEnv(r.Step, tenv)
		ec.add(err...)
		if !types.IsNumber(stpT) {
			ec.add(formatError(r.GetCtx(), "a for statement's step must be a number"))
			return types.Error{}, ec.errors
		}

		_, err = CheckWithEnv(r.Block, ntenv)
		ec.add(err...)

		return types.Default{}, ec.errors
	case ast.BinaryOp:

		switch r.Op {
		case "+":
			fallthrough
		case "*":
			fallthrough
		case "%":
			fallthrough
		case "-":
			lhs, err := CheckWithEnv(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := CheckWithEnv(r.Rhs, tenv)
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
			lhs, err := CheckWithEnv(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := CheckWithEnv(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Float{}, ec.errors
		case "//":
			lhs, err := CheckWithEnv(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := CheckWithEnv(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Int{}, ec.errors
		case "==", "~=":
			lhs, err := CheckWithEnv(r.Lhs, tenv)
			ec.add(err...)
			rhs, err := CheckWithEnv(r.Rhs, tenv)
			ec.add(err...)
			if lhs != rhs {
				ec.add(formatError(r.GetCtx(), "operator: %q requires that both operands have the same type", r.Op))
				return types.Error{}, ec.errors
			}
			return types.Bool{}, ec.errors
		case "<", ">", ">=", "<=":
			lhs, err := CheckWithEnv(r.Lhs, tenv)
			ec.add(err...)
			if !types.IsNumber(lhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, lhs.Name()))
				return types.Error{}, ec.errors
			}
			rhs, err := CheckWithEnv(r.Rhs, tenv)
			ec.add(err...)
			if !types.IsNumber(rhs) {
				ec.add(formatError(r.GetCtx(), "operator: %q expected float or int got: %s", r.Op, rhs.Name()))
				return types.Error{}, ec.errors
			}
			return types.Bool{}, ec.errors
		}
	case ast.Int:
		return types.Int{}, ec.errors
	case ast.Float:
		return types.Float{}, ec.errors
	case ast.String:
		return types.String{}, ec.errors
	case ast.Bool:
		return types.Bool{}, ec.errors
	case ast.Nil:
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
