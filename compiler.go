package regex

import(
	"./lexer"
	"./parser"
	"./oplist"
	"./preprocessor"
	"./vm"
	"fmt"
)

type Regex struct {
	Instructions []oplist.Instruct
}

func Compile(regex string) *Regex {
	regex = preprocessor.PreProcess(regex)
	lexed := lexer.Lex(regex)
	success, parseTree := parser.Parse(lexed)
	if !success {
		fmt.Println("Parsing Failed")
		return &Regex{}
	}
	return &Regex{parseTree.Generate().ToArray()}
}

func (regex *Regex) Match(input string) bool {
	return vm.ThompsonVM(input, regex.Instructions)
}

func Declare(name, regex string) {
	preprocessor.Variables[name] = regex
}

