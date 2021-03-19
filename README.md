# Compars [![release](https://img.shields.io/github/v/tag/mishamyrt/compars?sort=semver)](https://github.com/mishamyrt/compars/tags)

<img align="right" width="128" height="128"
     alt="Logo"
     src="./assets/logo@2x.png">

Parser of comments in source code. Supports multiple languages through dictionaries. The library contains a pure `Parse` function that returns found comments and their positions

## Features

* Mutli–language
* Pure functions
* Memory–frendly

## Usage

The function needs to pass a `*bufio.Scanner` and a structure of `CommentSymbolSet`. The set can be obtained from the `GetSetByExtension` function, which takes a file extension. The available extensions may be found in the [extensions.go](./pkg/symbols/extensions.go) file.

```go
package main
import (
    "github.com/mishamyrt/compars/v1"
    "github.com/mishamyrt/compars/v1/pkg/symbols"
)

func main() {
    file, _ := os.Open(path)
	set, _ := symbols.GetSetByExtension("go")
    comments := compars.Parse(file, set)
    fmt.Printf("%d comments are found", len(comments))
}
```

The result is represented by an array of comment structures.

```go
// CommentSymbolSet is language-specific symbol set
type CommentSymbolSet struct {
	Inline         string
	MultilineStart string
	MultilineEnd   string
}
```

## Documentation

Documentation about Compars commands and libraries can be found at usual [godoc.org](https://godoc.org/github.com/mishamyrt/compars).

## License

[MIT](./LICENSE).