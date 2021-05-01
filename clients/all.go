package clients

import (
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/clients/restclient"
	"github.com/l30n4rd01b4rr4t/covid-global-numbers-by-console/domain"
)

func GetAll(all *domain.All) error {
	client := restclient.NewClient()
	client.Get("/v3/covid-19/all")
	response := client.Response(all)
	return response
}
