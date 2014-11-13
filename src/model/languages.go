// languages
package model

import (
//	"fmt"
//	"sort"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/
type Language struct {
	Id       int
	Name     string
	Keywords []string
}

var Languages = map[string]string{"C#": "csharp", "Go": "go", "Python": "python",
	"Java": "java", "JavaScript": "javascript", "CSS": "css",
	"PHP": "php", "Ruby": "ruby", "Scala": "scala"}

/*
var Languages = []Language{
	Language{0, "Text", nil},
	Language{1, "Go",
		[]string{
			"break", "case", "chan", "const",
			"continue", "default", "defer", "else", "fallthrough", "for", "func",
			"go", "goto", "if", "import", "interface", "map", "package", "range",
			"return", "select", "struct", "switch", "type", "var",
		},
	},
	Language{2, "PHP",
		[]string{
			"abstract", "and", "as", "break", "callable", "case", "catch", "class",
			"clone", "const", "continue", "declare", "default", "do", "echo",
			"else", "elseif", "enddeclare", "endfor", "endforeach", "endif",
			"endswitch", "endwhile", "extends", "final", "finally", "for",
			"foreach", "function", "global", "goto", "if", "implements", "include",
			"include_once", "instanceof", "insteadof", "interface", "namespace",
			"new", "or", "print", "private", "protected", "public", "require",
			"require_once", "return", "static", "switch", "throw", "trait",
			"try", "use", "var", "while", "xor", "yield",
		},
	},
}
*/
