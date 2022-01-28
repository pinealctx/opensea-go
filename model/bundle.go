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
	Order
	Metadata     *SellOrderMetadata `opensea:"metadata" json:"metadata"`
	Maker        *BundleAccount     `opensea:"maker" json:"maker"`
	Taker        *BundleAccount     `opensea:"taker" json:"taker"`
	FeeRecipient *BundleAccount     `opensea:"fee_recipient" json:"feeRecipient"`
}

type BundleAccount struct {
	Account
	User int32 `opensea:"user" json:"user"`
}

type SellOrderMetadata struct {
	Bundle *SellOrderMetadataBundle `opensea:"bundle" json:"bundle"`
}

type SellOrderMetadataBundle struct {
	Assets      []*AssetAddress `opensea:"assets" json:"assets"`
	Schemas     []string        `opensea:"schemas" json:"schemas"`
	Name        string          `opensea:"name" json:"name"`
	Description string          `opensea:"description" json:"description"`
}
