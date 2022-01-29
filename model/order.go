package model

import "github.com/shopspring/decimal"

type Side uint8

const (
	SideBuy Side = iota
	SideSell
)

// SaleKind the kind of sell order
// use only_english=true for filtering for only English Auctions
type SaleKind uint8

const (
	// SaleKindFixedOrMinBidPrice fixed-price sales or min-bid auctions
	SaleKindFixedOrMinBidPrice SaleKind = iota
	// SaleKindDecliningPrice declining-price Dutch Auctions
	SaleKindDecliningPrice
)

type HowToCall uint8

const (
	Call HowToCall = iota
	DelegateCall
)

type FeeMethod uint8

const (
	ProtocolFee FeeMethod = iota
	SplitFee
)

type Order struct {
	CreatedDate          string           `opensea:"created_date" json:"createdDate"`
	ClosingDate          string           `opensea:"closing_date" json:"closingDate"`
	ClosingExtendable    bool             `opensea:"closing_extendable" json:"closingExtendable"`
	ExpirationTime       int              `opensea:"expiration_time" json:"expirationTime"`
	ListingTime          int              `opensea:"listing_time" json:"listingTime"`
	OrderHash            string           `opensea:"order_hash" json:"orderHash"`
	Metadata             *Metadata        `opensea:"metadata" json:"metadata"`
	Exchange             string           `opensea:"exchange" json:"exchange"`
	Maker                *Account         `opensea:"maker" json:"maker"`
	Taker                *Account         `opensea:"taker" json:"taker"`
	CurrentPrice         *decimal.Decimal `opensea:"current_price" json:"currentPrice"`
	CurrentBounty        string           `opensea:"current_bounty" json:"currentBounty"`
	BountyMultiple       string           `opensea:"bounty_multiple" json:"bountyMultiple"`
	MakerRelayerFee      string           `opensea:"maker_relayer_fee" json:"makerRelayerFee"`
	TakerRelayerFee      string           `opensea:"taker_relayer_fee" json:"takerRelayerFee"`
	MakerProtocolFee     string           `opensea:"maker_protocol_fee" json:"makerProtocolFee"`
	TakerProtocolFee     string           `opensea:"taker_protocol_fee" json:"takerProtocolFee"`
	MakerReferrerFee     string           `opensea:"maker_referrer_fee" json:"makerReferrerFee"`
	FeeRecipient         *Account         `opensea:"fee_recipient" json:"feeRecipient"`
	FeeMethod            FeeMethod        `opensea:"fee_method" json:"feeMethod"`
	Side                 Side             `opensea:"side" json:"side"`
	SaleKind             SaleKind         `opensea:"sale_kind" json:"saleKind"`
	Target               string           `opensea:"target" json:"target"`
	HowToCall            HowToCall        `opensea:"how_to_call" json:"howToCall"`
	Calldata             string           `opensea:"calldata" json:"calldata"`
	ReplacementPattern   string           `opensea:"replacement_pattern" json:"replacementPattern"`
	StaticTarget         string           `opensea:"static_target" json:"staticTarget"`
	StaticExtradata      string           `opensea:"static_extradata" json:"staticExtradata"`
	PaymentToken         string           `opensea:"payment_token" json:"paymentToken"`
	PaymentTokenContract *PaymentToken    `opensea:"payment_token_contract" json:"paymentTokenContract"`
	BasePrice            *decimal.Decimal `opensea:"base_price" json:"basePrice"`
	Extra                string           `opensea:"extra" json:"extra"`
	Quantity             string           `opensea:"quantity" json:"quantity"`
	Salt                 string           `opensea:"salt" json:"salt"`
	V                    *int             `opensea:"v" json:"v"`
	R                    *string          `opensea:"r" json:"r"`
	S                    *string          `opensea:"s" json:"s"`
	ApprovedOnChain      bool             `opensea:"approved_on_chain" json:"approvedOnChain"`
	Cancelled            bool             `opensea:"cancelled" json:"cancelled"`
	Finalized            bool             `opensea:"finalized" json:"finalized"`
	MarkedInvalid        bool             `opensea:"marked_invalid" json:"markedInvalid"`
	PrefixedHash         string           `opensea:"prefixed_hash" json:"prefixedHash"`

	ID          int64   `opensea:"id,omitempty" json:"id"`
	Asset       *Asset  `opensea:"asset" json:"asset"`
	AssetBundle *Bundle `opensea:"asset_bundle" json:"assetBundle"`
}

type Metadata struct {
	Asset           *AssetAddress `opensea:"asset" json:"asset"`
	Schema          string        `opensea:"schema" json:"schema"`
	ReferrerAddress string        `opensea:"referrerAddress" json:"referrerAddress"`
}
