package types

type ListingType string

const (
	POLICE                ListingType = "Police"
	HOSPITAL              ListingType = "Hospital"
	FIRE_DEPARTMENT       ListingType = "Fire Department"
	AMBULANCE             ListingType = "Ambulance"
	POISON_CONTROL        ListingType = "Poison Control"
	COAST_GUARD           ListingType = "Coast Guard"
	ELECTRICITY_EMERGENCY ListingType = "Electricity Emergency"
	GAS_LEAK_EMERGENCY    ListingType = "Gas Leak Emergency"
	ROAD_ASSISTANCE       ListingType = "Road Assistance"
	MENTAL_HEALTH         ListingType = "Mental Health"
	DOMESTIC_VIOLENCE     ListingType = "Domestic Violence"
	MISCELLANEOUS         ListingType = "Miscellaneous"
)

type FeatureType string

const (
	COUNTRY      FeatureType = "country"
	REGION       FeatureType = "region"
	POSTCODE     FeatureType = "postcode"
	DISTRICT     FeatureType = "district"
	PLACE        FeatureType = "place"
	LOCALITY     FeatureType = "locality"
	NEIGHBORHOOD FeatureType = "neighborhood"
)

type ContactType string

const (
	EMAIL ContactType = "Email"
	PHONE ContactType = "Phone"
)

type AdType string

const (
	LAWYER AdType = "Lawyer"
	DOCTOR AdType = "Doctor"
)
