package opensea

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/pinealctx/restgo"
	"net/http"
)

var ErrNotFound = errors.New("resource.not.found")

func WrapperRspError(rsp restgo.IResponse) error {
	switch rsp.StatusCode() {
	case http.StatusNotFound:
		return ErrNotFound
	default:
	}

	var ret, err = rsp.Data()
	if err != nil {
		return err
	}
	return errors.New(string(ret))
}

func ParseRsp(rsp restgo.IResponse, i interface{}) error {
	var data, err = rsp.Data()
	if err != nil {
		return err
	}

	var n = jsoniter.Config{
		MarshalFloatWith6Digits: false,
		EscapeHTML:              true,
		SortMapKeys:             true,
		UseNumber:               true,
		TagKey:                  "opensea",
		OnlyTaggedField:         false,
		ValidateJsonRawMessage:  true,
	}.Froze()

	return n.Unmarshal(data, i)
}
