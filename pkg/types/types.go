package types

<<<<<<< Updated upstream
type Division struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Type     string   `json:"type"`
	ParentId *int   `json:"parent_id"`
}

type Type string

const (
	COUNTRY      Type = "country"
	STATE        Type = "state"
	PROVINCE     Type = "province"
	OBLAST       Type = "oblast"
	LAND         Type = "land"
	REGION       Type = "region"
	COMARCA      Type = "comarca"
	RAION        Type = "raion"
	DISTRICT     Type = "district"
	MUNICIPALITY Type = "municipality"
	COMMUNE      Type = "commune"
	COMMUNITY    Type = "community"
	DEPARTMENT   Type = "department"
	CANTON       Type = "canton"
	PREFECTURE   Type = "prefecture"
	COUNTY       Type = "county"
	GOVERNORATE  Type = "governorate"
)

type Directory struct {
	Country string `json:"country"`
	State   string `json:"state"`
	City    string `json:"city"`

	Listings []*Listing `json:"listings"`
	Ads      []*Ad      `json:"ads"`
=======
type Feature struct {
	Id       int         `json:"id"`
	Name     string      `json:"name"`
	Type     FeatureType `json:"type"`
	ParentId *int        `json:"parent_id"`
>>>>>>> Stashed changes
}

type Listing struct {
	Id   int         `json:"id"`
	Name string      `json:"name"`
	Type ListingType `json:"type"`
	// TODO rename to feature_internal_id FeatureInternalId
	// TODO should be hidden from response
	FeatureId  int       `json:"feature_id"`
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
