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

	acc, err := client.Account()
	if err != nil {
		panic(err)
	}
	fmt.Println("==> ", acc)

	mte, err := client.AccountMTEngines()
	if err != nil {
		panic(err)
	}

	fmt.Println("==> ", mte)

}
