package main

type TypeT interface {
	typeDesc()
}

type FloatT struct {
}

func (f FloatT) typeDesc() {}

type IntT struct {
}

func (i IntT) typeDesc() {}

type StringT struct {
}

func (s StringT) typeDesc() {}

type ErrorT struct {
}

func (e *ErrorT) typeDesc() {}
