package model

// Bundle Bundles are groups of items for sale on OpenSea.
// You can buy them all at once in one transaction, and you can create them without any transactions or gas,
// as long as you've already approved the assets inside.
type Bundle struct {
	Maker         *Account     `json:"maker"`
	Slug          string       `json:"slug"`
	Assets        []*Asset     `json:"assets"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	ExternalLink  *string      `json:"external_link"`
	AssetContract *Contract    `json:"asset_contract"`
	Permalink     string       `json:"permalink"`
	SellOrders    []*SellOrder `json:"sell_orders"`
}

type SellOrder struct {
	CreatedDate       string `json:"created_date"`
	ClosingDate       string `json:"closing_date"`
	ClosingExtendable bool   `json:"closing_extendable"`
	ExpirationTime    int    `json:"expiration_time"`
	ListingTime       int    `json:"listing_time"`
	OrderHash         string `json:"order_hash"`
	Metadata          struct {
		Bundle struct {
			Assets []struct {
				ID      string `json:"id"`
				Address string `json:"address"`
			} `json:"assets"`
			Schemas     []string `json:"schemas"`
			Name        string   `json:"name"`
			Description string   `json:"description"`
		} `json:"bundle"`
	} `json:"metadata"`
	Exchange             string                `json:"exchange"`
	Maker                *BundleAccount        `json:"maker"`
	Taker                *BundleAccount        `json:"taker"`
	CurrentPrice         string                `json:"current_price"`
	CurrentBounty        string                `json:"current_bounty"`
	BountyMultiple       string                `json:"bounty_multiple"`
	MakerRelayerFee      string                `json:"maker_relayer_fee"`
	TakerRelayerFee      string                `json:"taker_relayer_fee"`
	MakerProtocolFee     string                `json:"maker_protocol_fee"`
	TakerProtocolFee     string                `json:"taker_protocol_fee"`
	MakerReferrerFee     string                `json:"maker_referrer_fee"`
	FeeRecipient         *BundleAccount        `json:"fee_recipient"`
	FeeMethod            int                   `json:"fee_method"`
	Side                 int                   `json:"side"`
	SaleKind             int                   `json:"sale_kind"`
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
	V                    int                   `json:"v"`
	R                    string                `json:"r"`
	S                    string                `json:"s"`
	ApprovedOnChain      bool                  `json:"approved_on_chain"`
	Cancelled            bool                  `json:"cancelled"`
	Finalized            bool                  `json:"finalized"`
	MarkedInvalid        bool                  `json:"marked_invalid"`
	PrefixedHash         string                `json:"prefixed_hash"`
}

type BundleAccount struct {
	Account
	User int32 `json:"user"`
}
