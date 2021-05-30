package cases

import (
	"github.com/mishamyrt/compars/pkg/symbols"
	"github.com/mishamyrt/compars/pkg/types"
)

var ShellTestCase = types.TestCase{
	Content: `
# Single line comment
go get "github.com/mishamyrt/compars"
	`,
	Results: []string{"Single line comment"},
	Set:     symbols.ExtensionsCommentSet["sh"],
}
