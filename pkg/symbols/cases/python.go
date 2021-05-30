package cases

import (
	"github.com/mishamyrt/compars/pkg/symbols"
	"github.com/mishamyrt/compars/pkg/types"
)

// PythonTestCase is comment parser test case for Python.
var PythonTestCase = types.TestCase{
	Content: `
# Single line
'''
Multiline single
'''
"""
Multiline quotes
"""
print("Hello, World!")
	`,
	Results: []string{"Single line", "Multiline single", "Multiline quotes"},
	Set:     symbols.ExtensionsCommentSet["py"],
}
