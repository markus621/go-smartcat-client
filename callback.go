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
	//AdditionalHeader additional headers for transmission with a callback
	AdditionalHeader struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}
	//Callback callback settings
	Callback struct {
		URL               string             `json:"url"`
		AdditionalHeaders []AdditionalHeader `json:"additionalHeaders"`
	}
	//LastError information about callback errors
	LastError struct {
		Created   string   `json:"created"`
		URL       string   `json:"url"`
		Reason    string   `json:"reason"`
		Code      int      `json:"code"`
		Content   string   `json:"content"`
		SourceIds []string `json:"sourceIds"`
	}
	//LastErrors error list
	LastErrors []LastError
)

//DelCallback Resetting the configuration of notifications reception
func (c *Client) DelCallback() (err error) {
	_, err = c.json(http.MethodDelete, uriCallback, nil, nil)
	return
}

//GetCallback Reading configurations of notifications reception of the account
func (c *Client) GetCallback() (out Callback, err error) {
	_, err = c.json(http.MethodGet, uriCallback, nil, &out)
	return
}

//SetCallback Setting configurations of notifications reception of the account
func (c *Client) SetCallback(in Callback) (out Callback, err error) {
	_, err = c.json(http.MethodPost, uriCallback, &in, &out)
	return
}

//GetCallbackLastErrors Reading the recent sending errors
func (c *Client) GetCallbackLastErrors(limit int) (out LastErrors, err error) {
	_, err = c.json(http.MethodGet, fmt.Sprintf("%s?limit=%d", uriCallbackLastErrors, limit), nil, &out)
	return
}
