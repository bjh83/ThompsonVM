package oplist

import()

const(
	Start = iota
	Char
	Jump
	Split
	Match
)

type Instruct struct {
	OpCode int
	Line1, Line2 int
}

type OpList struct {
	Head, Tail *OpNode
	Length int
}

type OpNode struct {
	Instruction Instruct
	Next *OpNode
}

func New() *OpList {
	return &OpList{}
}

func (list *OpList) add(opcode int, line1, line2 int) *Instruct {
	node := OpNode{Instruction: Instruct{opcode, line1, line2}}
	if list.Length == 0 {
		list.Head = &node
	} else {
		list.Tail.Next = &node
	}
	list.Tail = &node
	list.Length++
	return &node.Instruction //return a reference so that if we do not know an 
	// address we can figure it out later
}

func (list *OpList) AddSplit(line1, line2 int) *Instruct {
	return list.add(Split, line1, line2)
}

func (list *OpList) AddJump(line int) *Instruct {
	return list.add(Jump, line, -1)
}

func (list *OpList) AddChar(char uint8) {
	//No reason for this to return a reference
	list.add(Char, int(char), -1)
}

func (list *OpList) Start() {
	list.add(Start, -1, -1)
}

func (list *OpList) Finish() {
	list.add(Match, -1, -1)
}

func (list *OpList) Append(toAppend *OpList) {
	toAdd := list.Length
	if toAdd == 0 {
		list.Head = toAppend.Head
		list.Tail = toAppend.Tail
		list.Length = toAppend.Length
		return
	}
	for node := toAppend.Head; node != nil; node = node.Next {
		//Only need to check it here since chars are not stored in Line2
		if node.Instruction.Line1 != -1 && node.Instruction.OpCode != Char {
			node.Instruction.Line1 += toAdd
		}
		if node.Instruction.Line2 != -1 {
			node.Instruction.Line2 += toAdd
		}
	}
	list.Length += toAppend.Length
	list.Tail.Next = toAppend.Head
	list.Tail = toAppend.Tail
}

func (list *OpList) ToArray() []Instruct {
	instructs := make([]Instruct, list.Length)
	index := 0
	for node := list.Head; node != nil; node = node.Next {
		instructs[index] = node.Instruction
		index++
	}
	return instructs
}

