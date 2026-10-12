<p align="center">
    <img alt="Nitr" height="200" src="https://raw.githubusercontent.com/bitcav/go-memdev/master/images/logo.png" style="max-width:100%;">
</p>

<div align="center">

[![Go version](https://img.shields.io/github/go-mod/go-version/bitcav/go-memdev?style=for-the-badge&logo=go&logoColor=white)](https://github.com/bitcav/go-memdev/blob/master/go.mod)
[![CI](https://img.shields.io/github/actions/workflow/status/bitcav/go-memdev/ci.yml?branch=master&style=for-the-badge&logo=githubactions&logoColor=white)](https://github.com/bitcav/go-memdev/actions/workflows/ci.yml)
[![Release](https://img.shields.io/github/v/release/bitcav/go-memdev?style=for-the-badge&logo=github)](https://github.com/bitcav/go-memdev/releases)
[![Go Reference](https://img.shields.io/badge/pkg.go.dev-reference-3FB950?style=for-the-badge&logo=go&logoColor=white&labelColor=1B2127)](https://pkg.go.dev/github.com/bitcav/go-memdev)
[![MIT license](https://img.shields.io/badge/License-MIT-3FB950?style=for-the-badge&labelColor=1B2127)](https://github.com/bitcav/go-memdev/blob/master/LICENSE)
![Platform](https://img.shields.io/badge/platform-Linux%20%7C%20Windows-E8EDF2?style=for-the-badge&labelColor=1B2127)

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
	memInfo, _ := memdev.Info() //the returned value is a slice of Memory

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
        "type": "DDR3",
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
