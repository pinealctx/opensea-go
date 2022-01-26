package model

type Side int32

const (
	SideBuy Side = iota
	SideSell
)

// SaleKind the kind of sell order
// NOTE: use only_english=true for filtering for only English Auctions
type SaleKind int32

const (
	// SaleKindFixedOrMinBidPrice fixed-price sales or min-bid auctions
	SaleKindFixedOrMinBidPrice SaleKind = iota
	// SaleKindDecliningPrice declining-price Dutch Auctions
	SaleKindDecliningPrice
)

type Order struct {
	CreatedDate          string                `json:"created_date"`
	ClosingDate          string                `json:"closing_date"`
	ClosingExtendable    bool                  `json:"closing_extendable"`
	ExpirationTime       int                   `json:"expiration_time"`
	ListingTime          int                   `json:"listing_time"`
	OrderHash            string                `json:"order_hash"`
	Metadata             *Metadata             `json:"metadata"`
	Exchange             string                `json:"exchange"`
	Maker                *Account              `json:"maker"`
	Taker                *Account              `json:"taker"`
	CurrentPrice         string                `json:"current_price"`
	CurrentBounty        string                `json:"current_bounty"`
	BountyMultiple       string                `json:"bounty_multiple"`
	MakerRelayerFee      string                `json:"maker_relayer_fee"`
	TakerRelayerFee      string                `json:"taker_relayer_fee"`
	MakerProtocolFee     string                `json:"maker_protocol_fee"`
	TakerProtocolFee     string                `json:"taker_protocol_fee"`
	MakerReferrerFee     string                `json:"maker_referrer_fee"`
	FeeRecipient         *Account              `json:"fee_recipient"`
	FeeMethod            int                   `json:"fee_method"`
	Side                 Side                  `json:"side"`
	SaleKind             SaleKind              `json:"sale_kind"`
	Target               string                `json:"target"`
	HowToCall            int                   `json:"how_to_call"`
	Calldata             string                `json:"calldata"`
	ReplacementPattern   string                `json:"replacement_pattern"`
	StaticTarget         string                `json:"static_target"`
	StaticExtradata      string                `json:"static_extradata"`
	PaymentToken         string                `json:"payment_token"`
	PaymentTokenContract *PaymentTokenContract `json:"payment_token_contract"`
	BasePrice            string                `json:"base_price"`
	Extra                string                `json:"extra"`
	Quantity             string                `json:"quantity"`
	Salt                 string                `json:"salt"`
	V                    *int                  `json:"v"`
	R                    *string               `json:"r"`
	S                    *string               `json:"s"`
	ApprovedOnChain      bool                  `json:"approved_on_chain"`
	Cancelled            bool                  `json:"cancelled"`
	Finalized            bool                  `json:"finalized"`
	MarkedInvalid        bool                  `json:"marked_invalid"`
	PrefixedHash         string                `json:"prefixed_hash"`

	ID          int64   `json:"id,omitempty"`
	Asset       *Asset  `json:"asset"`
	AssetBundle *Bundle `json:"asset_bundle"`
}

type Metadata struct {
	Asset struct {
		ID      string `json:"id"`
		Address string `json:"address"`
	} `json:"asset"`
	Schema          string `json:"schema"`
	ReferrerAddress string `json:"referrerAddress"`
}

type PaymentTokenContract struct {
	ID       int    `json:"id"`
	Symbol   string `json:"symbol"`
	Address  string `json:"address"`
	ImageURL string `json:"image_url"`
	Name     string `json:"name"`
	Decimals int    `json:"decimals"`
	EthPrice string `json:"eth_price"`
	UsdPrice string `json:"usd_price"`
}
