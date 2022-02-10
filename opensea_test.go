package opensea

import (
	"context"
	"time"
)

var (
	ctx = context.Background()
)

func newClient() *Client {
	return New(WithTestNets(true), WithRetryWhenFreqLimit(time.Second, 10))
}
