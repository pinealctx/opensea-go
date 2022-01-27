package model

// Asset The primary object in the OpenSea API is the asset,
// which represents a unique digital item whose ownership is managed by the blockchain.
// The below CryptoSaga hero is an example of an asset shown on OpenSea.
type Asset struct {
	// OpenSea NFT ID
	ID int `json:"id" camel:"id"`
	// The token ID of the NFT
	TokenID  string `json:"token_id" camel:"tokenID"`
	NumSales int    `json:"num_sales" camel:"numSales"`
	// The background color to be displayed with the item
	BackgroundColor *string `json:"background_color" camel:"backgroundColor"`
	// An image for the item. Note that this is the cached URL we store on our end. The original image url is image_original_url
	ImageURL             string  `json:"image_url" camel:"imageURL"`
	ImagePreviewURL      string  `json:"image_preview_url" camel:"imagePreviewURL"`
	ImageThumbnailURL    string  `json:"image_thumbnail_url" camel:"imageThumbnailURL"`
	ImageOriginalURL     string  `json:"image_original_url" camel:"imageOriginalURL"`
	AnimationURL         *string `json:"animation_url" camel:"animationURL"`
	AnimationOriginalURL *string `json:"animation_original_url" camel:"animationOriginalURL"`
	// Name of the item
	Name        string  `json:"name" camel:"name"`
	Description *string `json:"description" camel:"description"`
	// External link to the original website for the item
	ExternalLink *string `json:"external_link" camel:"externalLink"`
	// Dictionary of data on the contract itself (see asset contract section)
	AssetContract *Contract   `json:"asset_contract" camel:"assetContract"`
	Permalink     string      `json:"permalink" camel:"permalink"`
	Collection    *Collection `json:"collection" camel:"collection"`
	Decimals      int         `json:"decimals" camel:"decimals"`
	TokenMetadata string      `json:"token_metadata" camel:"tokenMetadata"`
	// Dictionary of data on the owner (see account section)
	Owner      *Account     `json:"owner" camel:"owner"`
	SellOrders []*SellOrder `json:"sell_orders,omitempty" camel:"sellOrders,omitempty"`
	Creator    *Account     `json:"creator" camel:"creator"`
	// A list of traits associated with the item (see traits section)
	Traits []*Trait `json:"traits" camel:"traits"`
	// When this item was last sold (null if there was no last sale)
	LastSale                *LastSale       `json:"last_sale" camel:"lastSale"`
	TopBid                  interface{}     `json:"top_bid" camel:"topBid"`
	ListingDate             *string         `json:"listing_date" camel:"listingDate"`
	IsPresale               bool            `json:"is_presale" camel:"isPresale"`
	TransferFeePaymentToken interface{}     `json:"transfer_fee_payment_token" camel:"transferFeePaymentToken"`
	TransferFee             interface{}     `json:"transfer_fee" camel:"transferFee"`
	RelatedAssets           []*Asset        `json:"related_assets" camel:"relatedAssets"`
	Orders                  []*Order        `json:"orders" camel:"orders"`
	Auctions                []interface{}   `json:"auctions" camel:"auctions"`
	SupportsWyvern          bool            `json:"supports_wyvern" camel:"supportsWyvern"`
	TopOwnerships           []*TopOwnership `json:"top_ownerships" camel:"topOwnerships"`
	Ownership               interface{}     `json:"ownership" camel:"ownership"`
	HighestBuyerCommitment  interface{}     `json:"highest_buyer_commitment" camel:"highestBuyerCommitment"`
}

type Transaction struct {
	BlockHash        string   `json:"block_hash" camel:"blockHash"`
	BlockNumber      string   `json:"block_number" camel:"blockNumber"`
	FromAccount      *Account `json:"from_account" camel:"fromAccount"`
	ID               int      `json:"id" camel:"id"`
	Timestamp        string   `json:"timestamp" camel:"timestamp"`
	ToAccount        *Account `json:"to_account" camel:"toAccount"`
	TransactionHash  string   `json:"transaction_hash" camel:"transactionHash"`
	TransactionIndex string   `json:"transaction_index" camel:"transactionIndex"`
}

type LastSale struct {
	Asset struct {
		TokenID  string `json:"token_id" camel:"tokenID"`
		Decimals int    `json:"decimals" camel:"decimals"`
	} `json:"asset" camel:"asset"`
	AssetBundle    *Bundle       `json:"asset_bundle" camel:"assetBundle"`
	EventType      string        `json:"event_type" camel:"eventType"`
	EventTimestamp string        `json:"event_timestamp" camel:"eventTimestamp"`
	AuctionType    *AuctionType  `json:"auction_type" camel:"auctionType"`
	TotalPrice     string        `json:"total_price" camel:"totalPrice"`
	PaymentToken   *PaymentToken `json:"payment_token" camel:"paymentToken"`
	Transaction    *Transaction  `json:"transaction" camel:"transaction"`
	CreatedDate    string        `json:"created_date" camel:"createdDate"`
	Quantity       string        `json:"quantity" camel:"quantity"`
}

type TopOwnership struct {
	Owner    *Account `json:"owner" camel:"owner"`
	Quantity string   `json:"quantity" camel:"quantity"`
}

type AssetAddress struct {
	ID      string `json:"id" camel:"id"`
	Address string `json:"address" camel:"address"`
}
