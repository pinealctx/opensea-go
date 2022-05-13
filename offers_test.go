package opensea

import "testing"

func TestClient_Offers(t *testing.T) {
	var cli = newClient()
	var response, err = cli.Offers(ctx, &OffersRequest{
		AssetContractAddress: "0xed5af388653567af2f388e6224dc7c4b3241c544",
		TokenID:              "671",
		Limit:                20,
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("rsp", response)
}
