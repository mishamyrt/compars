package cases

import (
	"github.com/mishamyrt/compars/pkg/symbols"
	"github.com/mishamyrt/compars/pkg/types"
)

// PHPTestCase is comment parser test case for PHP.
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
