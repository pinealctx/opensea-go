package opensea

import "testing"

func TestClient_Bundles(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Bundles(ctx, &BundlesRequest{
		OnSale: true,
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}
