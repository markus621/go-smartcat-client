package smartcatclient

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

//go:generate easyjson

const (
	uriDocumentCreate     = "/api/integration/v1/project/document"
	uriDocumentExport     = "/api/integration/v1/document/export"
	uriDocumentTaskExport = "/api/integration/v1/document/export/%s"
)

//easyjson:json
type (
	//DocumentWorkflowStage model
	DocumentWorkflowStage struct {
		Progress             float64     `json:"progress"`
		WordsTranslated      uint64      `json:"wordsTranslated"`
		UnassignedWordsCount uint64      `json:"unassignedWordsCount"`
		Status               string      `json:"status"`
		Executives           []Executive `json:"executives"`
	}
	//Document model
	Document struct {
		ID                     string                  `json:"id"`
		Name                   string                  `json:"name"`
		CreationDate           time.Time               `json:"creationDate"`
		Deadline               time.Time               `json:"deadline"`
		SourceLanguage         string                  `json:"sourceLanguage"`
		DisassemblingStatus    string                  `json:"documentDisassemblingStatus"`
		TargetLanguage         string                  `json:"targetLanguage"`
		Status                 string                  `json:"status"`
		WordsCount             uint64                  `json:"wordsCount"`
		StatusModificationDate time.Time               `json:"statusModificationDate"`
		PretranslateCompleted  bool                    `json:"pretranslateCompleted"`
		WorkflowStages         []DocumentWorkflowStage `json:"workflowStages"`
		ExternalID             string                  `json:"externalId"`
		MetaInfo               string                  `json:"metaInfo"`
		PlaceholdersAreEnabled bool                    `json:"placeholdersAreEnabled"`
	}
	//DocumentList model
	DocumentList []Document
	//ExportTask model
	ExportTask struct {
		ID          string   `json:"id"`
		DocumentIds []string `json:"documentIds"`
	}
)

//CreateDocument Create new document in project
func (c *Client) CreateDocument(projectID string, form *Form) (out DocumentList, err error) {
	_, err = c.form(http.MethodPost, uriDocumentCreate+"?projectId="+projectID, form, &out)
	return
}

//ExportDocument Creating a document export task
func (c *Client) ExportDocument(ids []string) (out ExportTask, err error) {
	v := url.Values{"documentIds": ids}
	u := url.URL{Path: uriDocumentExport, RawQuery: v.Encode()}
	_, err = c.json(http.MethodPost, u.String(), nil, &out)
	return
}

//ExportDocumentByTaskID Downloading files by task number for export
func (c *Client) ExportDocumentByTaskID(id string) (out []byte, err error) {
	_, out, err = c.raw(http.MethodGet, fmt.Sprintf(uriDocumentTaskExport, id), nil)
	return
}
