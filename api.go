package crunchclient

import (
	"context"
	"fmt"
	"net/http"
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
	Description struct {
		Id                 string `json:"id"`
		CollectionId       string `json:"cid"`
		Name               string `json:"name"`
		TechnicalName      string `json:"technicalName"`
		Specfile           string `json:"specfile"`
		Yaml               bool   `json:"yaml"`
		RevisionOasCounter int32  `json:"revisionOasCounter"`
		Lock               bool   `json:"lock"`
		LockReason         string `json:"lockReason"`
	} `json:"desc"`
	Revision   ApiRevision   `json:"revision"`
	Assessment ApiAssessment `json:"assessment"`
	Scan       ApiScan       `json:"scan"`
	Tags       []ApiTag      `json:"tag"`
}

type ApiAssessment struct {
	IsProcessed  bool      `json:"isProcessed"`
	Last         time.Time `json:"last"`
	Error        bool      `json:"error"`
	IsValid      bool      `json:"isValid"`
	Grade        float64   `json:"grade"`
	NumErrors    int       `json:"numErrors"`
	NumInfos     int       `json:"numInfos"`
	NumLows      int       `json:"numLows"`
	NumMediums   int       `json:"numMediums"`
	NumHighs     int       `json:"numHighs"`
	NumCriticals int       `json:"numCriticals"`
	OasVersion   string    `json:"oasVersion"`
	Releasable   bool      `json:"releasable"`
	SqgPass      bool      `json:"sqgPass"`
	AuditVersion string    `json:"auditVersion"`
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

func (apis ApiService) GetByCollection(ctx context.Context, id string) (c ApiResult, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v2/collections/%s/apis", id))
	if err != nil {
		return
	}

	_, err = apis.client.doRequest(req, &c)

	return
}

func (apis ApiService) Get(ctx context.Context, id string) (c ApiItem, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v2/apis/%s", id))
	if err != nil {
		return
	}

	_, err = apis.client.doRequest(req, &c)

	return
}
