package compars_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/mishamyrt/compars"
	"github.com/mishamyrt/compars/pkg/symbols/cases"
	"github.com/mishamyrt/compars/pkg/types"
)

// ScannerFrom creates bufio.Scanner instance from string.
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
	cases.PHPTestCase,
	cases.ShellTestCase,
}

func RunTests(t *testing.T, parse CommentParser) {
	total := len(testCases)
	for i, c := range testCases {
		t.Logf("Running %d from %d\n", i+1, total)
		t.Log("Set: ", c.Set)
		res := parse(ScannerFrom(c.Content), c.Set)
		if len(res) != len(c.Results) {
			t.Logf("Obtained: %s", strings.Join(getStrings(res), ", "))
			t.Logf("Expected: %s", strings.Join(c.Results, ", "))
			t.Errorf("Wrong count: %d. Should be %d", len(res), len(c.Results))
			t.Fail()
		}

		for i, comment := range res {
			if i > len(c.Results)-1 {
				t.Errorf("Unknown text: \"%s\"", comment.Text)
				t.Fail()
			} else if comment.Text != c.Results[i] {
				t.Errorf("Wrong text: \"%s\". Should be \"%s\"", comment.Text, c.Results[i])
				t.Fail()
			}
		}
	}
}

func TestParser(t *testing.T) {
	RunTests(t, compars.Parse)
}
