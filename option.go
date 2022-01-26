package opensea

const (
	hostProd = "https://api.opensea.io"
	hostTest = "https://testnets-api.opensea.io"
)

var defaultOption = &option{baseURL: hostProd}

type option struct {
	baseURL string
	apiKey  string
}

type OptionFn func(*option)

func WithTestNet() OptionFn {
	return func(o *option) {
		o.baseURL = hostTest
	}
}

func WithProdNet() OptionFn {
	return func(o *option) {
		o.baseURL = hostProd
	}
}

func WithAPIKey(apiKey string) OptionFn {
	return func(o *option) {
		o.apiKey = apiKey
	}
}
