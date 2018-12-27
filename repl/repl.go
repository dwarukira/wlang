package repl

import (
	"io"
	"io/ioutil"
	"os"

	prompt "github.com/c-bata/go-prompt"
	"github.com/dwarukira/wakanda/evaluator"
	"github.com/dwarukira/wakanda/lexer"
	"github.com/dwarukira/wakanda/object"
	"github.com/dwarukira/wakanda/parser"
)

const PROMPT = ">> "

func completer(d prompt.Document) []prompt.Suggest {

	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

// func Executor(input string) {
// 	l := lexer.New(input)
// 	p := parser.New(l)
// 	program := p.ParseProgram()
// 	if len(p.Errors()) != 0 {
// 		printParserErrors(out, p.Errors())
// 		continue
// 	}

// 	evaluated := evaluator.Eval(program, env)
// 	if evaluated != nil {
// 		io.WriteString(out, evaluated.Inspect())
// 		io.WriteString(out, "\n")
// 	}

// }

func Start(in io.Reader, out io.Writer) {
	// scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	env = ReadFile("lib/buildin.wk", os.Stdout, env)

	exc := func(input string) {
		l := lexer.New(input)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			return
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, green(evaluated.Inspect()))
			io.WriteString(out, "\n")
		}

	}

	com := func(d prompt.Document) []prompt.Suggest {
		data := env.All()
		suggestions := []prompt.Suggest{
			{Text: "let"},
			{Text: "fn"},
			{Text: "if"},
		}

		for k := range data {
			suggest := prompt.Suggest{
				Text: k,
			}
			suggestions = append(suggestions, suggest)
		}
		return prompt.FilterHasPrefix(suggestions, d.GetWordBeforeCursor(), true)
	}
	p := prompt.New(exc, com, prompt.OptionPrefix(">> "))
	p.Run()

	// for {
	// 	fmt.Printf(PROMPT)
	// 	scanned := scanner.Scan()
	// 	if !scanned {
	// 		return
	// 	}

	// 	line := scanner.Text()
	// 	l := lexer.New(line)
	// 	p := parser.New(l)

	// 	program := p.ParseProgram()
	// 	if len(p.Errors()) != 0 {
	// 		printParserErrors(out, p.Errors())
	// 		continue
	// 	}

	// 	evaluated := evaluator.Eval(program, env)
	// 	if evaluated != nil {
	// 		io.WriteString(out, evaluated.Inspect())
	// 		io.WriteString(out, "\n")
	// 	}
	// }
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

func ReadFile(file string, out io.Writer, env *object.Environment) *object.Environment {
	// scanner := bufio.NewScanner(in)
	data, _ := ioutil.ReadFile(file)
	l := lexer.New(string(data))
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return nil
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

	return env
}
