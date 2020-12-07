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
	uriProjectDocument            = "/api/integration/v1/project/document"
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
		Progress  uint   `json:"progress"`
		StageType string `json:"stageType"`
	}
	//DocumentWorkflowStage model
	DocumentWorkflowStage struct {
		Progress             uint        `json:"progress"`
		WordsTranslated      uint64      `json:"wordsTranslated"`
		UnassignedWordsCount uint64      `json:"unassignedWordsCount"`
		Status               string      `json:"status"`
		Executives           []Executive `json:"executives"`
	}
	//Executive model
	Executive struct {
		ID                 string `json:"id"`
		AssignedWordsCount uint64 `json:"assignedWordsCount"`
		Progress           uint   `json:"progress"`
		SupplierType       string `json:"supplierType"`
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

//DelProject Delete the project
func (v *Client) DelProject(id string) (err error) {
	_, err = v.call(http.MethodDelete, fmt.Sprintf(uriProject, id), nil, nil)
	return
}

//GetProject Receive the project model
func (v *Client) GetProject(id string) (out Project, err error) {
	_, err = v.call(http.MethodGet, fmt.Sprintf(uriProject, id), nil, &out)
	return
}

//SetProject Change the project model
func (v *Client) SetProject(id string, in PatchProject) (out Project, err error) {
	_, err = v.call(http.MethodPut, fmt.Sprintf(uriProject, id), &in, &out)
	return
}

//CancelProject Cancel the project
func (v *Client) CancelProject(id string) (err error) {
	_, err = v.call(http.MethodPost, uriProjectCancel+"?projectId="+id, nil, nil)
	return
}

//RestoreProject Restore the project
func (v *Client) RestoreProject(id string) (err error) {
	_, err = v.call(http.MethodPost, uriProjectRestore+"?projectId="+id, nil, nil)
	return
}

//CompleteProject Complete the project
func (v *Client) CompleteProject(id string) (err error) {
	_, err = v.call(http.MethodPost, uriProjectComplete+"?projectId="+id, nil, nil)
	return
}

//ListProject List all projects
func (v *Client) ListProject() (out ProjectsList, err error) {
	_, err = v.call(http.MethodGet, uriProjectList, nil, &out)
	return
}

//GetProjectStatistics Receive statistics
func (v *Client) GetProjectStatistics(id string) (out StatisticsList, err error) {
	_, err = v.call(http.MethodGet, uriProjectStatistics, nil, &out)
	return
}

//GetProjectCompletedWorkStatistics Receiving statistics for the completed parts of the project
func (v *Client) GetProjectCompletedWorkStatistics(id string) (out CompletedWorkStatisticsList, err error) {
	_, err = v.call(http.MethodGet, uriProjectCWStatistics, nil, &out)
	return
}

//GetProjectTranslationMemories Receiving a list of the TMs plugged into the project
func (v *Client) GetProjectTranslationMemories(id string) (out TranslationMemories, err error) {
	_, err = v.call(http.MethodGet, uriProjectTranslationMemories, nil, &out)
	return
}

//SetProjectTranslationMemories Receiving a list of the TMs plugged into the project
func (v *Client) SetProjectTranslationMemories(id string, in TranslationMemories) (err error) {
	_, err = v.call(http.MethodGet, uriProjectTranslationMemories, &in, nil)
	return
}
