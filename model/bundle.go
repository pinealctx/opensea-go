package model

// Bundle Bundles are groups of items for sale on OpenSea.
// You can buy them all at once in one transaction, and you can create them without any transactions or gas,
// as long as you've already approved the assets inside.
type Bundle struct {
	Maker         *Account     `opensea:"maker" json:"maker"`
	Slug          string       `opensea:"slug" json:"slug"`
	Assets        []*Asset     `opensea:"assets" json:"assets"`
	Name          string       `opensea:"name" json:"name"`
	Description   string       `opensea:"description" json:"description"`
	ExternalLink  *string      `opensea:"external_link" json:"externalLink"`
	AssetContract *Contract    `opensea:"asset_contract" json:"assetContract"`
	Permalink     string       `opensea:"permalink" json:"permalink"`
	SellOrders    []*SellOrder `opensea:"sell_orders" json:"sellOrders"`
}

type SellOrder struct {
	CreatedDate       string `opensea:"created_date" json:"createdDate"`
	ClosingDate       string `opensea:"closing_date" json:"closingDate"`
	ClosingExtendable bool   `opensea:"closing_extendable" json:"closingExtendable"`
	ExpirationTime    int    `opensea:"expiration_time" json:"expirationTime"`
	ListingTime       int    `opensea:"listing_time" json:"listingTime"`
	OrderHash         string `opensea:"order_hash" json:"orderHash"`
	Metadata          struct {
		Bundle struct {
			Assets      []*AssetAddress `opensea:"assets" json:"assets"`
			Schemas     []string        `opensea:"schemas" json:"schemas"`
			Name        string          `opensea:"name" json:"name"`
			Description string          `opensea:"description" json:"description"`
		} `opensea:"bundle" json:"bundle"`
	} `opensea:"metadata" json:"metadata"`
	Exchange             string                `opensea:"exchange" json:"exchange"`
	Maker                *BundleAccount        `opensea:"maker" json:"maker"`
	Taker                *BundleAccount        `opensea:"taker" json:"taker"`
	CurrentPrice         string                `opensea:"current_price" json:"currentPrice"`
	CurrentBounty        string                `opensea:"current_bounty" json:"currentBounty"`
	BountyMultiple       string                `opensea:"bounty_multiple" json:"bountyMultiple"`
	MakerRelayerFee      string                `opensea:"maker_relayer_fee" json:"makerRelayerFee"`
	TakerRelayerFee      string                `opensea:"taker_relayer_fee" json:"takerRelayerFee"`
	MakerProtocolFee     string                `opensea:"maker_protocol_fee" json:"makerProtocolFee"`
	TakerProtocolFee     string                `opensea:"taker_protocol_fee" json:"takerProtocolFee"`
	MakerReferrerFee     string                `opensea:"maker_referrer_fee" json:"makerReferrerFee"`
	FeeRecipient         *BundleAccount        `opensea:"fee_recipient" json:"feeRecipient"`
	FeeMethod            int                   `opensea:"fee_method" json:"feeMethod"`
	Side                 int                   `opensea:"side" json:"side"`
	SaleKind             int                   `opensea:"sale_kind" json:"saleKind"`
	Target               string                `opensea:"target" json:"target"`
	HowToCall            int                   `opensea:"how_to_call" json:"howToCall"`
	Calldata             string                `opensea:"calldata" json:"calldata"`
	ReplacementPattern   string                `opensea:"replacement_pattern" json:"replacementPattern"`
	StaticTarget         string                `opensea:"static_target" json:"staticTarget"`
	StaticExtradata      string                `opensea:"static_extradata" json:"staticExtradata"`
	PaymentToken         string                `opensea:"payment_token" json:"paymentToken"`
	PaymentTokenContract *PaymentTokenContract `opensea:"payment_token_contract" json:"paymentTokenContract"`
	BasePrice            string                `opensea:"base_price" json:"basePrice"`
	Extra                string                `opensea:"extra" json:"extra"`
	Quantity             string                `opensea:"quantity" json:"quantity"`
	Salt                 string                `opensea:"salt" json:"salt"`
	V                    int                   `opensea:"v" json:"v"`
	R                    string                `opensea:"r" json:"r"`
	S                    string                `opensea:"s" json:"s"`
	ApprovedOnChain      bool                  `opensea:"approved_on_chain" json:"approvedOnChain"`
	Cancelled            bool                  `opensea:"cancelled" json:"cancelled"`
	Finalized            bool                  `opensea:"finalized" json:"finalized"`
	MarkedInvalid        bool                  `opensea:"marked_invalid" json:"markedInvalid"`
	PrefixedHash         string                `opensea:"prefixed_hash" json:"prefixedHash"`
}

type BundleAccount struct {
	Account
	User int32 `opensea:"user" json:"user"`
}
