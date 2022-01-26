package model

import "github.com/shopspring/decimal"

// Collection Collections are used to represent all the assets in a single (or multiple) contract addresses
// and help users group items from the same creator. They have one or more owners and are typically
// associated with important metadata such as creator royalties and descriptions.
// Visit it(https://docs.opensea.io/reference/collection-model) to learn anymore.
type Collection struct {
	// The collection name. Typically derived from the first contract imported to the collection but can be changed by the user
	Name string `json:"name"`
	// Description for the model
	Description      string  `json:"description"`
	ShortDescription *string `json:"short_description"`
	// The collection slug that is used to link to the collection on OpenSea.
	// This value can change by the owner but must be unique across all collection slugs in OpenSea
	Slug string `json:"slug"`
	// External link to the original website for the collection
	ExternalURL string `json:"external_url"`
	// An image for the collection. Note that this is the cached URL we store on our end.
	// The original image url is image_original_url
	ImageURL      string  `json:"image_url"`
	LargeImageURL *string `json:"large_image_url"`
	// Approved editors on this collection.
	Editors []string `json:"editors"`
	// The payment tokens accepted for this collection
	PaymentTokens []*PaymentToken `json:"payment_tokens"`
	// A list of the contracts that are associated with this collection
	PrimaryAssetContracts []*PrimaryAssetContract `json:"primary_asset_contracts"`
	// A dictionary listing all the trait types available within this collection
	Traits map[string]map[string]int32 `json:"traits"`
	// A dictionary containing some sales statistics related to this collection, including trade volume and floor prices
	Stats *CollectionStats `json:"stats"`
	// Image used in the horizontal top banner for the collection.
	BannerImageURL string `json:"banner_image_url"`
	// The payout address for the collection's royalties
	PayoutAddress string `json:"payout_address"`
	// The collector's fees that get paid out to them when sales are made for their collections
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points"`
	// The collection's approval status within OpenSea.
	// Can be not_requested (brand new collections), requested (collections that requested safelisting on our site),
	// approved (collections that are approved on our site and can be found in search results),
	// and verified (verified collections)
	SafelistRequestStatus string `json:"safelist_request_status"`

	CreatedDate                 string      `json:"created_date"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points"`
	DisplayData                 DisplayData `json:"display_data"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            *string     `json:"featured_image_url"`
	Hidden                      bool        `json:"hidden"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
	RequireEmail                bool        `json:"require_email"`

	TwitterUsername   *string `json:"twitter_username"`
	InstagramUsername string  `json:"instagram_username"`
	MediumUsername    *string `json:"medium_username"`

	TelegramURL *string `json:"telegram_url"`
	DiscordURL  string  `json:"discord_url"`
	ChatURL     *string `json:"chat_url"`
	WikiURL     *string `json:"wiki_url"`
}

type PaymentToken struct {
	ID       int              `json:"id"`
	Symbol   string           `json:"symbol"`
	Address  string           `json:"address"`
	ImageURL string           `json:"image_url"`
	Name     string           `json:"name"`
	Decimals int              `json:"decimals"`
	EthPrice *decimal.Decimal `json:"eth_price"`
	UsdPrice *decimal.Decimal `json:"usd_price"`
}

type PrimaryAssetContract struct {
	Address                     string  `json:"address"`
	AssetContractType           string  `json:"asset_contract_type"`
	CreatedDate                 string  `json:"created_date"`
	Name                        string  `json:"name"`
	NftVersion                  string  `json:"nft_version"`
	OpenseaVersion              *string `json:"opensea_version"`
	Owner                       int     `json:"owner"`
	SchemaName                  string  `json:"schema_name"`
	Symbol                      string  `json:"symbol"`
	TotalSupply                 string  `json:"total_supply"`
	Description                 string  `json:"description"`
	ExternalLink                string  `json:"external_link"`
	ImageURL                    string  `json:"image_url"`
	DefaultToFiat               bool    `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int     `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int     `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool    `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int     `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int     `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int     `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int     `json:"seller_fee_basis_points"`
	PayoutAddress               string  `json:"payout_address"`
}

type CollectionStats struct {
	OneDayVolume          float64 `json:"one_day_volume"`
	OneDayChange          float64 `json:"one_day_change"`
	OneDaySales           float64 `json:"one_day_sales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price"`
	SevenDayVolume        float64 `json:"seven_day_volume"`
	SevenDayChange        float64 `json:"seven_day_change"`
	SevenDaySales         float64 `json:"seven_day_sales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume"`
	ThirtyDayChange       float64 `json:"thirty_day_change"`
	ThirtyDaySales        float64 `json:"thirty_day_sales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price"`
	TotalVolume           float64 `json:"total_volume"`
	TotalSales            float64 `json:"total_sales"`
	TotalSupply           float64 `json:"total_supply"`
	Count                 float64 `json:"count"`
	NumOwners             int     `json:"num_owners"`
	AveragePrice          float64 `json:"average_price"`
	NumReports            int     `json:"num_reports"`
	MarketCap             float64 `json:"market_cap"`
	FloorPrice            float64 `json:"floor_price"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}
