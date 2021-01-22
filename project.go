package smartcatclient

import (
	"fmt"
	"net/http"
	"time"
)

//go:generate easyjson

const (
	uriProject                    = "/api/integration/v1/project/%s"
	uriProjectList                = "/api/integration/v1/project/list"
	uriProjectCancel              = "/api/integration/v1/project/cancel"
	uriProjectRestore             = "/api/integration/v1/project/restore"
	uriProjectComplete            = "/api/integration/v1/project/complete"
	uriProjectCreate              = "/api/integration/v1/project/create"
	uriProjectStatistics          = "/api/integration/v2/project/%s/statistics"
	uriProjectCWStatistics        = "/api/integration/v2/project/%s/completedWorkStatistics"
	uriProjectTranslationMemories = "/api/integration/v2/project/%s/translationmemories"
)

//easyjson:json
type (
	//ProjectsList model
	ProjectsList []Project
	//Project model
	Project struct {
		ID                     string          `json:"id"`
		AccountID              string          `json:"accountId"`
		Name                   string          `json:"name"`
		Description            string          `json:"description"`
		Deadline               time.Time       `json:"deadline"`
		CreationDate           time.Time       `json:"creationDate"`
		CreatedByUserID        string          `json:"createdByUserId"`
		CreatedByUserEmail     string          `json:"createdByUserEmail"`
		ModificationDate       time.Time       `json:"modificationDate"`
		SourceLanguageID       uint64          `json:"sourceLanguageId"`
		SourceLanguage         string          `json:"sourceLanguage"`
		TargetLanguages        []string        `json:"targetLanguages"`
		Status                 string          `json:"status"`
		StatusModificationDate time.Time       `json:"statusModificationDate"`
		DomainID               uint64          `json:"domainId"`
		ClientID               uint64          `json:"clientId"`
		Vendors                []Vendor        `json:"vendors"`
		WorkflowStages         []WorkflowStage `json:"workflowStages"`
		Documents              []Document      `json:"documents"`
		ExternalTag            string          `json:"externalTag"`
		Specializations        []string        `json:"specializations"`
		Managers               []string        `json:"managers"`
		Number                 []string        `json:"number"`
	}
	//NewProject model
	NewProject struct {
		Name                     string   `json:"name"`
		Description              string   `json:"description"`
		SourceLanguage           string   `json:"sourceLanguage"`
		TargetLanguages          []string `json:"targetLanguages"`
		AssignToVendor           bool     `json:"assignToVendor"`
		UseMT                    bool     `json:"useMT"`
		Pretranslate             bool     `json:"pretranslate"`
		UseTranslationMemory     bool     `json:"useTranslationMemory"`
		AutoPropagateRepetitions bool     `json:"autoPropagateRepetitions"`
		WorkflowStages           []string `json:"workflowStages"`
		IsForTesting             bool     `json:"isForTesting"`
	}
	//Cost model
	Cost struct {
		Value           float64 `json:"value"`
		Currency        string  `json:"currency"`
		AccuracyDegree  string  `json:"accuracyDegree"`
		DetailsFileName string  `json:"detailsFileName"`
		PaymentStatus   string  `json:"paymentStatus"`
	}
	//Vendor model
	Vendor struct {
		VendorAccountID    string `json:"vendorAccountId"`
		RemovedFromProject bool   `json:"removedFromProject"`
		Cost               Cost   `json:"cost"`
		CostDetailsFileID  string `json:"costDetailsFileId"`
	}
	//WorkflowStage model
	WorkflowStage struct {
		Progress  float64 `json:"progress"`
		StageType string  `json:"stageType"`
	}
	//Executive model
	Executive struct {
		ID                 string `json:"id"`
		AssignedWordsCount uint64 `json:"assignedWordsCount"`
		Progress           uint   `json:"progress"`
		SupplierType       string `json:"supplierType"`
	}
	//PatchProject model
	PatchProject struct {
		Name             string    `json:"name"`
		Description      string    `json:"description"`
		Deadline         time.Time `json:"deadline"`
		ClientID         string    `json:"clientId"`
		DomainID         uint64    `json:"domainId"`
		VendorAccountIDs []string  `json:"vendorAccountIds"`
		ExternalTag      string    `json:"externalTag"`
		Specializations  []string  `json:"specializations"`
		WorkflowStages   []string  `json:"workflowStages"`
	}
	//StatisticsList model
	StatisticsList []Statistics
	//Statistics model
	Statistics struct {
		Language   string                    `json:"language"`
		Statistics []StatisticsItem          `json:"statistics"`
		Documents  []StatisticsDocumentsItem `json:"documents"`
	}
	//StatisticsItem model
	StatisticsItem struct {
		Name                     string  `json:"name"`
		Words                    uint64  `json:"words"`
		Percent                  uint    `json:"percent"`
		Segments                 uint16  `json:"segments"`
		Pages                    float64 `json:"pages"`
		CharsWithoutSpaces       uint64  `json:"charsWithoutSpaces"`
		CharsWithSpaces          uint64  `json:"charsWithSpaces"`
		EffectiveWordsForBilling uint64  `json:"effectiveWordsForBilling"`
	}
	//StatisticsDocumentsItem model
	StatisticsDocumentsItem struct {
		Name       string           `json:"name"`
		Statistics []StatisticsItem `json:"statistics"`
	}
	//CompletedWorkStatisticsList model
	CompletedWorkStatisticsList []CompletedWorkStatistics
	//CompletedWorkStatistics model
	CompletedWorkStatistics struct {
		Executive struct {
			ID           string `json:"id"`
			SupplierType string `json:"supplierType"`
		} `json:"executive"`
		StageType      string                    `json:"stageType"`
		StageNumber    int                       `json:"stageNumber"`
		TargetLanguage string                    `json:"targetLanguage"`
		Total          []StatisticsItem          `json:"total"`
		Documents      []StatisticsDocumentsItem `json:"documents"`
	}
	//TranslationMemories model
	TranslationMemories []struct {
		ID             string `json:"id"`
		MatchThreshold int    `json:"matchThreshold"`
		IsWritable     bool   `json:"isWritable"`
	}
)

//CreateProject Create new project
func (c *Client) CreateProject(in NewProject) (out Project, err error) {
	form := NewForm()
	if err = form.AddJSON("model", &in); err != nil {
		return
	}
	_, err = c.form(http.MethodPost, uriProjectCreate, form, &out)
	return
}

//DelProject Delete the project
func (c *Client) DelProject(id string) (err error) {
	_, err = c.json(http.MethodDelete, fmt.Sprintf(uriProject, id), nil, nil)
	return
}

//GetProject Receive the project model
func (c *Client) GetProject(id string) (out Project, err error) {
	_, err = c.json(http.MethodGet, fmt.Sprintf(uriProject, id), nil, &out)
	return
}

//SetProject Change the project model
func (c *Client) SetProject(id string, in PatchProject) (out Project, err error) {
	_, err = c.json(http.MethodPut, fmt.Sprintf(uriProject, id), &in, &out)
	return
}

//CancelProject Cancel the project
func (c *Client) CancelProject(id string) (err error) {
	_, err = c.json(http.MethodPost, uriProjectCancel+"?projectId="+id, nil, nil)
	return
}

//RestoreProject Restore the project
func (c *Client) RestoreProject(id string) (err error) {
	_, err = c.json(http.MethodPost, uriProjectRestore+"?projectId="+id, nil, nil)
	return
}

//CompleteProject Complete the project
func (c *Client) CompleteProject(id string) (err error) {
	_, err = c.json(http.MethodPost, uriProjectComplete+"?projectId="+id, nil, nil)
	return
}

//ListProject List all projects
func (c *Client) ListProject() (out ProjectsList, err error) {
	_, err = c.json(http.MethodGet, uriProjectList, nil, &out)
	return
}

//GetProjectStatistics Receive statistics
func (c *Client) GetProjectStatistics(id string) (out StatisticsList, err error) {
	_, err = c.json(http.MethodGet, uriProjectStatistics, nil, &out)
	return
}

//GetProjectCompletedWorkStatistics Receiving statistics for the completed parts of the project
func (c *Client) GetProjectCompletedWorkStatistics(id string) (out CompletedWorkStatisticsList, err error) {
	_, err = c.json(http.MethodGet, uriProjectCWStatistics, nil, &out)
	return
}

//GetProjectTranslationMemories Receiving a list of the TMs plugged into the project
func (c *Client) GetProjectTranslationMemories(id string) (out TranslationMemories, err error) {
	_, err = c.json(http.MethodGet, uriProjectTranslationMemories, nil, &out)
	return
}

//SetProjectTranslationMemories Receiving a list of the TMs plugged into the project
func (c *Client) SetProjectTranslationMemories(id string, in TranslationMemories) (err error) {
	_, err = c.json(http.MethodGet, uriProjectTranslationMemories, &in, nil)
	return
}
