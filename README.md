# go-smartcat-client
Unofficial golang client for smartcat.com

[![Coverage Status](https://coveralls.io/repos/github/markus621/go-smartcat-client/badge.svg?branch=master)](https://coveralls.io/github/markus621/go-smartcat-client?branch=master)
[![Release](https://img.shields.io/github/release/markus621/go-smartcat-client.svg?style=flat-square)](https://github.com/markus621/go-smartcat-client/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/markus621/go-smartcat-client)](https://goreportcard.com/report/github.com/markus621/go-smartcat-client)
[![Build Status](https://travis-ci.com/markus621/go-smartcat-client.svg?branch=master)](https://travis-ci.com/markus621/go-smartcat-client)

# Example

```go
package main

import (
	"fmt"
	"os"

	cli "github.com/markus621/go-smartcat-client"
)

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
```