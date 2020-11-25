package smartcatclient

import "github.com/pkg/errors"

//go:generate easyjson

var (
	ErrUnknown = errors.New("unknown error")
)

//ErrorResponse ...
//easyjson:json
type ErrorResponse struct {
	Message string `json:"Message"`
}

func (v ErrorResponse) Error() string {
	return v.Message
}
