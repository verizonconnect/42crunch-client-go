package crunchclient

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (apis ApiService) ListApis(ctx context.Context, id string) (c ApiResult, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v2/collections/%s/apis?withTags=true", id))

	if err != nil {
		return
	}

	_, err = apis.client.doRequest(req, &c)

	return
}

func (apis ApiService) ReadApi(ctx context.Context, id string) (c ApiItem, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/apis/%s", id))
	if err != nil {
		return
	}

	_, err = apis.client.doRequest(req, &c)

	return
}

func (apis ApiService) ReadApiStatus(ctx context.Context, id string) (res ApiStatus, err error) {

	c, err := apis.ReadApi(ctx, id)
	if err != nil {
		return
	}

	// we don't care if the time conversion fail, the default will work
	lastAssessment, _ := time.Parse(time.RFC3339, c.Assessment.Last)
	lastScan, _ := time.Parse(time.RFC3339, c.Scan.Last)

	res = ApiStatus{
		LastAssessment:        lastAssessment,
		IsAssessmentProcessed: c.Assessment.IsProcessed,
		LastScan:              lastScan,
		IsScanProcessed:       c.Scan.IsProcessed,
	}

	return
}

func (apis ApiService) ReadAssessmentReport(ctx context.Context, id string) (report AssessmentReport, err error) {

	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/apis/%s/assessmentreport", id))
	if err != nil {
		return
	}

	// todo: we first need to ensure the api is processed before we can return this report

	var resp ApiAssessmentResponse
	_, err = apis.client.doRequest(req, &resp)

	if err != nil {
		return report, err
	}

	if resp.Encoding != "base64" {
		err = fmt.Errorf("unsupported data type")
		return report, err
	}

	sDec, err := b64.StdEncoding.DecodeString(resp.Data)
	if resp.Encoding != "base64" {
		err = fmt.Errorf("Unable to decode report document data: %w", err)
		return report, err
	}

	err = json.Unmarshal(sDec, &report)
	if err != nil {
		err = fmt.Errorf("Unable to decode report document: %w", err)
	}

	return report, err
}
