package opensea

import (
	"context"
	"os"
	"time"
)

const (
	envAPIKey = "OPENSEA_API_KEY"
)

var (
	ctx = context.Background()
)

func newClient() *Client {
	return New(
		WithTestNets(false),
		WithAPIKey(os.Getenv(envAPIKey)),
		WithRetryWhenFreqLimit(time.Second, 10),
	)
}
