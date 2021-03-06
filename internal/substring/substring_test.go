package substring_test

import (
	"testing"

	"github.com/mishamyrt/compars/internal/substring"
)

type Trimmer func(text string) string

type Subsequenter func(delimeter string, text string) string
type Previouser func(delimeter string, text string) string
type Midster func(startDel string, endDel string, text string) string

func TrimSuite(t *testing.T, trim Trimmer) {
	res := trim(" * Test 	")
	if res != "Test" {
		t.Errorf("Wrong trim: \"%s\"", res)
		t.Fail()
	}
}

func SubsequentSuite(t *testing.T, sub Subsequenter) {
	testString := "Junk :Test"
	res := sub(testString, ":")
	if res != "Test" {
		t.Errorf("Wrong subsequent: \"%s\"", res)
		t.Fail()
	}
	res = sub(testString, "*")
	if res != "" {
		t.Errorf("Wrong subsequent: \"%s\"", res)
		t.Fail()
	}
}

func PreviousSuite(t *testing.T, prev Previouser) {
	testString := "Test: Junk"
	res := prev(testString, ":")
	if res != "Test" {
		t.Errorf("Wrong previous: \"%s\"", res)
		t.Fail()
	}
	res = prev(testString, "*")
	if res != "" {
		t.Errorf("Wrong subsequent: \"%s\"", res)
		t.Fail()
	}
}

func TestTrim(t *testing.T) {
	TrimSuite(t, substring.Trim)
}

func TestSubsequent(t *testing.T) {
	SubsequentSuite(t, substring.GetSubsequent)
}

func TestPrevious(t *testing.T) {
	PreviousSuite(t, substring.GetPrevious)
}
