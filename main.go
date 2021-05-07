package main

import (
	"bufio"
	"strings"

	"github.com/mishamyrt/compars/pkg/substring"
	"github.com/mishamyrt/compars/pkg/types"
)

// Parse comments from given scanner.
func Parse(s *bufio.Scanner, set types.CommentSymbolSet) []types.Comment {
	var results []types.Comment
	var line string
	var lineNumber int
	var inMultiline bool
	var multilinePart string
	var multilineStartLine int
	var multilineSet types.MultilineSet

	appendComment := func(text string, number int) {
		results = append(results, types.Comment{
			Text: substring.Trim(text),
			Line: number,
		})
	}

	for s.Scan() {
		lineNumber++
		line = s.Text()

		if inMultiline {
			if strings.Contains(line, multilineSet.End) {
				inMultiline = false
				multilinePart += substring.GetPrevious(line, multilineSet.End)
				appendComment(multilinePart, multilineStartLine)

				continue
			} else {
				multilinePart += line
				continue
			}
		} else {
			for _, m := range set.Multiline {
				if strings.Contains(line, m.Start) {
					line = substring.GetSubsequent(line, m.Start)
					if strings.Contains(line, m.End) {
						line = substring.GetPrevious(line, m.End)
						appendComment(line, lineNumber)
						continue
					} else {
						inMultiline = true
						multilineSet = m
						multilinePart = line
						multilineStartLine = lineNumber
						continue
					}
				}
			}
		}

		for _, inlineSymbol := range set.Inline {
			if strings.Contains(line, inlineSymbol) {
				line = substring.GetSubsequent(line, inlineSymbol)
				appendComment(line, lineNumber)
			}
		}
	}
	return results
}
