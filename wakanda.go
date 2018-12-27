package main

import (
	"io"
	"io/ioutil"

	"github.com/dwarukira/wakanda/evaluator"
	"github.com/dwarukira/wakanda/lexer"
	"github.com/dwarukira/wakanda/object"
	"github.com/dwarukira/wakanda/parser"
)

func ReadFile(file string, out io.Writer) {
	// scanner := bufio.NewScanner(in)
	data, _ := ioutil.ReadFile(file)
	env := object.NewEnvironment()
	l := lexer.New(string(data))
	p := parser.New(l)

	program := p.ParseProgram()
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
