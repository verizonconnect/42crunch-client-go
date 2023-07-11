package crunchclient

import (
	"context"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"
)

// return all the api's for a specific api collection
func (apis ApiService) ListApis(ctx context.Context, id string) (c ApiResult, err error) {
	capis, err := apis.ListApisPaged(ctx, id, 1, 50)
	if err != nil {
		return
	}

	apiItems := make([]ApiItem, 0)
	apiItems = append(apiItems, capis.Items...)
	if capis.Count > 50 {
		pages := math.Ceil(float64(capis.Count) / 50)

		// we will start at page 2 here given we already got page one
		for i := 2; i < int(pages)+1; i++ {
			pagedResult, perr := apis.ListApisPaged(ctx, id, i, 50)
			if perr != nil {
				err = fmt.Errorf("unable to read page of while retrieving api's. %w", err)
				return
			}

			apiItems = append(apiItems, pagedResult.Items...)
		}
	}

	var result ApiResult
	result.Count = int32(len(apiItems))
	result.Items = apiItems

	return result, nil
}

// return paged results of api's for a specific api collection
func (apis ApiService) ListApisPaged(ctx context.Context, id string, page int, perPage int) (c ApiResult, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v2/collections/%s/apis?withTags=true&page=%d&perPage=%d", id, page, perPage))

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
