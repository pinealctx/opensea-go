package model

// Bundle Bundles are groups of items for sale on OpenSea.
// You can buy them all at once in one transaction, and you can create them without any transactions or gas,
// as long as you've already approved the assets inside.
type Bundle struct {
	Maker         *Account     `json:"maker" camel:"maker"`
	Slug          string       `json:"slug" camel:"slug"`
	Assets        []*Asset     `json:"assets" camel:"assets"`
	Name          string       `json:"name" camel:"name"`
	Description   string       `json:"description" camel:"description"`
	ExternalLink  *string      `json:"external_link" camel:"externalLink"`
	AssetContract *Contract    `json:"asset_contract" camel:"assetContract"`
	Permalink     string       `json:"permalink" camel:"permalink"`
	SellOrders    []*SellOrder `json:"sell_orders" camel:"sellOrders"`
}

type SellOrder struct {
	CreatedDate       string `json:"created_date" camel:"createdDate"`
	ClosingDate       string `json:"closing_date" camel:"closingDate"`
	ClosingExtendable bool   `json:"closing_extendable" camel:"closingExtendable"`
	ExpirationTime    int    `json:"expiration_time" camel:"expirationTime"`
	ListingTime       int    `json:"listing_time" camel:"listingTime"`
	OrderHash         string `json:"order_hash" camel:"orderHash"`
	Metadata          struct {
		Bundle struct {
			Assets      []*AssetAddress `json:"assets" camel:"assets"`
			Schemas     []string       `json:"schemas" camel:"schemas"`
			Name        string         `json:"name" camel:"name"`
			Description string         `json:"description" camel:"description"`
		} `json:"bundle" camel:"bundle"`
	} `json:"metadata" camel:"metadata"`
	Exchange             string                `json:"exchange" camel:"exchange"`
	Maker                *BundleAccount        `json:"maker" camel:"maker"`
	Taker                *BundleAccount        `json:"taker" camel:"taker"`
	CurrentPrice         string                `json:"current_price" camel:"currentPrice"`
	CurrentBounty        string                `json:"current_bounty" camel:"currentBounty"`
	BountyMultiple       string                `json:"bounty_multiple" camel:"bountyMultiple"`
	MakerRelayerFee      string                `json:"maker_relayer_fee" camel:"makerRelayerFee"`
	TakerRelayerFee      string                `json:"taker_relayer_fee" camel:"takerRelayerFee"`
	MakerProtocolFee     string                `json:"maker_protocol_fee" camel:"makerProtocolFee"`
	TakerProtocolFee     string                `json:"taker_protocol_fee" camel:"takerProtocolFee"`
	MakerReferrerFee     string                `json:"maker_referrer_fee" camel:"makerReferrerFee"`
	FeeRecipient         *BundleAccount        `json:"fee_recipient" camel:"feeRecipient"`
	FeeMethod            int                   `json:"fee_method" camel:"feeMethod"`
	Side                 int                   `json:"side" camel:"side"`
	SaleKind             int                   `json:"sale_kind" camel:"saleKind"`
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
	V                    int                   `json:"v" camel:"v"`
	R                    string                `json:"r" camel:"r"`
	S                    string                `json:"s" camel:"s"`
	ApprovedOnChain      bool                  `json:"approved_on_chain" camel:"approvedOnChain"`
	Cancelled            bool                  `json:"cancelled" camel:"cancelled"`
	Finalized            bool                  `json:"finalized" camel:"finalized"`
	MarkedInvalid        bool                  `json:"marked_invalid" camel:"markedInvalid"`
	PrefixedHash         string                `json:"prefixed_hash" camel:"prefixedHash"`
}

type BundleAccount struct {
	Account
	User int32 `json:"user" camel:"user"`
}
