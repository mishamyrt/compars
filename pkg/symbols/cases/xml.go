package cases

import (
	"github.com/mishamyrt/compars/v1/pkg/symbols"
	"github.com/mishamyrt/compars/v1/pkg/types"
)

var XMLTestCase = types.TestCase{
	Content: `
<html>
	<!-- Single line -->
	<!--
		Multiline
	-->
</html>
	`,
	Results: []string{"Single line", "Multiline"},
	Set:     symbols.ExtensionsCommentSet["html"],
}
