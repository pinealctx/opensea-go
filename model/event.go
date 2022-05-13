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
	ApprovedAccount *AccountID `opensea:"approved_account" json:"approvedAccount"`
	// A subfield containing a simplified version of the Asset or Asset Bundle on which this event happened
	Asset *Asset `opensea:"asset" json:"asset"`
	// Ditto
	AssetBundle     *Bundle      `opensea:"asset_bundle" json:"assetBundle"`
	AuctionType     *AuctionType `opensea:"auction_type" json:"auctionType"`
	BidAmount       string       `opensea:"bid_amount" json:"bidAmount"`
	CollectionSlug  string       `opensea:"collection_slug" json:"collectionSlug"`
	ContractAddress string       `opensea:"contract_address" json:"contractAddress"`
	// When the event was recorded
	CreatedDate             string           `opensea:"created_date" json:"createdDate"`
	CustomEventName         *string          `opensea:"custom_event_name" json:"customEventName"`
	DevFeePaymentEvent      interface{}      `opensea:"dev_fee_payment_event" json:"devFeePaymentEvent"`
	DevSellerFeeBasisPoints int              `opensea:"dev_seller_fee_basis_points" json:"devSellerFeeBasisPoints"`
	Duration                *int32           `opensea:"duration,string" json:"duration,string"`
	EndingPrice             *decimal.Decimal `opensea:"ending_price" json:"endingPrice"`
	// Describes the event type
	EventType EventType `opensea:"event_type" json:"eventType"`
	// The accounts associated with this event.
	FromAccount *AccountID `opensea:"from_account" json:"fromAccount"`
	// Ditto
	ToAccount *AccountID `opensea:"to_account" json:"toAccount"`
	ID        int64      `opensea:"id" json:"id"`
	// A boolean value that is true if the sale event was a private sale
	IsPrivate    *bool      `opensea:"is_private" json:"isPrivate"`
	OwnerAccount *AccountID `opensea:"owner_account" json:"ownerAccount"`
	// The payment asset used in this transaction, such as ETH, WETH or DAI
	PaymentToken *PaymentToken `opensea:"payment_token" json:"paymentToken"`
	// The amount of the item that was sold. Applicable for semi-fungible assets
	Quantity      string           `opensea:"quantity" json:"quantity"`
	Seller        *AccountID       `opensea:"seller" json:"seller"`
	StartingPrice *decimal.Decimal `opensea:"starting_price" json:"startingPrice"`
	// The total price that the asset was bought for. This includes any royalties that might have been collected
	TotalPrice    *decimal.Decimal `opensea:"total_price" json:"totalPrice"`
	Transaction   *Transaction     `opensea:"transaction" json:"transaction"`
	WinnerAccount *AccountID       `opensea:"winner_account" json:"winnerAccount"`
	// eg: 2017-07-21T17:32:28Z
	ListingTime *string `opensea:"listing_time" json:"listingTime"`
}
