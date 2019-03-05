package main

type Node struct {
}

type ChunkC struct {
	*Node
	Block BlockC
}

type BlockC struct {
	*Node
	StatLst []StatC
}

type StatC struct {
	*Node
	DefLst []DefC
}

type DefC struct {
	*StatC
	Id     string
	TypeId TypeT
}

type TypeC struct {
	Type TypeT
}
