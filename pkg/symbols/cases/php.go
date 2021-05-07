package cases

import (
	"github.com/mishamyrt/compars/v1/pkg/symbols"
	"github.com/mishamyrt/compars/v1/pkg/types"
)

var PHPTestCase = types.TestCase{
	Content: `
<?php
// Single line
# Another single line
/*
	Multiline
*/
echo 'Simple test';
	`,
	Results: []string{"Single line", "Another single line", "Multiline"},
	Set:     symbols.ExtensionsCommentSet["php"],
}
