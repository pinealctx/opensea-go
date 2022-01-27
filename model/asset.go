package model

// Asset The primary object in the OpenSea API is the asset,
// which represents a unique digital item whose ownership is managed by the blockchain.
// The below CryptoSaga hero is an example of an asset shown on OpenSea.
type Asset struct {
	// OpenSea NFT ID
	ID int `opensea:"id" json:"id"`
	// The token ID of the NFT
	TokenID  string `opensea:"token_id" json:"tokenID"`
	NumSales int    `opensea:"num_sales" json:"numSales"`
	// The background color to be displayed with the item
	BackgroundColor *string `opensea:"background_color" json:"backgroundColor"`
	// An image for the item. Note that this is the cached URL we store on our end. The original image url is image_original_url
	ImageURL             string  `opensea:"image_url" json:"imageURL"`
	ImagePreviewURL      string  `opensea:"image_preview_url" json:"imagePreviewURL"`
	ImageThumbnailURL    string  `opensea:"image_thumbnail_url" json:"imageThumbnailURL"`
	ImageOriginalURL     string  `opensea:"image_original_url" json:"imageOriginalURL"`
	AnimationURL         *string `opensea:"animation_url" json:"animationURL"`
	AnimationOriginalURL *string `opensea:"animation_original_url" json:"animationOriginalURL"`
	// Name of the item
	Name        string  `opensea:"name" json:"name"`
	Description *string `opensea:"description" json:"description"`
	// External link to the original website for the item
	ExternalLink *string `opensea:"external_link" json:"externalLink"`
	// Dictionary of data on the contract itself (see asset contract section)
	AssetContract *Contract   `opensea:"asset_contract" json:"assetContract"`
	Permalink     string      `opensea:"permalink" json:"permalink"`
	Collection    *Collection `opensea:"collection" json:"collection"`
	Decimals      int         `opensea:"decimals" json:"decimals"`
	TokenMetadata string      `opensea:"token_metadata" json:"tokenMetadata"`
	// Dictionary of data on the owner (see account section)
	Owner      *Account     `opensea:"owner" json:"owner"`
	SellOrders []*SellOrder `opensea:"sell_orders,omitempty" json:"sellOrders,omitempty"`
	Creator    *Account     `opensea:"creator" json:"creator"`
	// A list of traits associated with the item (see traits section)
	Traits []*Trait `opensea:"traits" json:"traits"`
	// When this item was last sold (null if there was no last sale)
	LastSale                *LastSale       `opensea:"last_sale" json:"lastSale"`
	TopBid                  interface{}     `opensea:"top_bid" json:"topBid"`
	ListingDate             *string         `opensea:"listing_date" json:"listingDate"`
	IsPresale               bool            `opensea:"is_presale" json:"isPresale"`
	TransferFeePaymentToken interface{}     `opensea:"transfer_fee_payment_token" json:"transferFeePaymentToken"`
	TransferFee             interface{}     `opensea:"transfer_fee" json:"transferFee"`
	RelatedAssets           []*Asset        `opensea:"related_assets" json:"relatedAssets"`
	Orders                  []*Order        `opensea:"orders" json:"orders"`
	Auctions                []interface{}   `opensea:"auctions" json:"auctions"`
	SupportsWyvern          bool            `opensea:"supports_wyvern" json:"supportsWyvern"`
	TopOwnerships           []*TopOwnership `opensea:"top_ownerships" json:"topOwnerships"`
	Ownership               interface{}     `opensea:"ownership" json:"ownership"`
	HighestBuyerCommitment  interface{}     `opensea:"highest_buyer_commitment" json:"highestBuyerCommitment"`
}

type Transaction struct {
	BlockHash        string   `opensea:"block_hash" json:"blockHash"`
	BlockNumber      string   `opensea:"block_number" json:"blockNumber"`
	FromAccount      *Account `opensea:"from_account" json:"fromAccount"`
	ID               int      `opensea:"id" json:"id"`
	Timestamp        string   `opensea:"timestamp" json:"timestamp"`
	ToAccount        *Account `opensea:"to_account" json:"toAccount"`
	TransactionHash  string   `opensea:"transaction_hash" json:"transactionHash"`
	TransactionIndex string   `opensea:"transaction_index" json:"transactionIndex"`
}

type LastSale struct {
	Asset struct {
		TokenID  string `opensea:"token_id" json:"tokenID"`
		Decimals int    `opensea:"decimals" json:"decimals"`
	} `opensea:"asset" json:"asset"`
	AssetBundle    *Bundle       `opensea:"asset_bundle" json:"assetBundle"`
	EventType      string        `opensea:"event_type" json:"eventType"`
	EventTimestamp string        `opensea:"event_timestamp" json:"eventTimestamp"`
	AuctionType    *AuctionType  `opensea:"auction_type" json:"auctionType"`
	TotalPrice     string        `opensea:"total_price" json:"totalPrice"`
	PaymentToken   *PaymentToken `opensea:"payment_token" json:"paymentToken"`
	Transaction    *Transaction  `opensea:"transaction" json:"transaction"`
	CreatedDate    string        `opensea:"created_date" json:"createdDate"`
	Quantity       string        `opensea:"quantity" json:"quantity"`
}

type TopOwnership struct {
	Owner    *Account `opensea:"owner" json:"owner"`
	Quantity string   `opensea:"quantity" json:"quantity"`
}

type AssetAddress struct {
	ID      string `opensea:"id" json:"id"`
	Address string `opensea:"address" json:"address"`
}
