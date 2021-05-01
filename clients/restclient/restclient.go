package restclient

import (
	"github.com/dghubble/sling"
)

type httpclient struct {
	*sling.Sling
}

func NewClient() httpclient {
	return httpclient{sling.New().Base("https://disease.sh")}
}

func (client httpclient) Response(response interface{}) error {
	err := new(Error)
	_, e := client.Receive(response, err)
	return client.check(e, err)
}

// compare the errors and return the relevant one
func (client httpclient) check(e error, err *Error) error {
	if err.Exists() {
		return err
	}
	return e
}

type Error struct {
	Message string          `json:"message"`
	Errors  string          `json:"error"`
	Status  int             `json:"status"`
	Cause   [][]interface{} `json:"errors"`
}

func (e Error) Error() string {
	return e.Message
}

func (e Error) Exists() bool {
	return len(e.Message) > 0 || len(e.Errors) > 0
}
