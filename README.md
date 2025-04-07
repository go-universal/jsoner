# Jsoner (Extended GoLang JSON Marshal)

![GitHub Tag](https://img.shields.io/github/v/tag/go-universal/jsoner?sort=semver&label=version)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-universal/jsoner.svg)](https://pkg.go.dev/github.com/go-universal/jsoner)
[![License](https://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/go-universal/jsoner/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-universal/jsoner)](https://goreportcard.com/report/github.com/go-universal/jsoner)
![Contributors](https://img.shields.io/github/contributors/go-universal/jsoner)
![Issues](https://img.shields.io/github/issues/go-universal/jsoner)

`jsoner` is a Go package that provides utilities for marshaling Go structs into JSON with advanced filtering capabilities. It allows you to exclude specific fields or nested paths from the JSON output.

## Installation

To install the package, run:

```bash
go get github.com/go-universal/jsoner
```

## Usage

Import the package in your Go code:

```go
import "github.com/go-universal/jsoner"
```

### Marshal

Converts the input value (`v`) into JSON format, excluding specified fields or paths.

```go
func Marshal(v any, exclude ...string) ([]byte, error)
```

```go
type Book struct {
    Title string
    ISBN  string `json:"isbn"`
}

type Author struct {
    Name   string `json:"name"`
    Family string `json:"family"`
    Books  []Book `json:"author_books"`
}

author := Author{
    Name:   "John",
    Family: "Doe",
    Books: []Book{
        {Title: "Book 1", ISBN: "12345"},
        {Title: "Book 2", ISBN: "67890"},
    },
}

jsonData, err := jsoner.Marshal(author, "family", "author_books.isbn")
if err != nil {
    panic(err)
}

fmt.Println(string(jsonData))
// Output: {"name":"John","author_books":[{"Title":"Book 1"},{"Title":"Book 2"}]}
```

### MarshalIndent

Converts the input value (`v`) into indented JSON format, excluding specified fields or paths.

```go
func MarshalIndent(v any, indent string, exclude ...string) ([]byte, error)
```

```go
type Book struct {
    Title string
    ISBN  string `json:"isbn"`
}

type Author struct {
    Name   string `json:"name"`
    Family string `json:"family"`
    Books  []Book `json:"author_books"`
}

author := Author{
    Name:   "John",
    Family: "Doe",
    Books: []Book{
        {Title: "Book 1", ISBN: "12345"},
        {Title: "Book 2", ISBN: "67890"},
    },
}

jsonData, err := jsoner.MarshalIndent(author, "  ", "family", "author_books.isbn")
if err != nil {
    panic(err)
}

fmt.Println(string(jsonData))
// Output:
// {
//   "name": "John",
//   "author_books": [
//     {
//       "Title": "Book 1"
//     },
//     {
//       "Title": "Book 2"
//     }
//   ]
// }
```

## License

This project is licensed under the ISC License. See the [LICENSE](LICENSE) file for details.
