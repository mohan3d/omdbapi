# omdbapi
[![GoDoc](https://godoc.org/github.com/mohan3d/omdbapi?status.svg)](https://godoc.org/github.com/mohan3d/omdbapi)
[![Build Status](https://api.travis-ci.org/mohan3d/omdbapi.svg?branch=master)](https://travis-ci.org/mohan3d/omdbapi)
[![Go Report Card](https://goreportcard.com/badge/github.com/mohan3d/omdbapi)](https://goreportcard.com/report/github.com/mohan3d/omdbapi)

Golang package to access the OMDbAPI https://www.omdbapi.com

## Installation
```bash
$ go get github.com/mohan3d/omdbapi
```

## Usage
```go
package main

import (
	"fmt"

	"github.com/mohan3d/omdbapi"
)

func savePoster(poster []byte, path string) {
	// Implement your own save function.
}

func main() {
	client := omdbapi.New("<OMDBAPI_KEY>")

	movieInfo, err := client.Title("Enemy at the gates")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Release Year: %v\nGenre: %v\n", movieInfo.Year, movieInfo.Genre)

	movieInfo, err = client.ID("tt0215750")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Title: %v\n", movieInfo.Title)

	poster, err := client.Poster("tt0215750")

	if err != nil {
		panic(err)
	}

	savePoster(poster, "<FILE_PATH>")
}
```

## Testing
**OMDBAPI_KEY** must be exported to environment variables before running tests.

```bash
$ export OMDBAPI_KEY=<YOUR_OMDBAPI_KEY>
```