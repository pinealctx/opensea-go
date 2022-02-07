package opensea

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/pinealctx/restgo"
	"net/http"
)

var (
	ErrNotFound        = errors.New("resource.not.found")
	ErrTooManyRequests = errors.New("too.many.requests")
)

var (
	jConfig = jsoniter.Config{
		MarshalFloatWith6Digits: false,
		EscapeHTML:              true,
		SortMapKeys:             true,
		UseNumber:               true,
		TagKey:                  "opensea",
		OnlyTaggedField:         false,
		ValidateJsonRawMessage:  true,
	}.Froze()
)

func ParseRsp(rsp restgo.IResponse, i interface{}) error {
	switch rsp.StatusCode() {
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	}

	var data, err = rsp.Data()
	if err != nil {
		return err
	}

	return jConfig.Unmarshal(data, i)
}
