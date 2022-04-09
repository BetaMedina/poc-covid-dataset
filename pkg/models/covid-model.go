package models

type CovidModel struct {
	Date                           string  `json:"date,omitempty"`
	State                          string  `json:"state,omitempty"`
	City                           string  `json:"city,omitempty"`
	Place_type                     string  `json:"placeType,omitempty"`
	Confirmed                      int32   `json:"confirmed,omitempty"`
	Deaths                         int32   `json:"deaths,omitempty"`
	Order_for_place                int32   `json:"orderForPlace,omitempty"`
	Estimated_population_2019      int32   `json:"estimatedPopulation2019,omitempty"`
	Estimated_population           int32   `json:"estimatedPopulation,omitempty"`
	City_ibge_code                 int32   `json:"ibgeCode,omitempty"`
	Confirmed_per_100k_inhabitants float64 `json:"confirmedPer100kInhabitants,omitempty"`
	Death_rate                     float64 `json:"deathRatew,omitempty"`
}
