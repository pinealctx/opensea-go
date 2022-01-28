# opensea-go

Golang's library for OpenSea APIs (https://docs.opensea.io/reference).

[![Go Reference](https://pkg.go.dev/badge/github.com/pinealctx/opensea-go.svg)](https://pkg.go.dev/github.com/pinealctx/opensea-go)

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
	var cli = opensea.New(opensea.WithTestNet())
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