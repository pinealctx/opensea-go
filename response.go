package opensea

import (
	"errors"
	"github.com/json-iterator/go"
	"github.com/pinealctx/restgo"
	"net/http"
)

var ErrNotFound = errors.New("resource.not.found")

func WrapperRspError(rsp restgo.IResponse) error {
	var r = rsp.GetResponse()
	if r.StatusCode != http.StatusOK {
		switch r.StatusCode {
		case http.StatusNotFound:
			return ErrNotFound
		}
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

	var nJson = jsoniter.Config{
		MarshalFloatWith6Digits: false,
		EscapeHTML:              true,
		SortMapKeys:             true,
		UseNumber:               true,
		TagKey:                  "opensea",
		OnlyTaggedField:         false,
		ValidateJsonRawMessage:  true,
	}.Froze()

	return nJson.Unmarshal(data, i)
}
