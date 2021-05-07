package cases

import (
	"github.com/mishamyrt/compars/pkg/symbols"
	"github.com/mishamyrt/compars/pkg/types"
)

var XMLTestCase = types.TestCase{
	Content: `
<html>
	<!-- Single line XML comment -->
	<!--
		Multiline XML comment
	-->
</html>
	`,
	Results: []string{"Single line XML comment", "Multiline XML comment"},
	Set:     symbols.ExtensionsCommentSet["html"],
}
