[![GoDoc](https://godoc.org/github.com/narqo/go-badge?status.svg)](https://godoc.org/github.com/bitcav/go-memdev)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitcav/go-memdev)](https://goreportcard.com/report/github.com/bitcav/go-memdev)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/bitcav/go-memdev/blob/master/LICENSE)
 
# Go-MemDev

Crossplatform (Windows/Linux) **memory devices** information library.

## Installation
```
go get github.com/bitcav/go-memdev
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/bitcav/go-memdev"
)

func main() {
	memInfo, err := memdev.Info()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(memInfo)
}

```
