package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Collection Used for retrieving more in-depth information about individual collections,
// including real time statistics like floor price
func (c *Client) Collection(ctx context.Context, req *CollectionRequest) (*model.Collection, error) {
	var rsp, err = c.Get(ctx, "/api/v1/collection/:collection_slug", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response CollectionResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, WrapperRspError(rsp)
	}
	return response.Collection, nil
}

type CollectionRequest struct {
	CollectionSlug string `path:"collection_slug,required"`
}

type CollectionResponse struct {
	Collection *model.Collection `opensea:"collection"`
}
