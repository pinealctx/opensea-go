package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	var cli = newClient()
	var asset, err = cli.Asset(ctx, &AssetRequest{
		AssetContractAddress: "0x88b48f654c30e99bc2e4a1559b4dcf1ad93fa656",
		TokenID:              "81808355449469934498948910091380706836086362421539041654606043397932313477121",
		AccountAddress:       "",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(asset)
}
