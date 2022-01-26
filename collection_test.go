package opensea

import "testing"

func TestClient_Collection(t *testing.T) {
	var cli = newClient()
	var collection, err = cli.Collection(ctx, &CollectionRequest{
		CollectionSlug: "doodles-official",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(collection)
}
