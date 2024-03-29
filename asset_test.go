package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0xed5af388653567af2f388e6224dc7c4b3241c544",
		TokenID:              "671",
		AccountAddress:       "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, o := range asset.Orders {
		t.Log(o)
	}
}
