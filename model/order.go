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
	CreatedDate          string                `json:"created_date" camel:"createdDate"`
	ClosingDate          string                `json:"closing_date" camel:"closingDate"`
	ClosingExtendable    bool                  `json:"closing_extendable" camel:"closingExtendable"`
	ExpirationTime       int                   `json:"expiration_time" camel:"expirationTime"`
	ListingTime          int                   `json:"listing_time" camel:"listingTime"`
	OrderHash            string                `json:"order_hash" camel:"orderHash"`
	Metadata             *Metadata             `json:"metadata" camel:"metadata"`
	Exchange             string                `json:"exchange" camel:"exchange"`
	Maker                *Account              `json:"maker" camel:"maker"`
	Taker                *Account              `json:"taker" camel:"taker"`
	CurrentPrice         string                `json:"current_price" camel:"currentPrice"`
	CurrentBounty        string                `json:"current_bounty" camel:"currentBounty"`
	BountyMultiple       string                `json:"bounty_multiple" camel:"bountyMultiple"`
	MakerRelayerFee      string                `json:"maker_relayer_fee" camel:"makerRelayerFee"`
	TakerRelayerFee      string                `json:"taker_relayer_fee" camel:"takerRelayerFee"`
	MakerProtocolFee     string                `json:"maker_protocol_fee" camel:"makerProtocolFee"`
	TakerProtocolFee     string                `json:"taker_protocol_fee" camel:"takerProtocolFee"`
	MakerReferrerFee     string                `json:"maker_referrer_fee" camel:"makerReferrerFee"`
	FeeRecipient         *Account              `json:"fee_recipient" camel:"feeRecipient"`
	FeeMethod            int                   `json:"fee_method" camel:"feeMethod"`
	Side                 Side                  `json:"side" camel:"side"`
	SaleKind             SaleKind              `json:"sale_kind" camel:"saleKind"`
	Target               string                `json:"target" camel:"target"`
	HowToCall            int                   `json:"how_to_call" camel:"howToCall"`
	Calldata             string                `json:"calldata" camel:"calldata"`
	ReplacementPattern   string                `json:"replacement_pattern" camel:"replacementPattern"`
	StaticTarget         string                `json:"static_target" camel:"staticTarget"`
	StaticExtradata      string                `json:"static_extradata" camel:"staticExtradata"`
	PaymentToken         string                `json:"payment_token" camel:"paymentToken"`
	PaymentTokenContract *PaymentTokenContract `json:"payment_token_contract" camel:"paymentTokenContract"`
	BasePrice            string                `json:"base_price" camel:"basePrice"`
	Extra                string                `json:"extra" camel:"extra"`
	Quantity             string                `json:"quantity" camel:"quantity"`
	Salt                 string                `json:"salt" camel:"salt"`
	V                    *int                  `json:"v" camel:"v"`
	R                    *string               `json:"r" camel:"r"`
	S                    *string               `json:"s" camel:"s"`
	ApprovedOnChain      bool                  `json:"approved_on_chain" camel:"approvedOnChain"`
	Cancelled            bool                  `json:"cancelled" camel:"cancelled"`
	Finalized            bool                  `json:"finalized" camel:"finalized"`
	MarkedInvalid        bool                  `json:"marked_invalid" camel:"markedInvalid"`
	PrefixedHash         string                `json:"prefixed_hash" camel:"prefixedHash"`

	ID          int64   `json:"id,omitempty" camel:"id"`
	Asset       *Asset  `json:"asset" camel:"asset"`
	AssetBundle *Bundle `json:"asset_bundle" camel:"assetBundle"`
}

type Metadata struct {
	Asset           *AssetAddress `json:"asset" camel:"asset"`
	Schema          string       `json:"schema" camel:"schema"`
	ReferrerAddress string       `json:"referrerAddress" camel:"referrerAddress"`
}

type PaymentTokenContract struct {
	ID       int    `json:"id" camel:"id"`
	Symbol   string `json:"symbol" camel:"symbol"`
	Address  string `json:"address" camel:"address"`
	ImageURL string `json:"image_url" camel:"imageURL"`
	Name     string `json:"name" camel:"name"`
	Decimals int    `json:"decimals" camel:"decimals"`
	EthPrice string `json:"eth_price" camel:"ethPrice"`
	UsdPrice string `json:"usd_price" camel:"usdPrice"`
}
