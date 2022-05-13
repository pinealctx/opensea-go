package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Offers This endpoint is used to fetch the set of active offers on a given NFT.
func (c *Client) Offers(ctx context.Context, req *OffersRequest) (*OffersResponse, error) {
	var rsp, err = c.get(ctx, "/api/v1/asset/:asset_contract_address/:token_id/offers", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response OffersResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type OffersRequest struct {
	// Address of the contract for this NFT
	AssetContractAddress string `path:"asset_contract_address,required"`
	// Token ID for this item
	TokenID string `path:"token_id,required"`
	// Limit. Defaults to 20, capped at 50.
	Limit int32 `query:"limit"`
}

type OffersResponse struct {
	Offers []*model.Order `opensea:"offers"`
}
