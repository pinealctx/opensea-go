# opensea-go

[![Go Reference](https://pkg.go.dev/badge/github.com/pinealctx/opensea-go.svg)](https://pkg.go.dev/github.com/pinealctx/opensea-go)
[![codecov](https://codecov.io/gh/pinealctx/opensea-go/branch/main/graph/badge.svg?token=36396OPMPT)](https://codecov.io/gh/pinealctx/opensea-go)
[![golangci-lint](https://github.com/pinealctx/opensea-go/actions/workflows/ci.yml/badge.svg)](https://github.com/pinealctx/opensea-go/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pinealctx/opensea-go)](https://goreportcard.com/report/github.com/pinealctx/opensea-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Golang's library for OpenSea APIs (https://docs.opensea.io/reference).

## Get Started

#### Get it

```shell
go get -u github.com/pinealctx/opensea-go
```

#### Use it

```go
package main

import (
	"context"
	"github.com/pinealctx/opensea-go"
	"log"
)

func main() {
	var cli = opensea.New(opensea.WithTestNets(true))
	var asset, err = cli.Asset(context.Background(), &opensea.AssetRequest{
		AssetContractAddress: "0x66583bd73a27c9245b723ff6a58f872234c3a50a",
		TokenID:              "3",
	})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(asset)
}
```
