package model

// Asset The primary object in the OpenSea API is the asset,
// which represents a unique digital item whose ownership is managed by the blockchain.
// The below CryptoSaga hero is an example of an asset shown on OpenSea.
type Asset struct {
	// OpenSea NFT ID
	ID int `json:"id"`
	// The token ID of the NFT
	TokenID  string `json:"token_id"`
	NumSales int    `json:"num_sales"`
	// The background color to be displayed with the item
	BackgroundColor interface{} `json:"background_color"`
	// An image for the item. Note that this is the cached URL we store on our end. The original image url is image_original_url
	ImageURL             string      `json:"image_url"`
	ImagePreviewURL      string      `json:"image_preview_url"`
	ImageThumbnailURL    string      `json:"image_thumbnail_url"`
	ImageOriginalURL     string      `json:"image_original_url"`
	AnimationURL         interface{} `json:"animation_url"`
	AnimationOriginalURL interface{} `json:"animation_original_url"`
	// Name of the item
	Name        string      `json:"name"`
	Description interface{} `json:"description"`
	// External link to the original website for the item
	ExternalLink interface{} `json:"external_link"`
	// Dictionary of data on the contract itself (see asset contract section)
	AssetContract *Contract   `json:"asset_contract"`
	Permalink     string      `json:"permalink"`
	Collection    *Collection `json:"collection"`
	Decimals      int         `json:"decimals"`
	TokenMetadata string      `json:"token_metadata"`
	// Dictionary of data on the owner (see account section)
	Owner      *Account    `json:"owner"`
	SellOrders interface{} `json:"sell_orders"`
	Creator    *Account    `json:"creator"`
	// A list of traits associated with the item (see traits section)
	Traits []*Trait `json:"traits"`
	// When this item was last sold (null if there was no last sale)
	LastSale                *LastSale       `json:"last_sale"`
	TopBid                  interface{}     `json:"top_bid"`
	ListingDate             interface{}     `json:"listing_date"`
	IsPresale               bool            `json:"is_presale"`
	TransferFeePaymentToken interface{}     `json:"transfer_fee_payment_token"`
	TransferFee             interface{}     `json:"transfer_fee"`
	RelatedAssets           []interface{}   `json:"related_assets"`
	Orders                  []*Order        `json:"orders"`
	Auctions                []interface{}   `json:"auctions"`
	SupportsWyvern          bool            `json:"supports_wyvern"`
	TopOwnerships           []*TopOwnership `json:"top_ownerships"`
	Ownership               interface{}     `json:"ownership"`
	HighestBuyerCommitment  interface{}     `json:"highest_buyer_commitment"`
}

type Transaction struct {
	BlockHash        string   `json:"block_hash"`
	BlockNumber      string   `json:"block_number"`
	FromAccount      *Account `json:"from_account"`
	ID               int      `json:"id"`
	Timestamp        string   `json:"timestamp"`
	ToAccount        *Account `json:"to_account"`
	TransactionHash  string   `json:"transaction_hash"`
	TransactionIndex string   `json:"transaction_index"`
}

type LastSale struct {
	Asset struct {
		TokenID  string `json:"token_id"`
		Decimals int    `json:"decimals"`
	} `json:"asset"`
	AssetBundle    interface{}   `json:"asset_bundle"`
	EventType      string        `json:"event_type"`
	EventTimestamp string        `json:"event_timestamp"`
	AuctionType    interface{}   `json:"auction_type"`
	TotalPrice     string        `json:"total_price"`
	PaymentToken   *PaymentToken `json:"payment_token"`
	Transaction    *Transaction  `json:"transaction"`
	CreatedDate    string        `json:"created_date"`
	Quantity       string        `json:"quantity"`
}

type TopOwnership struct {
	Owner    *Account `json:"owner"`
	Quantity string   `json:"quantity"`
}
