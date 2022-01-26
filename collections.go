package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Collections Use this endpoint to fetch collections on OpenSea
func (c *Client) Collections(ctx context.Context, req *CollectionsRequest) ([]*model.Collection, error) {
	var rsp, err = c.Get(ctx, "/api/v1/collections", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response CollectionsResponse
	err = rsp.JSONUnmarshal(&response)
	if err != nil {
		return nil, WrapperRspError(rsp)
	}
	return response.Collections, nil
}

type CollectionsRequest struct {
	// A wallet address. If specified, will return collections where the owner owns at least one asset belonging to smart contracts in the collection.
	// The number of assets the account owns is shown as owned_asset_count for each collection.
	AssetOwner string `query:"asset_owner"`
	// Offset
	Offset int32 `query:"offset,required"`
	// Limit
	Limit int32 `query:"limit,required"`
}

type CollectionsResponse struct {
	Collections []*model.Collection `json:"collections"`
}
