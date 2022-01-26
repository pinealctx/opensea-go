package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Bundles are groups of items for sale on OpenSea.
// You can buy them all at once in one transaction, and you can create them without any transactions or gas,
// as long as you've already approved the assets inside.
func (c *Client) Bundles(ctx context.Context, req *BundlesRequest) ([]*model.Bundle, error) {
	var rsp, err = c.Get(ctx, "/api/v1/bundles", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response BundlesResponse
	err = rsp.JSONUnmarshal(&response)
	if err != nil {
		return nil, WrapperRspError(rsp)
	}
	return response.Bundles, nil
}

type BundlesRequest struct {
	// If set to true, only show bundles currently on sale.
	// If set to false, only show bundles that have been sold or cancelled.
	OnSale bool `json:"on_sale"`
	// The address of the owner of the assets
	Owner string `query:"owner"`
	// The NFT contract address for the assets
	AssetContractAddress string `query:"asset_contract_address"`
	// An array of contract addresses to search for (e.g. ?asset_contract_addresses=0x1...&asset_contract_addresses=0x2...).
	// Will return a list of assets with contracts matching any of the addresses in this array.
	// If "token_ids" is also specified, then it will only return assets that match each (address, token_id) pairing, respecting order.
	AssetContractAddresses []string `query:"asset_contract_addresses"`
	// An array of token IDs to search for (e.g. ?token_ids=1&token_ids=209).
	// Will return a list of assets with token_id matching any of the IDs in this array.
	TokenIDs []string `query:"token_ids"`
	// Limit
	Limit int32 `query:"limit,required"`
	// Offset
	Offset int32 `query:"offset,required"`
}

type BundlesResponse struct {
	Bundles []*model.Bundle `json:"bundles"`
}
