package opensea

import (
	"context"
	"time"
)

const (
	hostProd = "https://api.opensea.io"
	hostTest = "https://testnets-api.opensea.io"
)

var defaultOption = &option{baseURL: hostProd}

type option struct {
	test               bool
	baseURL            string
	apiKey             string
	retryWhenFreqLimit bool
	retryInterval      time.Duration
	retryCount         int
}

type OptionFn func(*option)

func WithTestNets(test bool) OptionFn {
	return func(o *option) {
		o.test = test
		if test {
			o.baseURL = hostTest
		} else {
			o.baseURL = hostProd
		}
	}
}

func WithAPIKey(apiKey string) OptionFn {
	return func(o *option) {
		o.apiKey = apiKey
	}
}

func WithRetryWhenFreqLimit(interval time.Duration, count int) OptionFn {
	return func(o *option) {
		o.retryWhenFreqLimit = true
		o.retryInterval = interval
		o.retryCount = count
	}
}

type ctxKey int

const (
	countCtxKey ctxKey = iota
)

func WrapperCountContext(ctx context.Context, count int) context.Context {
	return context.WithValue(ctx, countCtxKey, count)
}
