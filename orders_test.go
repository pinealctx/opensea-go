package opensea

import "testing"

func TestClient_Orders(t *testing.T) {
	var cli = newClient()
	var response, err = cli.Orders(ctx, &OrdersRequest{
		AssetContractAddress: "0xed5af388653567af2f388e6224dc7c4b3241c544",
		TokenID:              "671",
		Limit:                20,
		Side:                 1,
		Offset:               0,
		OrderBy:              "created_date",
		OrderDirection:       "desc",
		Bundled:              false,
		IncludeBundled:       false,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("Count", response.Count)
	for _, item := range response.Orders {
		t.Log(item)
	}
}
