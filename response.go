package opensea

import (
	"errors"
	"fmt"
	"github.com/json-iterator/go"
	"github.com/pinealctx/restgo"
	"net/http"
)

var (
	ErrNotFound        = errors.New("resource.not.found")
	ErrTooManyRequests = errors.New("too.many.requests")
	ErrInternalServer  = errors.New("internal.server.error")
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
	var statusCode = rsp.StatusCode()
	switch statusCode {
	case http.StatusOK:
		var data, err = rsp.Data()
		if err != nil {
			return err
		}
		return jConfig.Unmarshal(data, i)
	case http.StatusBadRequest, http.StatusNotAcceptable:
		var data, err = rsp.Data()
		if err != nil {
			return err
		}
		return fmt.Errorf("bad.request:%s", data)
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusInternalServerError:
		return ErrInternalServer
	default:
		return fmt.Errorf("unknown.http.status.code: %d", statusCode)
	}
}
