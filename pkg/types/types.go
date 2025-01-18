package types

type Type string

type Feature struct {
	Id         string      `json:"id"`
	InternalId int         `json:",omitempty" db:"internal_id"`
	Name       string      `json:"name"`
	Type       FeatureType `json:"type"`
	ParentId   *int        `json:"parent_id"`
	Listings   []*Listing  `json:"listings"`
}

type Listing struct {
	Id         string      `json:"id"`
	InternalId int         `json:",omitempty" db:"internal_id"`
	Name       string      `json:"name"`
	Type       ListingType `json:"type"`
	FeatureId  int         `json:"feature_id"`
	Address    string      `json:"address"`
	ContactIds []int       `json:"contact_ids"`
	Details    *string     `json:"details"`
	Contacts   []Contact   `json:"contacts"`
}

type Contact struct {
	Id         string      `json:"id"`
	InternalId int         `json:",omitempty" db:"internal_id"`
	Name       string      `json:"name"`
	Type       ContactType `json:"type"`
	Details    *string     `json:"details"`
}

// TODO Implement these structs
type Ad struct {
	Type  AdType `json:"type"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type Directory struct {
	Country  string     `json:"country"`
	State    string     `json:"state"`
	City     string     `json:"city"`
	Listings []*Listing `json:"listings"`
	Ads      []*Ad      `json:"ads"`
}
