package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0x66583bd73a27c9245b723ff6a58f872234c3a50a",
		TokenID:              "3",
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
