package model

// EventType Describes the event type
type EventType string

const (
	// ETCreated new auctions
	ETCreated EventType = "created"
	// ETSuccessful sales
	ETSuccessful EventType = "successful"
	// ETCancelled cancelled auctions
	ETCancelled    EventType = "cancelled"
	ETBidEntered   EventType = "bid_entered"
	ETBidWithdrawn EventType = "bid_withdrawn"
	ETTransfer     EventType = "transfer"
	ETOfferEntered EventType = "offer_entered"
	ETApprove      EventType = "approve"
)

// AuctionType Auction type
type AuctionType string

const (
	// ATEnglish English Auctions
	ATEnglish AuctionType = "english"
	// ATDutch fixed-price and declining-price sell orders (Dutch Auctions)
	ATDutch AuctionType = "dutch"
	// ATMinPrice CryptoPunks bidding auctions
	ATMinPrice AuctionType = "min-price"
)

// Event Asset events represent state changes that occur for assets.
// This includes putting them on sale, bidding on them, selling them, cancelling sales, transferring them, and more.
type Event struct {
	ApprovedAccount interface{} `json:"approved_account"`
	// A subfield containing a simplified version of the Asset or Asset Bundle on which this event happened
	Asset *Asset `json:"asset"`
	// Ditto
	AssetBundle     *Bundle      `json:"asset_bundle"`
	AuctionType     *AuctionType `json:"auction_type"`
	BidAmount       string       `json:"bid_amount"`
	CollectionSlug  string       `json:"collection_slug"`
	ContractAddress string       `json:"contract_address"`
	// When the event was recorded
	CreatedDate             string      `json:"created_date"`
	CustomEventName         interface{} `json:"custom_event_name"`
	DevFeePaymentEvent      interface{} `json:"dev_fee_payment_event"`
	DevSellerFeeBasisPoints int         `json:"dev_seller_fee_basis_points"`
	Duration                interface{} `json:"duration"`
	EndingPrice             interface{} `json:"ending_price"`
	// Describes the event type
	EventType string `json:"event_type"`
	// The accounts associated with this event.
	FromAccount *Account `json:"from_account"`
	// Ditto
	ToAccount *Account `json:"to_account"`
	ID        int64    `json:"id"`
	// A boolean value that is true if the sale event was a private sale
	IsPrivate    interface{} `json:"is_private"`
	OwnerAccount interface{} `json:"owner_account"`
	// The payment asset used in this transaction, such as ETH, WETH or DAI
	PaymentToken *PaymentToken `json:"payment_token"`
	// The amount of the item that was sold. Applicable for semi-fungible assets
	Quantity      string      `json:"quantity"`
	Seller        interface{} `json:"seller"`
	StartingPrice interface{} `json:"starting_price"`
	// The total price that the asset was bought for. This includes any royalties that might have been collected
	TotalPrice    interface{} `json:"total_price"`
	Transaction   interface{} `json:"transaction"`
	WinnerAccount interface{} `json:"winner_account"`
	ListingTime   interface{} `json:"listing_time"`
}
