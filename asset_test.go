package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb",
		TokenID:              "1",
		AccountAddress:       "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(asset)
}
