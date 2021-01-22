# go-smartcat-client
Unofficial golang client for smartcat.com

[![Coverage Status](https://coveralls.io/repos/github/markus621/go-smartcat-client/badge.svg?branch=master)](https://coveralls.io/github/markus621/go-smartcat-client?branch=master)
[![Release](https://img.shields.io/github/release/markus621/go-smartcat-client.svg?style=flat-square)](https://github.com/markus621/go-smartcat-client/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/markus621/go-smartcat-client)](https://goreportcard.com/report/github.com/markus621/go-smartcat-client)
[![Build Status](https://travis-ci.com/markus621/go-smartcat-client.svg?branch=master)](https://travis-ci.com/markus621/go-smartcat-client)

## Official documentation

[Swagger 2.0](https://smartcat.com/api/swagger/docs/v1) + [Swagger UI](https://smartcat.com/api/methods/)

# Example

### _Create project and upload documents_

```go
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

```

### _Get all documents and export_
```go
import (
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"

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

	list, err := client.ListProject()
	if err != nil {
		panic(err)
	}

	ids := make([]string, 0)
	for _, project := range list {
		fmt.Println(project.ID, project.Name, project.Status)
		for _, doc := range project.Documents {
			fmt.Println(doc.ID, doc.Name, doc.Status, doc.SourceLanguage, doc.TargetLanguage, doc.WorkflowStages[0].Progress)

			if doc.WorkflowStages[0].Progress == 100 {
				ids = append(ids, doc.ID)
			}
		}
	}

	task, err := client.ExportDocument(ids)
	if err != nil {
		panic(err)
	}

	<-time.After(3 * time.Second)

	data, err := client.ExportDocumentByTaskID(task.ID)
	if err != nil {
		panic(err)
	}

	// as text file if one document
	if len(ids) == 1 {
		fmt.Println(string(data))
		os.Exit(0)
	}

	// as zip file if many documents
	z, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		panic(err)
	}
	for _, f := range z.File {
		reader, err := f.Open()
		if err != nil {
			panic(err)
		}
		b, err := ioutil.ReadAll(reader)
		fmt.Println(f.Name, string(b), err)
	}
}

```