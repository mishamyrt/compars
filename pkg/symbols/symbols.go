package symbols

import "github.com/mishamyrt/compars/pkg/types"

var cStyleMultiline = types.MultilineSet{
	Start: "/*",
	End:   "*/",
}

// CommentSymbols is allowed comment symbol set.
var CommentSymbols = map[string]types.CommentSymbolSet{
	"c-style": {
		Inline:    []string{"//"},
		Multiline: []types.MultilineSet{cStyleMultiline},
	},
	"python": {
		Inline: []string{"#"},
		Multiline: []types.MultilineSet{
			{
				Start: "'''",
				End:   "'''",
			},
			{
				Start: "\"\"\"",
				End:   "\"\"\"",
			},
		},
	},
	"xml": {
		Inline: []string{},
		Multiline: []types.MultilineSet{
			{
				Start: "<!--",
				End:   "-->",
			},
		},
	},
	"php": {
		Inline:    []string{"//", "#"},
		Multiline: []types.MultilineSet{cStyleMultiline},
	},
	"sh": {
		Inline: []string{"#"},
	},
}
