package types

// TestCase for symbols set.
type TestCase struct {
	Content string
	Results []string
	Set     CommentSymbolSet
}

// Config map of keyword with error level.
type Config map[string]uint8

// KewordsMap is a map where key is keyword and value is level.
type KewordsMap map[string]string

// ConfigFile is list with keywords and error levels.
type ConfigFile struct {
	Keywords KewordsMap
}

// Comment with content and position.
type Comment struct {
	Line int
	Text string
}

// MultilineSet is language-specific mutliline symbol set.
type MultilineSet struct {
	Start string
	End   string
}

// CommentSymbolSet is language-specific symbol set.
type CommentSymbolSet struct {
	Inline    []string
	Multiline []MultilineSet
}
