package main

import (
	"os"

	cli "github.com/markus621/go-smartcat-client"
)

//nolint: errcheck
func main() {

	conf := cli.Config{
		AccountID: os.Getenv(`SMARTCAT_ACCOUNT_ID`),
		AuthKey:   os.Getenv(`SMARTCAT_AUTH_KEY`),
		URL:       cli.HostURL,
	}

	client := cli.NewClient(conf)
	client.Debug(true, os.Stdout)

	_, _ = client.GetAccount()
	_, _ = client.GetAccountMTEngines()
	_, _ = client.SetCallback(cli.Callback{
		URL: "https://demo.example/callback",
		AdditionalHeaders: []cli.AdditionalHeader{
			{Name: "x-header", Value: "demo"},
		},
	})
	_, _ = client.GetCallback()
	_ = client.DelCallback()
	_, _ = client.GetCallbackLastErrors(10)
}
