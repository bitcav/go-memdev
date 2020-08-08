# MemDev

Crossplatform (Windows/Linux) memory devices information library

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
