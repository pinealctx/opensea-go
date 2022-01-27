package opensea

import "testing"

func TestClient_Orders(t *testing.T) {
	var cli = newClient()
	var response, err = cli.Orders(ctx, &OrderRequest{
		AssetContractAddress: "0x622062a3d249cccee76864a148c89484a23a6144",
		TokenID:              "2",
		Limit:                20,
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
