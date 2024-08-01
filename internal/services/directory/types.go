package directory

type Directory struct {
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
