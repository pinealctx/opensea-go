package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Asset Used to fetch more in-depth information about an individual asset
func (c *Client) Asset(ctx context.Context, req *AssetRequest) (*model.Asset, error) {
	var rsp, err = c.Get(ctx, "/api/v1/asset/:asset_contract_address/:token_id", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response model.Asset
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type AssetRequest struct {
	// The NFT contract address for the assets
	AssetContractAddress string `path:"asset_contract_address,required"`
	// Address of the contract for this NFT
	TokenID string `path:"token_id,required"`
	// Address of an owner of the token.
	// If you include this, the response will include an ownership object that includes
	// the number of tokens owned by the address provided instead of the top_ownerships object
	// included in the standard response, which provides the number of tokens owned by each of
	// the 10 addresses with the greatest supply of the token.
	AccountAddress string `query:"account_address"`
}
