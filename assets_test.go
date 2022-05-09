package opensea

import "testing"

func TestClient_Assets(t *testing.T) {
	var cli = newClient()
	var rsp, err = cli.Assets(ctx, &AssetsRequest{
		OrderDirection: "desc",
		Offset:         0,
		Limit:          10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range rsp.Assets {
		t.Log(c)
	}
	if rsp.Next == "" {
		return
	}
	rsp, err = cli.Assets(ctx, &AssetsRequest{
		OrderDirection: "desc",
		Offset:         0,
		Limit:          10,
		Cursor:         rsp.Next,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range rsp.Assets {
		t.Log(c)
	}
}

func TestClient_Assets_IncludeOrders(t *testing.T) {
	var cli = newClient()
	var rsp, err = cli.Assets(ctx, &AssetsRequest{
		OrderDirection: "desc",
		Offset:         0,
		Limit:          10,
		IncludeOrders:  true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range rsp.Assets {
		t.Log(c)
	}
}
