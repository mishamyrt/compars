package comments_test

import (
	"bufio"
	"testing"

	"github.com/mishamyrt/compars/v1/comments"
	"github.com/mishamyrt/compars/v1/testhelpers"
	"github.com/mishamyrt/compars/v1/types"
)

type CommentParser func(s *bufio.Scanner, set types.CommentSymbolSet) []types.Comment

type TestCase struct {
	Count int
	Text  string
	Set   types.CommentSymbolSet
}

var Cases = []TestCase{
	{
		// For the first case, I set a symbol set directly to test the algorithm itself, not the data
		Set: types.CommentSymbolSet{
			Inline:         "//",
			MultilineStart: "/*",
			MultilineEnd:   "*/",
		},
		Count: 3,
		Text: `
			// Comment
			/*
			* Comment
			*/
			/*
			Comment
			*/`,
	},
	{
		Set:   comments.CommentSymbols["python"],
		Count: 2,
		Text: `
			# Comment
			'''
			Comment
			'''`,
	},
	{
		Set:   comments.CommentSymbols["html"],
		Count: 1,
		Text:  "<!-- Comment -->",
	},
}

func ParseSuite(t *testing.T, parse CommentParser) {
	for _, c := range Cases {
		res := parse(testhelpers.ScannerFrom(c.Text), c.Set)
		if len(res) != c.Count {
			t.Errorf("Wrong count: %d vs %d", len(res), c.Count)
			t.Fail()
		}
		for _, comment := range res {
			if comment.Text != "Comment" {
				t.Errorf("Wrong text: \"%s\" from \"%s\"", comment.Text, c.Text)
				t.Fail()
			}
		}
	}
}

func TestParse(t *testing.T) {
	ParseSuite(t, comments.Parse)
}
