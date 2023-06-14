package crunchclient

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func (apis ApiService) ListApis(ctx context.Context, id string) (c ApiResult, err error) {
	req, err := apis.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/collections/%s/apis", id))

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
