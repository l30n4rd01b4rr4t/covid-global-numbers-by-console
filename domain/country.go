package domain

type Country struct {
	Updated                int64   `json:"updated"`
	Cases                  int64   `json:"cases"`
	TodayCases             int64   `json:"todayCases"`
	Deaths                 int64   `json:"deaths"`
	TodayDeaths            int64   `json:"todayDeaths"`
	Recovered              int64   `json:"recovered"`
	Active                 int64   `json:"active"`
	Critical               int64   `json:"critical"`
	CasesPerOneMillion     float64 `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float64 `json:"deathsPerOneMillion"`
	Tests                  int64   `json:"tests"`
	TestsPerOneMillion     float64 `json:"testsPerOneMillion"`
	Population             int64   `json:"population"`
	ActivePerOneMillion    float64 `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64 `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float64 `json:"criticalPerOneMillion"`
	Continent              string  `json:"continent"`
	Country                string  `json:"country"`
	CountryInfo            struct {
		ID   int     `json:"_id"`
		Iso2 string  `json:"iso2"`
		Iso3 string  `json:"iso3"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
		Flag string  `json:"flag"`
	} `json:"countryInfo"`
}
