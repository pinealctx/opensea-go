package opensea

import "testing"

func TestClient_Events(t *testing.T) {
	var cli = newClient()
	var list, err = cli.Events(ctx, &EventsRequest{
		Offset: 0,
		Limit:  10,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range list {
		t.Log(c)
	}
}
