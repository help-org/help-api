package types

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
}

type Listing struct {
	Type  ListingType `json:"type"`
	Name  string      `json:"name"`
	Phone string      `json:"phone"`
}

type Ad struct {
	Type  AdType `json:"type"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type ListingType string

const (
	POLICE ListingType = "Police"
	FIRE   ListingType = "Fire"
	EMS    ListingType = "EMS"
)

type AdType string

const (
	LAWYER AdType = "Lawyer"
	DOCTOR AdType = "Doctor"
)
