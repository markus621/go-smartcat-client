package smartcatclient

import "github.com/pkg/errors"

//go:generate easyjson

var (
	//ErrUnknown unknown error
	ErrUnknown = errors.New("unknown error")
)

//ErrorResponse model error response from the server
//easyjson:json
type ErrorResponse struct {
	Message string `json:"Message"`
}

//Error error interface
func (v ErrorResponse) Error() string {
	return v.Message
}
