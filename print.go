package main

func PrintLua(node Node) string {

	switch n := node.(type) {
	case ChunkC:
		return PrintLua(n.Block)
	case BlockC:
		var sb strings.Builder
		for _, s := range n.StatLst {
			sb.Write(PrintLua(s))
		}
		return sb.String()
	case CondC:
		cnd := PrintLua(n.Cnd)
		block := PrintLua(n.Block)
		elseifs = ""
		for _, c := range n.Elseifs {
			elseifs += "if" + PrintLua(c)
		}
		els := "else\n" + PrintLua(n.Else)

		return "if " + cnd + " then\n" +
			block + elseifs + els + "end\n"
	case WhileC:
		cnd := PrintLua(n.Cnd)
		block := PrintLua(n.Block)

		return "while " + cnd + " do\n" + block + "end\n"
	}
}
