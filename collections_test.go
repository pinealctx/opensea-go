package opensea

import "testing"

func TestClient_Collections(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Collections(ctx, &CollectionsRequest{
		AssetOwner: "",
		Offset:     0,
		Limit:      10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}
