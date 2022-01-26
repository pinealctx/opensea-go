package opensea

import (
	"github.com/pinealctx/restgo"
	"net/http"
)

type Client struct {
	*restgo.Client
}

func New(fnList ...OptionFn) *Client {
	var o = defaultOption
	for _, fn := range fnList {
		fn(o)
	}
	var header = make(http.Header)
	if o.apiKey != "" {
		header.Set("X-API-KEY", o.apiKey)
	}
	header.Set("Accept", "application/json")
	return &Client{
		Client: restgo.New(restgo.WithBaseURL(o.baseURL), restgo.WithGlobalHeader(header)),
	}
}
