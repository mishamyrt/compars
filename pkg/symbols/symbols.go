package symbols

import "github.com/mishamyrt/compars/v1/pkg/types"

// CommentSymbols is allowed comment symbol set
var CommentSymbols = map[string]types.CommentSymbolSet{
	"c-style": {
		Inline:         "//",
		MultilineStart: "/*",
		MultilineEnd:   "*/",
	},
	"python": {
		Inline:         "#",
		MultilineStart: "'''",
		MultilineEnd:   "'''",
	},
	"xml": {
		Inline:         "",
		MultilineStart: "<!--",
		MultilineEnd:   "-->",
	},
}