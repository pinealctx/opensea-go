package opensea

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/pinealctx/restgo"
)

func WrapperRspError(rsp restgo.IResponse) error {
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

	var njson = jsoniter.Config{
		MarshalFloatWith6Digits: false,
		EscapeHTML:              true,
		SortMapKeys:             true,
		UseNumber:               true,
		TagKey:                  "opensea",
		OnlyTaggedField:         false,
		ValidateJsonRawMessage:  true,
	}.Froze()

	return njson.Unmarshal(data, i)
}
