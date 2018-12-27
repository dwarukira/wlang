package evaluator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dwarukira/wakanda/object"
)

var builtins = map[string]*object.Builtin{
	//array and string helper funcations
	"len": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}

			default:
				return newError("argument to `len` not supported, got %s",
					args[0].Type())
			}
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				if len(arr.Elements) > 0 {
					return arr.Elements[0]
				}
			case *object.String:
				str := args[0].(*object.String).Value
				length := len(str)
				if length > 0 {
					return &object.String{Value: string(str[0])}
				}
			default:
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			return NULL
		},
	},

	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					return arr.Elements[length-1]
				}
			case *object.String:
				str := args[0].(*object.String).Value
				length := len(str)
				if length > 0 {
					return &object.String{Value: string(str[length-1])}
				}
			default:
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			switch args[0].(type) {
			case *object.Array:
				arr := args[0].(*object.Array)
				length := len(arr.Elements)
				if length > 0 {
					newElements := make([]object.Object, length-1, length-1)
					copy(newElements, arr.Elements[1:length])
					return &object.Array{Elements: newElements}
				}
			case *object.String:
				str := args[0].(*object.String).Value
				length := len(str)
				if length > 0 {
					return &object.String{Value: str[1:length]}

				}
			default:
				return newError("argument to `first` must be ARRAY or STRING, got %s",
					args[0].Type())
			}
			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 && args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	// io
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"input": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}
			scanner := bufio.NewScanner(os.Stdin)
			scanned := scanner.Scan()
			if !scanned {
				return NULL
			}
			return &object.String{Value: scanner.Text()}
		},
	},
}
