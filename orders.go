package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Orders How to fetch orders from the OpenSea system
func (c *Client) Orders(ctx context.Context, req *OrdersRequest) (*OrdersResponse, error) {
	var params = restgo.ObjectParams(req)
	if c.test {
		params = append(params, restgo.NewURLQueryParam(APIKey, c.apiKey))
	}
	var rsp, err = c.get(ctx, "/wyvern/v1/orders", params...)
	if err != nil {
		return nil, err
	}
	var response OrdersResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type OrdersRequest struct {
	// Filter by smart contract address for the asset category.
	// Needs to be defined together with token_id or token_ids.
	AssetContractAddress string `query:"asset_contract_address"`
	// Filter by the address of the smart contract of the payment token that is accepted or offered by the order
	PaymentTokenAddress string `query:"payment_token_address"`
	// Filter by the order maker's wallet address
	Maker string `query:"maker"`
	// Filter by the order taker's wallet address.
	// Orders open for any taker have the null address as their taker.
	Taker string `query:"taker"`
	// Filter by the asset owner's wallet address
	Owner string `query:"owner"`
	// When "true", only show English Auction sell orders,
	// which wait for the highest bidder. When "false", exclude those.
	IsEnglish bool `query:"is_english"`
	// Only show orders for bundles of assets
	Bundled bool `query:"bundled,required"`
	// Include orders on bundles where all assets in the bundle share the address provided
	// in asset_contract_address or where the bundle's maker is the address provided in owner
	IncludeBundled bool `query:"include_bundled,required"`
	// Only show orders listed after this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	ListedAfter string `query:"listed_after"`
	// Only show orders listed before this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	ListedBefore string `query:"listed_before"`
	// Filter by the token ID of the order's asset.
	// Needs to be defined together with asset_contract_address.
	TokenID string `query:"token_id"`
	// Filter by a list of token IDs for the order's asset.
	// Needs to be defined together with asset_contract_address.
	TokenIDs []string `query:"token_ids"`
	// Filter by the side of the order.
	// 0 for buy orders and 1 for sale orders.
	Side model.Side `query:"side,required"`
	// Filter by the kind of sell order.
	// 0 for fixed-price sales or min-bid auctions,
	// and 1 for declining-price Dutch Auctions.
	// use only_english=true for filtering for only English Auctions
	SaleKind model.SaleKind `query:"sale_kind"`
	// Limit
	Limit int32 `query:"limit,required"`
	// Offset
	Offset int32 `query:"offset,required"`
	// How to sort the orders. Can be created_date for when they were made,
	// or eth_price to see the lowest-priced orders first (converted to their ETH values).
	// eth_price is only supported when asset_contract_address and token_id are also defined.
	OrderBy string `query:"order_by"`
	// Can be asc or desc for ascending or descending sort.
	// For example, to see the cheapest orders, do order_direction asc and order_by eth_price.
	OrderDirection string `query:"order_direction,required"`
}

type OrdersResponse struct {
	Count  int32          `opensea:"count"`
	Orders []*model.Order `opensea:"orders"`
}
