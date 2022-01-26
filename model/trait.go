package model

type TraitDisplayType string

const (
	TraitDTNumber          TraitDisplayType = "string"
	TraitDTBoostPercentage TraitDisplayType = "boost_percentage"
	TraitDTBoostNumber     TraitDisplayType = "boost_number"
	TraitDTDate            TraitDisplayType = "date"
)

// Trait Traits are special properties on the item, that can either be numbers or strings.
// Below is an example of how OpenSea displays the traits for a specific item.
type Trait struct {
	// The name of the trait (for example color)
	TraitType string `json:"trait_type" camel:"traitType"`
	// The value of this trait (can be a string or number)
	Value interface{} `json:"value" camel:"value"`
	// How this trait will be displayed (options are number, boost_percentage, boost_number, and date).
	// See the adding metadata section for more details
	DisplayType *TraitDisplayType `json:"display_type" camel:"displayType"`
	MaxValue    interface{}       `json:"max_value" camel:"maxValue"`
	TraitCount  int               `json:"trait_count" camel:"traitCount"`
	Order       interface{}       `json:"order" camel:"order"`
}
