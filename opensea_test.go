package opensea

import "context"

var (
	ctx = context.Background()
)

func newClient() *Client {
	return New(WithTestNets(true))
}
