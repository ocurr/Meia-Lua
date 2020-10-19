package main

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// LuaErrorListener represents an Error Listener for the Meia-Lua language.
type LuaErrorListener struct {
	*antlr.DefaultErrorListener
	del *antlr.DiagnosticErrorListener
}

// NewLuaErrorListener returns a new LuaErrorListener.
func NewLuaErrorListener() *LuaErrorListener {
	el := new(LuaErrorListener)
	el.del = antlr.NewDiagnosticErrorListener(true)
	return el
}

// ReportAmbiguity is called to report ambiguities in the token stream.
func (d *LuaErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	d.del.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

// ReportAttemptingFullContext is a NO-OP.
func (d *LuaErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
}

// ReportContextSensitivity is called to report context sensitive tokens in the token stream.
func (d *LuaErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	d.del.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}
