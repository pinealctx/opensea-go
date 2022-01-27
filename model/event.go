package model

import "github.com/shopspring/decimal"

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
	ApprovedAccount *Account `json:"approved_account" camel:"approvedAccount"`
	// A subfield containing a simplified version of the Asset or Asset Bundle on which this event happened
	Asset *Asset `json:"asset" camel:"asset"`
	// Ditto
	AssetBundle     *Bundle      `json:"asset_bundle" camel:"assetBundle"`
	AuctionType     *AuctionType `json:"auction_type" camel:"auctionType"`
	BidAmount       string       `json:"bid_amount" camel:"bidAmount"`
	CollectionSlug  string       `json:"collection_slug" camel:"collectionSlug"`
	ContractAddress string       `json:"contract_address" camel:"contractAddress"`
	// When the event was recorded
	CreatedDate             string           `json:"created_date" camel:"createdDate"`
	CustomEventName         *string          `json:"custom_event_name" camel:"customEventName"`
	DevFeePaymentEvent      interface{}      `json:"dev_fee_payment_event" camel:"devFeePaymentEvent"`
	DevSellerFeeBasisPoints int              `json:"dev_seller_fee_basis_points" camel:"devSellerFeeBasisPoints"`
	Duration                *int32           `json:"duration,string" camel:"duration,string"`
	EndingPrice             *decimal.Decimal `json:"ending_price" camel:"endingPrice"`
	// Describes the event type
	EventType EventType `json:"event_type" camel:"eventType"`
	// The accounts associated with this event.
	FromAccount *Account `json:"from_account" camel:"fromAccount"`
	// Ditto
	ToAccount *Account `json:"to_account" camel:"toAccount"`
	ID        int64    `json:"id" camel:"id"`
	// A boolean value that is true if the sale event was a private sale
	IsPrivate    *bool    `json:"is_private" camel:"isPrivate"`
	OwnerAccount *Account `json:"owner_account" camel:"ownerAccount"`
	// The payment asset used in this transaction, such as ETH, WETH or DAI
	PaymentToken *PaymentToken `json:"payment_token" camel:"paymentToken"`
	// The amount of the item that was sold. Applicable for semi-fungible assets
	Quantity      string           `json:"quantity" camel:"quantity"`
	Seller        *Account         `json:"seller" camel:"seller"`
	StartingPrice *decimal.Decimal `json:"starting_price" camel:"startingPrice"`
	// The total price that the asset was bought for. This includes any royalties that might have been collected
	TotalPrice    interface{} `json:"total_price" camel:"totalPrice"`
	Transaction   interface{} `json:"transaction" camel:"transaction"`
	WinnerAccount interface{} `json:"winner_account" camel:"winnerAccount"`
	// eg: 2017-07-21T17:32:28Z
	ListingTime *string `json:"listing_time" camel:"listingTime"`
}
