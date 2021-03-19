package cases

import (
	"github.com/mishamyrt/compars/v1/pkg/symbols"
	"github.com/mishamyrt/compars/v1/pkg/types"
)

var PythonTestCase = types.TestCase{
	Content: `
# Single line
"""
Multiline
"""
print("Hello, World!")
	`,
	Results: []string{"Single line", "Multiline"},
	Set:     symbols.ExtensionsCommentSet["html"],
}
