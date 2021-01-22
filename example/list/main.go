package main

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
