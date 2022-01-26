package opensea

import (
	"errors"
	"github.com/pinealctx/restgo"
)

func WrapperRspError(rsp restgo.IResponse) error {
	var ret, err = rsp.Data()
	if err != nil {
		return err
	}
	return errors.New(string(ret))
}
