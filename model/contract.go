package model

// ContractType Types of asset contracts
// Given by the asset_contract_type in the OpenSea API
type ContractType string

const (
	Fungible     ContractType = "fungible"
	SemiFungible ContractType = "semi-fungible"
	NonFungible  ContractType = "non-fungible"
	Unknown      ContractType = "unknown"
)

// Contract Asset contracts contain data about the contract itself,
// such as the CryptoKitties contract or the CoolCats contract.
type Contract struct {
	// Address of the asset contract
	Address           string       `json:"address" camel:"address"`
	AssetContractType ContractType `json:"asset_contract_type" camel:"assetContractType"`
	CreatedDate       string       `json:"created_date" camel:"createdDate"`
	// Name of the asset contract
	Name           string  `json:"name" camel:"name"`
	NftVersion     string  `json:"nft_version" camel:"nftVersion"`
	OpenseaVersion *string `json:"opensea_version" camel:"openseaVersion"`
	Owner          int     `json:"owner" camel:"owner"`
	SchemaName     string  `json:"schema_name" camel:"schemaName"`
	// Symbol, such as CKITTY
	Symbol      string `json:"symbol" camel:"symbol"`
	TotalSupply string `json:"total_supply" camel:"totalSupply"`
	// Description of the asset contract
	Description string `json:"description" camel:"description"`
	// Link to the original website for this contract
	ExternalLink string `json:"external_link" camel:"externalLink"`
	// Image associated with the asset contract
	ImageURL                    string      `json:"image_url" camel:"imageURL"`
	DefaultToFiat               bool        `json:"default_to_fiat" camel:"defaultToFiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points" camel:"devBuyerFeeBasisPoints"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points" camel:"devSellerFeeBasisPoints"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers" camel:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points" camel:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points" camel:"openseaSellerFeeBasisPoints"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points" camel:"buyerFeeBasisPoints"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points" camel:"sellerFeeBasisPoints"`
	PayoutAddress               string      `json:"payout_address" camel:"payoutAddress"`
	Collection                  *Collection `json:"collection,omitempty" camel:"collection"`
}
