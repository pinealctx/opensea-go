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
	Address           string       `opensea:"address" json:"address"`
	AssetContractType ContractType `opensea:"asset_contract_type" json:"assetContractType"`
	CreatedDate       string       `opensea:"created_date" json:"createdDate"`
	// Name of the asset contract
	Name           string  `opensea:"name" json:"name"`
	NftVersion     string  `opensea:"nft_version" json:"nftVersion"`
	OpenseaVersion *string `opensea:"opensea_version" json:"openseaVersion"`
	Owner          int     `opensea:"owner" json:"owner"`
	SchemaName     string  `opensea:"schema_name" json:"schemaName"`
	// Symbol, such as CKITTY
	Symbol      string `opensea:"symbol" json:"symbol"`
	TotalSupply string `opensea:"total_supply" json:"totalSupply"`
	// Description of the asset contract
	Description string `opensea:"description" json:"description"`
	// Link to the original website for this contract
	ExternalLink string `opensea:"external_link" json:"externalLink"`
	// Image associated with the asset contract
	ImageURL                    string      `opensea:"image_url" json:"imageURL"`
	DefaultToFiat               bool        `opensea:"default_to_fiat" json:"defaultToFiat"`
	DevBuyerFeeBasisPoints      int         `opensea:"dev_buyer_fee_basis_points" json:"devBuyerFeeBasisPoints"`
	DevSellerFeeBasisPoints     int         `opensea:"dev_seller_fee_basis_points" json:"devSellerFeeBasisPoints"`
	OnlyProxiedTransfers        bool        `opensea:"only_proxied_transfers" json:"onlyProxiedTransfers"`
	OpenseaBuyerFeeBasisPoints  int         `opensea:"opensea_buyer_fee_basis_points" json:"openseaBuyerFeeBasisPoints"`
	OpenseaSellerFeeBasisPoints int         `opensea:"opensea_seller_fee_basis_points" json:"openseaSellerFeeBasisPoints"`
	BuyerFeeBasisPoints         int         `opensea:"buyer_fee_basis_points" json:"buyerFeeBasisPoints"`
	SellerFeeBasisPoints        int         `opensea:"seller_fee_basis_points" json:"sellerFeeBasisPoints"`
	PayoutAddress               string      `opensea:"payout_address" json:"payoutAddress"`
	Collection                  *Collection `opensea:"collection,omitempty" json:"collection"`
}
