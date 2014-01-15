package vm

import(
	. "../oplist"
	"./queue"
)

func ThompsonVM(input string, instructions []Instruct) bool {
	cqueue := queue.New()
	nqueue := queue.New()
	cqueue.Push(0)
	for strCount := 0; strCount <= len(input); strCount++ {
		for cqueue.Len() > 0 {
			instrCount := cqueue.Pop()
			instr := instructions[instrCount]
			switch instr.OpCode {
			case Start:
				cqueue.Push(1) //This will cause a bootch is Start is not the first instr
				if instrCount != 0 {
					return false
				}
				break
			case Char:
				if strCount >= len(input) || input[strCount] != uint8(instr.Line1) {
					break //Thread dies
				}
				nqueue.Push(instrCount + 1)
				break
			case Jump:
				cqueue.Push(instr.Line1)
				break
			case Split:
				cqueue.Push(instr.Line1)
				cqueue.Push(instr.Line2)
				break
			case Match:
				if strCount == len(input) {
					return true
				}
				break
			}
		}
		cqueue, nqueue = nqueue, cqueue
	}
	return false
}

