package opensea

import "testing"

func TestClient_Contract(t *testing.T) {
	var cli = newClient()
	var contract, err = cli.Contract(ctx, &ContractRequest{
		AssetContractAddress: "0xb47e3cd837ddf8e4c57f05d70ab865de6e193bbb",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(contract)
}
