package clients

import (
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/clients/restclient"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
)

func GetCountries(countries *[]domain.Country) error {
	client := restclient.NewClient()
	client.Get("/v3/covid-19/countries")
	response := client.Response(countries)
	return response
}
