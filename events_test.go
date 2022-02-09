package opensea

import "testing"

func TestClient_Events(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Events(ctx, &EventsRequest{
		AssetContractAddress: "0x88b48f654c30e99bc2e4a1559b4dcf1ad93fa656",
		TokenID:              "81808355449469934498948910091380706836086362421539041654606043397932313477121",
		Offset:               0,
		Limit:                10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}
