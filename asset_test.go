package opensea

import "testing"

func TestClient_Asset(t *testing.T) {
	// 0x5206e78b21ce315ce284fb24cf05e0585a93b1d9
	// 0xf368deaf4d8f3d7b9aa2c9c991a652c6c0670879fffe86b2898fda2b791ee37f
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
