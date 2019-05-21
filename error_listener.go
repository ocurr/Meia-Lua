package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type LuaErrorListener struct {
	*antlr.DefaultErrorListener
	del *antlr.DiagnosticErrorListener
}

func NewLuaErrorListener() *LuaErrorListener {
	el := new(LuaErrorListener)
	el.del = antlr.NewDiagnosticErrorListener(true)
	return el
}

func (d *LuaErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.del.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// NO-OP
func (d *LuaErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	return
}

func (d *LuaErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.del.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}
