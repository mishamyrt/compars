package symbols

import "github.com/mishamyrt/compars/v1/pkg/types"

// CommentSymbols is allowed comment symbol set
var CommentSymbols = map[string]types.CommentSymbolSet{
	"c-style": {
		Inline: []string{"//"},
		Multiline: []types.MultilineSet{
			{
				Start: "/*",
				End:   "*/",
			},
		},
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
}
