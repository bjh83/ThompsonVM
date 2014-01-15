package lexer

import("fmt")

const(Buffer = 255)

const(
	//XXX: ANYTHING with precedence higher than the left associative 
	//operators must be listed before RParen
	LParen = iota + Buffer
	//XXX:From this point on the parser will ignore if not explicitly handled!!!
	RParen = iota + Buffer
	Pipe = iota + Buffer
	Star = iota + Buffer
	Plus = iota + Buffer
	Ques = iota + Buffer
	Epsilon = iota + Buffer
)

func Lex(regex string) []int {
	nextEscape := false
	escaped := false
	out := make([]int, len(regex))
	out_index := 0
	for i := 0; i < len(regex); i++ {
		switch regex[i] {
		case '\\':
			if escaped {
				out[out_index] = '|'
				out_index++
				escaped = false
			} else {
				nextEscape = true
			}
			break
		case '|':
			if escaped {
				out[out_index] = '|'
				out_index++
				escaped = false
			} else {
				out[out_index] = Pipe
				out_index++
			}
			break
		case '*':
			if escaped {
				out[out_index] = '*'
				out_index++
				escaped = false
			} else {
				out[out_index] = Star
				out_index++
			}
			break
		case '+':
			if escaped {
				out[out_index] = '+'
				out_index++
				escaped = false
			} else {
				out[out_index] = Plus
				out_index++
			}
			break
		case '?':
			if escaped {
				out[out_index] = '?'
				out_index++
				escaped = false
			} else {
				out[out_index] = Ques
				out_index++
			}
			break
		case '(':
			if escaped {
				out[out_index] = '('
				out_index++
				escaped = false
			} else {
				out[out_index] = LParen
				out_index++
			}
			break
		case ')':
			if escaped {
				out[out_index] = ')'
				out_index++
				escaped = false
			} else {
				out[out_index] = RParen
				out_index++
			}
			break
		default:
			out[out_index] = int(regex[i])
			out_index++
			break
		}
		if escaped {
			//ERROR: Nothing consumed the escape
			fmt.Println("The escape before character: ", regex[i], " index: ", i, "was not consumed!")
		}
		escaped = nextEscape
		nextEscape = false
	}
	return out
}

