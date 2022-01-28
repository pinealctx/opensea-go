package opensea

import "testing"

func TestClient_Contract(t *testing.T) {
	var cli = newClient()
	var contract, err = cli.Contract(ctx, &ContractRequest{
		AssetContractAddress: "0x0000000000000000000000000000000000000000",
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(contract)
}
