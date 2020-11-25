package smartcatclient

import "encoding/base64"

const (
	//HostURL  if you are using the European server.
	HostURL = `https://smartcat.ai`
	//USHostURL  if you are using the American server.
	USHostURL = `https://us.smartcat.ai`
	//EAHostURL  if you are using the Asian server.
	EAHostURL = `https://ea.smartcat.ai`
)

type Config struct {
	AccountID string
	AuthKey   string
	URL       string
}

func (c Config) AuthToken() string {
	return `Basic ` + base64.StdEncoding.EncodeToString([]byte(c.AccountID+`:`+c.AuthKey))
}
