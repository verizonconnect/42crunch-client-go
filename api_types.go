package crunchclient

import (
	"time"
)

type ApiService struct {
	client *Client
}

type ApiResult struct {
	Count int32     `json:"num"`
	Items []ApiItem `json:"list"`
}

type ApiItem struct {
	Description ApiDescription `json:"desc"`
	Revision    ApiRevision    `json:"revision"`
	Assessment  ApiAssessment  `json:"assessment"`
	Scan        ApiScan        `json:"scan"`
	Tags        []ApiTag       `json:"tags"`
}

type ApiDescription struct {
	Id                 string `json:"id"`
	CollectionId       string `json:"cid"`
	Name               string `json:"name"`
	TechnicalName      string `json:"technicalName"`
	Specfile           string `json:"specfile"`
	Yaml               bool   `json:"yaml"`
	RevisionOasCounter int32  `json:"revisionOasCounter"`
	Lock               bool   `json:"lock"`
	LockReason         string `json:"lockReason"`
}

type ApiAssessment struct {
	IsProcessed  bool    `json:"isProcessed"`
	Last         string  `json:"last"` // todo: this should ideally be time.Time, but it sometimes comes back nil, so need to check that.
	Error        bool    `json:"error"`
	IsValid      bool    `json:"isValid"`
	Grade        float64 `json:"grade"`
	NumErrors    int     `json:"numErrors"`
	NumInfos     int     `json:"numInfos"`
	NumLows      int     `json:"numLows"`
	NumMediums   int     `json:"numMediums"`
	NumHighs     int     `json:"numHighs"`
	NumCriticals int     `json:"numCriticals"`
	OasVersion   string  `json:"oasVersion"`
	Releasable   bool    `json:"releasable"`
	SqgPass      bool    `json:"sqgPass"`
	AuditVersion string  `json:"auditVersion"`
}

type ApiScan struct {
	IsProcessed bool        `json:"isProcessed"`
	Last        string      `json:"last"`
	NumHighs    int         `json:"numHighs"`
	NumMediums  int         `json:"numMediums"`
	NumLows     int         `json:"numLows"`
	State       string      `json:"state"`
	ExitCode    int         `json:"exitCode"`
	RequestDone int         `json:"requestDone"`
	TotalIssues int         `json:"totalIssues"`
	Mode        int         `json:"mode"`
	SqgPass     interface{} `json:"sqgPass"`
	ScanVersion string      `json:"scanVersion"`
}

type ApiRevision struct {
	ID                        string    `json:"id"`
	Aid                       string    `json:"aid"`
	CreateAt                  time.Time `json:"createAt"`
	TaskID                    string    `json:"taskId"`
	RevisionNumber            string    `json:"revisionNumber"`
	RevisionVersion           string    `json:"revisionVersion"`
	RevisionDate              string    `json:"revisionDate"`
	SecuredRevisionOasCounter string    `json:"SecuredRevisionOasCounter"`
	ParentID                  string    `json:"parentId"`
	Yaml                      bool      `json:"yaml"`
	OasFile                   string    `json:"oasFile"`
	ReadSpecFile              bool      `json:"readSpecFile"`
}

type ApiTag struct {
	CategoryID          string      `json:"categoryId"`
	CategoryName        string      `json:"categoryName"`
	CategoryDescription string      `json:"categoryDescription"`
	TagID               string      `json:"tagId"`
	TagName             string      `json:"tagName"`
	TagDescription      string      `json:"tagDescription"`
	Color               string      `json:"color"`
	IsProtected         bool        `json:"isProtected"`
	IsFreeForm          bool        `json:"isFreeForm"`
	IsExclusive         bool        `json:"isExclusive"`
	Dependencies        interface{} `json:"dependencies"`
}

type ApiStatus struct {
	LastAssessment        time.Time
	LastScan              time.Time
	IsAssessmentProcessed bool
	IsScanProcessed       bool
}

type ApiAssessmentResponse struct {
	Tid      string    `json:"tid"`
	Aid      string    `json:"aid"`
	Cid      string    `json:"cid"`
	Date     time.Time `json:"date"`
	Data     string    `json:"data"`
	Type     string    `json:"type"`
	Encoding string    `json:"enc"`
	Log      string    `json:"log"`
}

type AssessmentReport struct {
	Index                   []string `json:"index"`
	AssessmentVersion       string   `json:"assessmentVersion"`
	AssessmentReportVersion string   `json:"assessmentReportVersion"`
	Commit                  string   `json:"commit"`
	OasVersion              string   `json:"oasVersion"`
	APIVersion              string   `json:"apiVersion"`
	FileID                  string   `json:"fileId"`
	APIID                   string   `json:"apiId"`
	OpenapiState            string   `json:"openapiState"`
	Score                   float64  `json:"score"`
	Valid                   bool     `json:"valid"`
	Criticality             int      `json:"criticality"`
	IssueCounter            int      `json:"issueCounter"`
	SemanticErrors          struct {
		Issues map[string]AssessmentIssue `json:"issues"`
	} `json:"semanticErrors"`
	ValidationErrors struct {
		Issues map[string]AssessmentIssue `json:"issues"`
	} `json:"validationErrors"`
	Warnings struct {
		Issues map[string]AssessmentIssue `json:"issues"`
	} `json:"warnings"`
	OperationsNoAuthentication []int `json:"operationsNoAuthentication"`
	MinimalReport              bool  `json:"minimalReport"`
	MaxEntriesPerIssue         int   `json:"maxEntriesPerIssue"`
	MaxImpactedPerEntry        int   `json:"maxImpactedPerEntry"`
	Security                   struct {
		IssueCounter         int                        `json:"issueCounter"`
		Score                int                        `json:"score"`
		Criticality          int                        `json:"criticality"`
		Issues               map[string]AssessmentIssue `json:"issues"`
		SubgroupIssueCounter struct {
			Authentication IssueCounter `json:"authentication"`
			Authorization  IssueCounter `json:"authorization"`
			Transport      IssueCounter `json:"transport"`
		} `json:"subgroupIssueCounter"`
	} `json:"security"`
	Data struct {
		IssueCounter         int                        `json:"issueCounter"`
		Score                float64                    `json:"score"`
		Criticality          int                        `json:"criticality"`
		Issues               map[string]AssessmentIssue `json:"issues"`
		SubgroupIssueCounter struct {
			Parameters         IssueCounter `json:"parameters"`
			ResponseHeader     IssueCounter `json:"responseHeader"`
			ResponseDefinition IssueCounter `json:"responseDefinition"`
			Schema             IssueCounter `json:"schema"`
			Paths              IssueCounter `json:"paths"`
		} `json:"subgroupIssueCounter"`
	}
	IssuesKey     []string                `json:"issuesKey"`
	SkippedIssues []string                `json:"skippedIssues"`
	Summary       AssessmentReportSummary `json:"summary"`
}

type AssessmentIssue struct {
	Description string `json:"description"`
	Issues      []struct {
		Score               float64 `json:"score"`
		Pointer             int     `json:"pointer"`
		TooManyImpacted     bool    `json:"tooManyImpacted"`
		Criticality         int     `json:"criticality"`
		Response            bool    `json:"response"`
		SpecificDescription string  `json:"specificDescription"`
	} `json:"issues"`
	TotalIssues  int     `json:"totalIssues"`
	IssueCounter int     `json:"issueCounter"`
	Score        float64 `json:"score"`
	Criticality  int     `json:"criticality"`
	TooManyError bool    `json:"tooManyError"`
}

type IssueCounter struct {
	None     int `json:"none"`
	Info     int `json:"info"`
	Low      int `json:"low"`
	Medium   int `json:"medium"`
	High     int `json:"high"`
	Critical int `json:"critical"`
}

type AssessmentReportSummary struct {
	OasVersion                       string           `json:"oasVersion"`
	APIVersion                       string           `json:"apiVersion"`
	Basepath                         string           `json:"basepath"`
	APIName                          string           `json:"apiName"`
	Description                      string           `json:"description"`
	Endpoints                        []string         `json:"endpoints"`
	PathCounter                      int              `json:"pathCounter"`
	OperationCounter                 int              `json:"operationCounter"`
	ParameterCounter                 int              `json:"parameterCounter"`
	RequestBodyCounter               int              `json:"requestBodyCounter"`
	RequestContentType               map[string]int64 `json:"requestContentType"`
	ResponseContentType              map[string]int64 `json:"responseContentType"`
	ComponentsSchemasCounter         int              `json:"componentsSchemasCounter"`
	ComponentsResponsesCounter       int              `json:"componentsResponsesCounter"`
	ComponentsParametersCounter      int              `json:"componentsParametersCounter"`
	ComponentsExamplesCounter        int              `json:"componentsExamplesCounter"`
	ComponentsRequestBodiesCounter   int              `json:"componentsRequestBodiesCounter"`
	ComponentsHeadersCounter         int              `json:"componentsHeadersCounter"`
	ComponentsSecuritySchemesCounter int              `json:"componentsSecuritySchemesCounter"`
	ComponentsLinksCounter           int              `json:"componentsLinksCounter"`
	ComponentsCallbacksCounter       int              `json:"componentsCallbacksCounter"`
}

type ApiIssue struct {
	Id           string
	Description  string
	Pointer      string
	Score        int32
	DisplayScore string
	Criticality  int
	Severity     string
}
