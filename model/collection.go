package model

import "github.com/shopspring/decimal"

// Collection Collections are used to represent all the assets in a single (or multiple) contract addresses
// and help users group items from the same creator. They have one or more owners and are typically
// associated with important metadata such as creator royalties and descriptions.
// Visit it(https://docs.opensea.io/reference/collection-model) to learn anymore.
type Collection struct {
	// The collection name. Typically derived from the first contract imported to the collection but can be changed by the user
	Name string `opensea:"name" json:"name"`
	// Description for the model
	Description      string  `opensea:"description" json:"description"`
	ShortDescription *string `opensea:"short_description" json:"shortDescription"`
	// The collection slug that is used to link to the collection on OpenSea.
	// This value can change by the owner but must be unique across all collection slugs in OpenSea
	Slug string `opensea:"slug" json:"slug"`
	// External link to the original website for the collection
	ExternalURL string `opensea:"external_url" json:"externalURL"`
	// An image for the collection. Note that this is the cached URL we store on our end.
	// The original image url is image_original_url
	ImageURL      string  `opensea:"image_url" json:"imageURL"`
	LargeImageURL *string `opensea:"large_image_url" json:"largeImageURL"`
	// Approved editors on this collection.
	Editors []string `opensea:"editors" json:"editors"`
	// The payment tokens accepted for this collection
	PaymentTokens []*PaymentToken `opensea:"payment_tokens" json:"paymentTokens"`
	// A list of the contracts that are associated with this collection
	PrimaryAssetContracts []*PrimaryAssetContract `opensea:"primary_asset_contracts" json:"primaryAssetContracts"`
	// A dictionary listing all the trait types available within this collection
	Traits map[string]map[string]int32 `opensea:"traits" json:"traits"`
	// A dictionary containing some sales statistics related to this collection, including trade volume and floor prices
	Stats *CollectionStats `opensea:"stats" json:"stats"`
	// Image used in the horizontal top banner for the collection.
	BannerImageURL string `opensea:"banner_image_url" json:"bannerImageURL"`
	// The payout address for the collection's royalties
	PayoutAddress string `opensea:"payout_address" json:"payoutAddress"`
	// The collector's fees that get paid out to them when sales are made for their collections
	DevSellerFeeBasisPoints string `opensea:"dev_seller_fee_basis_points" json:"devSellerFeeBasisPoints"`
	// The collection's approval status within OpenSea.
	// Can be not_requested (brand new collections), requested (collections that requested safelisting on our site),
	// approved (collections that are approved on our site and can be found in search results),
	// and verified (verified collections)
	SafelistRequestStatus string `opensea:"safelist_request_status" json:"safelistRequestStatus"`

	CreatedDate                 string      `opensea:"created_date" json:"createdDate"`
	DefaultToFiat               bool        `opensea:"default_to_fiat" json:"defaultToFiat"`
	DevBuyerFeeBasisPoints      string      `opensea:"dev_buyer_fee_basis_points" json:"devBuyerFeeBasisPoints"`
	DisplayData                 DisplayData `opensea:"display_data" json:"displayData"`
	Featured                    bool        `opensea:"featured" json:"featured"`
	FeaturedImageURL            *string     `opensea:"featured_image_url" json:"featuredImageURL"`
	Hidden                      bool        `opensea:"hidden" json:"hidden"`
	IsSubjectToWhitelist        bool        `opensea:"is_subject_to_whitelist" json:"isSubjectToWhitelist"`
	OnlyProxiedTransfers        bool        `opensea:"only_proxied_transfers" json:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  string      `opensea:"opensea_buyer_fee_basis_points" json:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints string      `opensea:"opensea_seller_fee_basis_points" json:"openseaSellerFeeBasisPoints"`
	RequireEmail                bool        `opensea:"require_email" json:"requireEmail"`

	TwitterUsername   *string `opensea:"twitter_username" json:"twitterUsername"`
	InstagramUsername string  `opensea:"instagram_username" json:"instagramUsername"`
	MediumUsername    *string `opensea:"medium_username" json:"mediumUsername"`

	TelegramURL *string `opensea:"telegram_url" json:"telegramURL"`
	DiscordURL  string  `opensea:"discord_url" json:"discordURL"`
	ChatURL     *string `opensea:"chat_url" json:"chatURL"`
	WikiURL     *string `opensea:"wiki_url" json:"wikiURL"`
}

type PaymentToken struct {
	ID       int              `opensea:"id" json:"id"`
	Symbol   string           `opensea:"symbol" json:"symbol"`
	Address  string           `opensea:"address" json:"address"`
	ImageURL string           `opensea:"image_url" json:"imageURL"`
	Name     string           `opensea:"name" json:"name"`
	Decimals int              `opensea:"decimals" json:"decimals"`
	EthPrice *decimal.Decimal `opensea:"eth_price" json:"ethPrice"`
	UsdPrice *decimal.Decimal `opensea:"usd_price" json:"usdPrice"`
}

type PrimaryAssetContract struct {
	Address                     string  `opensea:"address" json:"address"`
	AssetContractType           string  `opensea:"asset_contract_type" json:"assetContractType"`
	CreatedDate                 string  `opensea:"created_date" json:"createdDate"`
	Name                        string  `opensea:"name" json:"name"`
	NftVersion                  string  `opensea:"nft_version" json:"nftVersion"`
	OpenseaVersion              *string `opensea:"opensea_version" json:"openseaVersion"`
	Owner                       int     `opensea:"owner" json:"owner"`
	SchemaName                  string  `opensea:"schema_name" json:"schemaName"`
	Symbol                      string  `opensea:"symbol" json:"symbol"`
	TotalSupply                 string  `opensea:"total_supply" json:"totalSupply"`
	Description                 string  `opensea:"description" json:"description"`
	ExternalLink                string  `opensea:"external_link" json:"externalLink"`
	ImageURL                    string  `opensea:"image_url" json:"imageURL"`
	DefaultToFiat               bool    `opensea:"default_to_fiat" json:"defaultToFiat"`
	DevBuyerFeeBasisPoints      int     `opensea:"dev_buyer_fee_basis_points" json:"devBuyerFeeBasisPoints"`
	DevSellerFeeBasisPoints     int     `opensea:"dev_seller_fee_basis_points" json:"devSellerFeeBasisPoints"`
	OnlyProxiedTransfers        bool    `opensea:"only_proxied_transfers" json:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  int     `opensea:"opensea_buyer_fee_basis_points" json:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints int     `opensea:"opensea_seller_fee_basis_points" json:"openseaSellerFeeBasisPoints"`
	BuyerFeeBasisPoints         int     `opensea:"buyer_fee_basis_points" json:"buyerFeeBasisPoints"`
	SellerFeeBasisPoints        int     `opensea:"seller_fee_basis_points" json:"sellerFeeBasisPoints"`
	PayoutAddress               string  `opensea:"payout_address" json:"payoutAddress"`
}

type CollectionStats struct {
	OneDayVolume          float64 `opensea:"one_day_volume" json:"oneDayVolume"`
	OneDayChange          float64 `opensea:"one_day_change" json:"oneDayChange"`
	OneDaySales           float64 `opensea:"one_day_sales" json:"oneDaySales"`
	OneDayAveragePrice    float64 `opensea:"one_day_average_price" json:"oneDayAveragePrice"`
	SevenDayVolume        float64 `opensea:"seven_day_volume" json:"sevenDayVolume"`
	SevenDayChange        float64 `opensea:"seven_day_change" json:"sevenDayChange"`
	SevenDaySales         float64 `opensea:"seven_day_sales" json:"sevenDaySales"`
	SevenDayAveragePrice  float64 `opensea:"seven_day_average_price" json:"sevenDayAveragePrice"`
	ThirtyDayVolume       float64 `opensea:"thirty_day_volume" json:"thirtyDayVolume"`
	ThirtyDayChange       float64 `opensea:"thirty_day_change" json:"thirtyDayChange"`
	ThirtyDaySales        float64 `opensea:"thirty_day_sales" json:"thirtyDaySales"`
	ThirtyDayAveragePrice float64 `opensea:"thirty_day_average_price" json:"thirtyDayAveragePrice"`
	TotalVolume           float64 `opensea:"total_volume" json:"totalVolume"`
	TotalSales            float64 `opensea:"total_sales" json:"totalSales"`
	TotalSupply           float64 `opensea:"total_supply" json:"totalSupply"`
	Count                 float64 `opensea:"count" json:"count"`
	NumOwners             int     `opensea:"num_owners" json:"numOwners"`
	AveragePrice          float64 `opensea:"average_price" json:"averagePrice"`
	NumReports            int     `opensea:"num_reports" json:"numReports"`
	MarketCap             float64 `opensea:"market_cap" json:"marketCap"`
	FloorPrice            float64 `opensea:"floor_price" json:"floorPrice"`
}

type DisplayData struct {
	CardDisplayStyle string `opensea:"card_display_style" json:"cardDisplayStyle"`
}
