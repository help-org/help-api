package types

type Division struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	ParentId *int   `json:"parent_id"`
}

type Type string

type Directory struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`

	Listings []*Listing `json:"listings"`
	Ads      []*Ad      `json:"ads"`
}

type Feature struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Type     FeatureType `json:"type"`
	ParentId *int        `json:"parent_id"`
}

type Listing struct {
	Id   int         `json:"id"`
	Name string      `json:"name"`
	Type ListingType `json:"type"`
	// TODO rename to feature_internal_id FeatureInternalId
	// TODO should be hidden from response
	FeatureId  int       `json:"feature_id"`
	ParentId   int       `json:"parent_id"`
	Address    string    `json:"address"`
	ContactIds []int     `json:"contact_ids"`
	Details    *string   `json:"details"`
	Contacts   []Contact `json:"contacts"`
	// last_modified
}

type Contact struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Type     FeatureType `json:"type"`
	ParentId *int        `json:"parent_id"`
}

// TODO Implement these structs
type Ad struct {
	Type  AdType `json:"type"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
