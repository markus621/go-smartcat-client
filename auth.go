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

//Config client settings for connecting to the server
type Config struct {
	AccountID string
	AuthKey   string
	URL       string
}

//AuthToken generating an authorization token
func (c Config) AuthToken() string {
	return `Basic ` + base64.StdEncoding.EncodeToString([]byte(c.AccountID+`:`+c.AuthKey))
}
