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
	Address           string       `json:"address"`
	AssetContractType ContractType `json:"asset_contract_type"`
	CreatedDate       string       `json:"created_date"`
	// Name of the asset contract
	Name           string  `json:"name"`
	NftVersion     string  `json:"nft_version"`
	OpenseaVersion *string `json:"opensea_version"`
	Owner          int     `json:"owner"`
	SchemaName     string  `json:"schema_name"`
	// Symbol, such as CKITTY
	Symbol      string `json:"symbol"`
	TotalSupply string `json:"total_supply"`
	// Description of the asset contract
	Description string `json:"description"`
	// Link to the original website for this contract
	ExternalLink string `json:"external_link"`
	// Image associated with the asset contract
	ImageURL                    string      `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
	Collection                  *Collection `json:"collection,omitempty"`
}
