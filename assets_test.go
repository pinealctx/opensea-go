package opensea

import "testing"

func TestClient_Assets(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Assets(ctx, &AssetsRequest{
		OrderDirection: "desc",
		Offset:         0,
		Limit:          10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}

func TestClient_Assets_IncludeOrders(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Assets(ctx, &AssetsRequest{
		OrderDirection: "desc",
		Offset:         0,
		Limit:          10,
		IncludeOrders:  true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}
