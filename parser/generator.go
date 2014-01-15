package parser

import(
	"../oplist"
	. "../lexer"
)

func (start *Start) Generate() *oplist.OpList {
	oplist := oplist.New()
	oplist.Start()
	oplist.Append(start.generate())
	oplist.Finish()
	return oplist
}

func (start *Start) generate() *oplist.OpList {
	oplist := oplist.New()
	if !start.Right.Empty {
		//this means there should be a split here
		save1 := oplist.AddSplit(1, -1) //we do not know the second address yet
		oplist.Append(start.Left.generate())
		save2 := oplist.AddJump(-1)
		save1.Line2 = oplist.Length
		oplist.Append(start.Right.generate())
		save2.Line1 = oplist.Length
	} else {
		oplist.Append(start.Left.generate())
	}
	return oplist
}

func (start *Start_) generate() *oplist.OpList {
	oplist := oplist.New()
	if !start.Right.Empty {
		save1 := oplist.AddSplit(1, -1) //we do not know the second address yet
		oplist.Append(start.Left.generate())
		save2 := oplist.AddJump(-1)
		save1.Line2 = oplist.Length
		oplist.Append(start.Right.generate())
		save2.Line1 = oplist.Length
	} else {
		oplist.Append(start.Left.generate())
	}
	return oplist
}

func (juxt *Juxt) generate() *oplist.OpList {
	oplist := juxt.Left.generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.generate())
	}
	return oplist
}

func (juxt *Juxt_) generate() *oplist.OpList {
	oplist := juxt.Left.generate()
	if !juxt.Right.Empty {
		oplist.Append(juxt.Right.generate())
	}
	return oplist
}

func (quant *Quant) generate() *oplist.OpList {
	oplist := oplist.New()
	switch quant.Type {
	case Star:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.generate())
		oplist.AddJump(0)
		save.Line2 = oplist.Length
		break
	case Plus:
		oplist.Append(quant.Left.generate())
		oplist.AddSplit(0, oplist.Length + 1) //Either do again or end
		break
	case Ques:
		save := oplist.AddSplit(1, -1)
		oplist.Append(quant.Left.generate())
		save.Line2 = oplist.Length
		break
	case Epsilon:
		oplist = quant.Left.generate()
		break
	}
	return oplist
}

func (ident *Ident) generate() *oplist.OpList {
	if ident.Left != nil {
		return ident.Left.generate()
	}
	oplist := oplist.New()
	oplist.AddChar(ident.Char)
	return oplist
}

