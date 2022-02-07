package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0xe0916a45ec7c64a0e4365cd546a3d8534d362e05",
		TokenID:              "2",
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
