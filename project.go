package smartcatclient

import "net/http"

//go:generate easyjson

const (
	uriProject     = "/api/integration/v1/project"
	uriProjectList = "/api/integration/v1/project/list"
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
		Deadline               string          `json:"deadline"`
		CreationDate           string          `json:"creationDate"`
		CreatedByUserID        string          `json:"createdByUserId"`
		CreatedByUserEmail     string          `json:"createdByUserEmail"`
		ModificationDate       string          `json:"modificationDate"`
		SourceLanguageID       uint64          `json:"sourceLanguageId"`
		SourceLanguage         string          `json:"sourceLanguage"`
		TargetLanguages        []string        `json:"targetLanguages"`
		Status                 string          `json:"status"`
		StatusModificationDate string          `json:"statusModificationDate"`
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
		CreationDate           string                  `json:"creationDate"`
		Deadline               string                  `json:"deadline"`
		SourceLanguage         string                  `json:"sourceLanguage"`
		DisassemblingStatus    string                  `json:"documentDisassemblingStatus"`
		TargetLanguage         string                  `json:"targetLanguage"`
		Status                 string                  `json:"status"`
		WordsCount             uint64                  `json:"wordsCount"`
		StatusModificationDate string                  `json:"statusModificationDate"`
		PretranslateCompleted  bool                    `json:"pretranslateCompleted"`
		WorkflowStages         []DocumentWorkflowStage `json:"workflowStages"`
		ExternalID             string                  `json:"externalId"`
		MetaInfo               string                  `json:"metaInfo"`
		PlaceholdersAreEnabled bool                    `json:"placeholdersAreEnabled"`
	}
	//PatchProject model
	PatchProject struct {
		Name             string   `json:"name"`
		Description      string   `json:"description"`
		Deadline         string   `json:"deadline"`
		ClientID         string   `json:"clientId"`
		DomainID         uint64   `json:"domainId"`
		VendorAccountIDs []string `json:"vendorAccountIds"`
		ExternalTag      string   `json:"externalTag"`
		Specializations  []string `json:"specializations"`
		WorkflowStages   []string `json:"workflowStages"`
	}
)

//DelProject Delete the project
func (v *Client) DelProject(id string) (err error) {
	_, err = v.call(http.MethodDelete, uriProject+"/"+id, nil, nil)
	return
}

//GetProject Receive the project model
func (v *Client) GetProject(id string) (out Project, err error) {
	_, err = v.call(http.MethodGet, uriProject+"/"+id, nil, &out)
	return
}

//SetProject Change the project model
func (v *Client) SetProject(id string, in PatchProject) (out Project, err error) {
	_, err = v.call(http.MethodPut, uriProject+"/"+id, &in, &out)
	return
}

//ListProject List all projects
func (v *Client) ListProject() (out ProjectsList, err error) {
	_, err = v.call(http.MethodGet, uriProjectList, nil, &out)
	return
}
