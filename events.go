package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// Events The /events endpoint provides a list of events that occur on the assets that OpenSea tracks.
// The "event_type" field indicates what type of event it is (transfer, successful auction, etc).
func (c *Client) Events(ctx context.Context, req *EventsRequest) (*EventsResponse, error) {
	var rsp, err = c.get(ctx, "/api/v1/events", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response EventsResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

type EventsRequest struct {
	// Restrict to events on OpenSea auctions. Can be true or false
	OnlyOpensea bool `query:"only_opensea"`
	// The token's id to optionally filter by
	TokenID string `query:"token_id"`
	// The NFT contract address for the assets
	AssetContractAddress string `query:"asset_contract_address"`
	// Limit responses to events from a collection. Case sensitive and must match the collection slug exactly.
	// Will return all assets from all contracts in a collection.
	// For more information on collections, see our collections documentation.
	CollectionSlug   string `query:"collection_slug"`
	CollectionEditor string `query:"collection_editor"`
	// A user account's wallet address to filter for events on an account
	AccountAddress string `query:"account_address"`
	// The event type to filter.
	EventType model.EventType `query:"event_type"`
	// Filter by an auction type.
	AuctionType model.AuctionType `query:"auction_type"`
	// Only show events listed before this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	OccurredBefore string `query:"occurred_before"`
	// Only show events listed after this timestamp. Seconds since the Unix epoch.
	// eg: 2017-07-21T17:32:28Z
	OccurredAfter string `query:"occurred_after"`
	// A cursor pointing to the page to retrieve
	Cursor string `query:"cursor"`
}

type EventsResponse struct {
	// List of Event Object
	AssetEvents []*model.Event `opensea:"asset_events"`
	// A cursor to be supplied as a query param to retrieve the next page
	Next string `opensea:"next"`
	// A cursor to be supplied as a query param to retrieve the previous page
	Previous string `opensea:"previous"`
}
