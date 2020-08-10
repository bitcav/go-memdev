<p align="center">
    <img alt="Nitr" height="200" src="https://raw.githubusercontent.com/bitcav/go-memdev/master/images/logo.png" style="max-width:100%;">
</p>

<div align="center">

![go](https://raw.githubusercontent.com/bitcav/nitr/master/images/goversion.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/bitcav/go-memdev)](https://goreportcard.com/report/github.com/bitcav/go-memdev)
[![GoDoc](https://godoc.org/github.com/narqo/go-badge?status.svg)](https://godoc.org/github.com/bitcav/go-memdev)
[![Build Status](https://travis-ci.org/bitcav/go-memdev.svg?branch=master)](https://travis-ci.org/bitcav/go-memdev)
![preview](https://img.shields.io/badge/platform-linux%20%7C%20%20win-lightgrey)
[![MIT license](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/bitcav/go-memdev/blob/master/LICENSE)

</div>
 
# go-memdev

A Go package that access to **Memory Devices** information.

## Installation
```
go get -u github.com/bitcav/go-memdev
```

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/bitcav/go-memdev"
)

func main() {
	memInfo, _ := memdev.Info() //the returned value is a struct

	jsonOutput, _ := json.MarshalIndent(memInfo, "", "    ")

	fmt.Println(string(jsonOutput))
}

```

The output is below.

```json
[
    {
        "bank": "DIMM A",
        "size": 4096,
        "unit": "MB",
        "type": "FBD2",
        "formFactor": "SODIMM",
        "manufacturer": "Samsung",
        "serial": "A49F8D93",
        "assetTag": "03153300",
        "partNumber": "M471B5173DB0-YK0  ",
        "speed": 1600,
        "dataWidth": 64,
        "totalWidth": 64
    }
]

```


## Running

### Build
```
go build
```

### Run
:lock: In order to access the system BIOS requires running with elevated privileges.

* Linux:
```
sudo ./main
```

* Windows:

You can launch the program as an administrator by right-clicking on the executable file and choosing "Run as administrator."
