package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/dwarukira/wakanda/evaluator"
	"github.com/dwarukira/wakanda/lexer"
	"github.com/dwarukira/wakanda/object"
	"github.com/dwarukira/wakanda/parser"
	"github.com/dwarukira/wakanda/repl"
)

func ReadFile(file string, out io.Writer) {
	// scanner := bufio.NewScanner(in)
	data, _ := ioutil.ReadFile(file)
	env := object.NewEnvironment()
	env = repl.ReadFile("lib/buildin.wk", os.Stdout, env)
	l := lexer.New(string(data))
	p := parser.New(l)
	program := p.ParseProgram()
	// fmt.Println(program)
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
