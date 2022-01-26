package opensea

import (
	"context"
	"github.com/pinealctx/opensea-go/model"
	"github.com/pinealctx/restgo"
)

// CollectionStats Use this endpoint to fetch stats for a specific collection,
// including realtime floor price statistics
func (c *Client) CollectionStats(ctx context.Context, req *CollectionRequest) (*model.CollectionStats, error) {
	var rsp, err = c.Get(ctx, "/api/v1/collection/:collection_slug/stats", restgo.ObjectParams(req)...)
	if err != nil {
		return nil, err
	}
	var response CollectionStatsResponse
	err = ParseRsp(rsp, &response)
	if err != nil {
		return nil, WrapperRspError(rsp)
	}
	return response.Stats, nil
}

type CollectionStatsResponse struct {
	Stats *model.CollectionStats `opensea:"stats"`
}
