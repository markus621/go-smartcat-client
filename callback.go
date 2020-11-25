package smartcatclient

import (
	"fmt"
	"net/http"
)

//go:generate easyjson

const (
	uriCallback           = "/api/integration/v1/callback"
	uriCallbackLastErrors = "/api/integration/v1/callback/lastErrors"
)

//easyjson:json
type (
	AdditionalHeader struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	Callback struct {
		URL               string             `json:"url"`
		AdditionalHeaders []AdditionalHeader `json:"additionalHeaders"`
	}
	LastError struct {
		Created   string   `json:"created"`
		URL       string   `json:"url"`
		Reason    string   `json:"reason"`
		Code      int      `json:"code"`
		Content   string   `json:"content"`
		SourceIds []string `json:"sourceIds"`
	}
	LastErrors []LastError
)

//DelCallback Resetting the configuration of notifications reception
func (v *Client) DelCallback() (err error) {
	err, _ = v.call(http.MethodDelete, uriCallback, nil, nil)
	return
}

//GetCallback Reading configurations of notifications reception of the account
func (v *Client) GetCallback() (out Callback, err error) {
	err, _ = v.call(http.MethodGet, uriCallback, nil, &out)
	return
}

//SetCallback Setting configurations of notifications reception of the account
func (v *Client) SetCallback(in Callback) (out Callback, err error) {
	err, _ = v.call(http.MethodPost, uriCallback, &in, &out)
	return
}

//GetCallbackLastErrors Reading the recent sending errors
func (v *Client) GetCallbackLastErrors(limit int) (out LastErrors, err error) {
	err, _ = v.call(http.MethodGet, fmt.Sprintf("%s?limit=%d", uriCallbackLastErrors, limit), nil, &out)
	return
}
