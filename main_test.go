package main

import (
	"bufio"
	"fmt"
	"strings"
	"testing"

	"github.com/mishamyrt/compars/v1/pkg/symbols/cases"
	"github.com/mishamyrt/compars/v1/pkg/types"
)

// ScannerFrom creates bufio.Scanner instance from string
func ScannerFrom(input string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(input))
}

type CommentParser func(s *bufio.Scanner, set types.CommentSymbolSet) []types.Comment

func getStrings(r []types.Comment) (comments []string) {
	for _, c := range r {
		comments = append(comments, c.Text)
	}
	return
}

var testCases = []types.TestCase{
	cases.PythonTestCase,
	cases.XMLTestCase,
}

func RunTests(t *testing.T, parse CommentParser) {
	total := len(testCases)
	for i, c := range testCases {
		fmt.Printf("Running %d from %d\n", i, total)
		res := parse(ScannerFrom(c.Content), c.Set)
		if len(res) != len(c.Results) {
			t.Errorf("Wrong count: \"%d\". Should be \"%d\"", len(res), len(c.Results))
			t.Logf("Obtained: %s", strings.Join(getStrings(res), ", "))
			t.Logf("Expected: %s", strings.Join(c.Results, ", "))
			t.Fail()
		}
		for i, comment := range res {
			if comment.Text != c.Results[i] {
				t.Errorf("Wrong text: \"%s\". Should be \"%s\"", comment.Text, c.Results[i])
				t.Fail()
			}
		}
	}
}

func TestParserImpl(t *testing.T) {
	RunTests(t, Parse)
}
