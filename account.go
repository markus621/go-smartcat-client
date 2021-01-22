package smartcatclient

import (
	"net/http"
)

//go:generate easyjson

const (
	uriAccount          = "/api/integration/v1/account"
	uriAccountMTEngines = "/api/integration/v1/account/mtengines"
)

//easyjson:json
type (
	//Account information about your current account.
	Account struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		IsPersonal bool   `json:"isPersonal"`
		Type       string `json:"type"`
	}
	//AccountMTEngine information about the machine translation engine
	AccountMTEngine struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	//AccountMTEngines list of available machine translation engines
	AccountMTEngines []AccountMTEngine
)

//GetAccount Receiving the account details
func (c *Client) GetAccount() (out Account, err error) {
	_, err = c.json(http.MethodGet, uriAccount, nil, &out)
	return
}

//GetAccountMTEngines Receiving MT engines available for the account
func (c *Client) GetAccountMTEngines() (out AccountMTEngines, err error) {
	_, err = c.json(http.MethodGet, uriAccountMTEngines, nil, &out)
	return
}
