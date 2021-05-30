package symbols

import (
	"errors"

	"github.com/mishamyrt/compars/pkg/types"
)

// ExtensionsCommentSet is map with file extensions and comment sets.
var ExtensionsCommentSet = map[string]types.CommentSymbolSet{
	"js":   CommentSymbols["c-style"],
	"jsx":  CommentSymbols["c-style"],
	"ts":   CommentSymbols["c-style"],
	"tsx":  CommentSymbols["c-style"],
	"go":   CommentSymbols["c-style"],
	"php":  CommentSymbols["php"],
	"py":   CommentSymbols["python"],
	"html": CommentSymbols["xml"],
	"xml":  CommentSymbols["xml"],
	"sh":   CommentSymbols["sh"],
	"bash": CommentSymbols["sh"],
}

// GetSetByExtension returns comment symbol set for the extension
func GetSetByExtension(ext string) (types.CommentSymbolSet, error) {
	ext = ext[1 : len(ext)-1]
	if val, ok := ExtensionsCommentSet[ext]; ok {
		return val, nil
	}
	return types.CommentSymbolSet{}, errors.New("extension isn't supported yet")
}
