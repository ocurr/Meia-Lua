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

type ErrorT struct {
}

func (e *ErrorT) typeDesc() {}
