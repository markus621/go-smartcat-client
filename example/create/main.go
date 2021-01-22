package main

import (
	"fmt"
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

	project, err := client.CreateProject(cli.NewProject{
		Name:                     "TS-10",
		Description:              "Перевод с русского на английский",
		SourceLanguage:           "en-US",
		TargetLanguages:          []string{"ru", "id"},
		AssignToVendor:           false,
		UseMT:                    false,
		Pretranslate:             false,
		UseTranslationMemory:     false,
		AutoPropagateRepetitions: false,
		WorkflowStages:           []string{"translation"},
		IsForTesting:             false,
	})
	if err != nil {
		panic(err)
	}

	form := cli.NewForm()
	form.AddFile("base1.json", []byte(`{"main":"hello world"}`))
	form.AddFile("base2.json", []byte(`{"main2":"hello my world"}`))

	docs, err := client.CreateDocument(project.ID, form)
	if err != nil {
		panic(err)
	}

	fmt.Println("Create new docs")
	for _, doc := range docs {
		fmt.Println(doc.Status, doc.Name)
		for _, ws := range doc.WorkflowStages {
			fmt.Println(ws.Status, ws.Progress)
		}
	}

	fmt.Println("Get status all docs")
	project, err = client.GetProject(project.ID)
	if err != nil {
		panic(err)
	}

	for _, doc := range project.Documents {
		fmt.Println(doc.Status, doc.Name)
		for _, ws := range doc.WorkflowStages {
			fmt.Println(ws.Status, ws.Progress)
		}
	}

}
