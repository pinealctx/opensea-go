package model

import "github.com/shopspring/decimal"

// Collection Collections are used to represent all the assets in a single (or multiple) contract addresses
// and help users group items from the same creator. They have one or more owners and are typically
// associated with important metadata such as creator royalties and descriptions.
// Visit it(https://docs.opensea.io/reference/collection-model) to learn anymore.
type Collection struct {
	// The collection name. Typically derived from the first contract imported to the collection but can be changed by the user
	Name string `json:"name" camel:"name"`
	// Description for the model
	Description      string  `json:"description" camel:"description"`
	ShortDescription *string `json:"short_description" camel:"shortDescription"`
	// The collection slug that is used to link to the collection on OpenSea.
	// This value can change by the owner but must be unique across all collection slugs in OpenSea
	Slug string `json:"slug" camel:"slug"`
	// External link to the original website for the collection
	ExternalURL string `json:"external_url" camel:"externalURL"`
	// An image for the collection. Note that this is the cached URL we store on our end.
	// The original image url is image_original_url
	ImageURL      string  `json:"image_url" camel:"imageURL"`
	LargeImageURL *string `json:"large_image_url" camel:"largeImageURL"`
	// Approved editors on this collection.
	Editors []string `json:"editors" camel:"editors"`
	// The payment tokens accepted for this collection
	PaymentTokens []*PaymentToken `json:"payment_tokens" camel:"paymentTokens"`
	// A list of the contracts that are associated with this collection
	PrimaryAssetContracts []*PrimaryAssetContract `json:"primary_asset_contracts" camel:"primaryAssetContracts"`
	// A dictionary listing all the trait types available within this collection
	Traits map[string]map[string]int32 `json:"traits" camel:"traits"`
	// A dictionary containing some sales statistics related to this collection, including trade volume and floor prices
	Stats *CollectionStats `json:"stats" camel:"stats"`
	// Image used in the horizontal top banner for the collection.
	BannerImageURL string `json:"banner_image_url" camel:"bannerImageURL"`
	// The payout address for the collection's royalties
	PayoutAddress string `json:"payout_address" camel:"payoutAddress"`
	// The collector's fees that get paid out to them when sales are made for their collections
	DevSellerFeeBasisPoints string `json:"dev_seller_fee_basis_points" camel:"devSellerFeeBasisPoints"`
	// The collection's approval status within OpenSea.
	// Can be not_requested (brand new collections), requested (collections that requested safelisting on our site),
	// approved (collections that are approved on our site and can be found in search results),
	// and verified (verified collections)
	SafelistRequestStatus string `json:"safelist_request_status" camel:"safelistRequestStatus"`

	CreatedDate                 string      `json:"created_date" camel:"createdDate"`
	DefaultToFiat               bool        `json:"default_to_fiat" camel:"defaultToFiat"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points" camel:"devBuyerFeeBasisPoints"`
	DisplayData                 DisplayData `json:"display_data" camel:"displayData"`
	Featured                    bool        `json:"featured" camel:"featured"`
	FeaturedImageURL            *string     `json:"featured_image_url" camel:"featuredImageURL"`
	Hidden                      bool        `json:"hidden" camel:"hidden"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist" camel:"isSubjectToWhitelist"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" camel:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points" camel:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points" camel:"openseaSellerFeeBasisPoints"`
	RequireEmail                bool        `json:"require_email" camel:"requireEmail"`

	TwitterUsername   *string `json:"twitter_username" camel:"twitterUsername"`
	InstagramUsername string  `json:"instagram_username" camel:"instagramUsername"`
	MediumUsername    *string `json:"medium_username" camel:"mediumUsername"`

	TelegramURL *string `json:"telegram_url" camel:"telegramURL"`
	DiscordURL  string  `json:"discord_url" camel:"discordURL"`
	ChatURL     *string `json:"chat_url" camel:"chatURL"`
	WikiURL     *string `json:"wiki_url" camel:"wikiURL"`
}

type PaymentToken struct {
	ID       int              `json:"id" camel:"id"`
	Symbol   string           `json:"symbol" camel:"symbol"`
	Address  string           `json:"address" camel:"address"`
	ImageURL string           `json:"image_url" camel:"imageURL"`
	Name     string           `json:"name" camel:"name"`
	Decimals int              `json:"decimals" camel:"decimals"`
	EthPrice *decimal.Decimal `json:"eth_price" camel:"ethPrice"`
	UsdPrice *decimal.Decimal `json:"usd_price" camel:"usdPrice"`
}

type PrimaryAssetContract struct {
	Address                     string  `json:"address" camel:"address"`
	AssetContractType           string  `json:"asset_contract_type" camel:"assetContractType"`
	CreatedDate                 string  `json:"created_date" camel:"createdDate"`
	Name                        string  `json:"name" camel:"name"`
	NftVersion                  string  `json:"nft_version" camel:"nftVersion"`
	OpenseaVersion              *string `json:"opensea_version" camel:"openseaVersion"`
	Owner                       int     `json:"owner" camel:"owner"`
	SchemaName                  string  `json:"schema_name" camel:"schemaName"`
	Symbol                      string  `json:"symbol" camel:"symbol"`
	TotalSupply                 string  `json:"total_supply" camel:"totalSupply"`
	Description                 string  `json:"description" camel:"description"`
	ExternalLink                string  `json:"external_link" camel:"externalLink"`
	ImageURL                    string  `json:"image_url" camel:"imageURL"`
	DefaultToFiat               bool    `json:"default_to_fiat" camel:"defaultToFiat"`
	DevBuyerFeeBasisPoints      int     `json:"dev_buyer_fee_basis_points" camel:"devBuyerFeeBasisPoints"`
	DevSellerFeeBasisPoints     int     `json:"dev_seller_fee_basis_points" camel:"devSellerFeeBasisPoints"`
	OnlyProxiedTransfers        bool    `json:"only_proxied_transfers" camel:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  int     `json:"opensea_buyer_fee_basis_points" camel:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints int     `json:"opensea_seller_fee_basis_points" camel:"openseaSellerFeeBasisPoints"`
	BuyerFeeBasisPoints         int     `json:"buyer_fee_basis_points" camel:"buyerFeeBasisPoints"`
	SellerFeeBasisPoints        int     `json:"seller_fee_basis_points" camel:"sellerFeeBasisPoints"`
	PayoutAddress               string  `json:"payout_address" camel:"payoutAddress"`
}

type CollectionStats struct {
	OneDayVolume          float64 `json:"one_day_volume" camel:"oneDayVolume"`
	OneDayChange          float64 `json:"one_day_change" camel:"oneDayChange"`
	OneDaySales           float64 `json:"one_day_sales" camel:"oneDaySales"`
	OneDayAveragePrice    float64 `json:"one_day_average_price" camel:"oneDayAveragePrice"`
	SevenDayVolume        float64 `json:"seven_day_volume" camel:"sevenDayVolume"`
	SevenDayChange        float64 `json:"seven_day_change" camel:"sevenDayChange"`
	SevenDaySales         float64 `json:"seven_day_sales" camel:"sevenDaySales"`
	SevenDayAveragePrice  float64 `json:"seven_day_average_price" camel:"sevenDayAveragePrice"`
	ThirtyDayVolume       float64 `json:"thirty_day_volume" camel:"thirtyDayVolume"`
	ThirtyDayChange       float64 `json:"thirty_day_change" camel:"thirtyDayChange"`
	ThirtyDaySales        float64 `json:"thirty_day_sales" camel:"thirtyDaySales"`
	ThirtyDayAveragePrice float64 `json:"thirty_day_average_price" camel:"thirtyDayAveragePrice"`
	TotalVolume           float64 `json:"total_volume" camel:"totalVolume"`
	TotalSales            float64 `json:"total_sales" camel:"totalSales"`
	TotalSupply           float64 `json:"total_supply" camel:"totalSupply"`
	Count                 float64 `json:"count" camel:"count"`
	NumOwners             int     `json:"num_owners" camel:"numOwners"`
	AveragePrice          float64 `json:"average_price" camel:"averagePrice"`
	NumReports            int     `json:"num_reports" camel:"numReports"`
	MarketCap             float64 `json:"market_cap" camel:"marketCap"`
	FloorPrice            float64 `json:"floor_price" camel:"floorPrice"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style" camel:"cardDisplayStyle"`
}
