package opensea

import (
	"context"
	"github.com/pinealctx/restgo"
	"net/http"
	"time"
)

const (
	APIKey = "X-API-KEY"
)

type Client struct {
	*restgo.Client
	*option
}

func New(fnList ...OptionFn) *Client {
	var o = defaultOption
	for _, fn := range fnList {
		fn(o)
	}
	var header = make(http.Header)
	if o.apiKey != "" {
		header.Set(APIKey, o.apiKey)
	}
	header.Set("Accept", "application/json")
	return &Client{
		Client: restgo.New(restgo.WithBaseURL(o.baseURL), restgo.WithGlobalHeader(header)),
		option: o,
	}
}

func (c *Client) get(ctx context.Context, resource string, params ...restgo.IParam) (restgo.IResponse, error) {
	var rsp, err = c.Execute(ctx, "GET", resource, params...)
	if err != nil {
		return nil, err
	}
	if c.retryWhenFreqLimit && rsp.StatusCode() == http.StatusTooManyRequests {
		var countValue = ctx.Value(countCtxKey)
		var count = c.retryCount
		if countValue != nil {
			count = countValue.(int)
		}
		if count != 0 {
			time.Sleep(c.retryInterval)
			return c.get(WrapperCountContext(ctx, count-1), resource, params...)
		}
	}
	return rsp, nil
}
