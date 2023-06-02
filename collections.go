package crunchclient

import (
	"context"
	"fmt"
	"net/http"
)

type CollectionsResult struct {
	Count int32        `json:"num"`
	Items []Collection `json:"list"`
}

type Collection struct {
	Description struct {
		Id            string `json:"id"`
		Name          string `json:"name"`
		IsShared      bool   `json:"isShared"`
		TechnicalName string `json:"technicalName"`
	} `json:"desc"`
	Summary struct {
		Org struct {
			Name string `json:"name"`
		} `json:"org"`
		Read     bool  `json:"read"`
		Write    bool  `json:"write"`
		ApiCount int32 `json:"apis"`
	} `json:"summary"`
	Owner struct {
		Id        string `json:"id"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
	} `json:"owner"`
}

type CollectionsService struct {
	client *Client
}

func (cs CollectionsService) GetAll(ctx context.Context) (c CollectionsResult, err error) {
	req, err := cs.client.newRequest(ctx, http.MethodGet, "/api/v1/collections")
	if err != nil {
		return
	}

	_, err = cs.client.doRequest(req, &c)

	return
}

func (cs CollectionsService) Get(ctx context.Context, id string) (c Collection, err error) {
	req, err := cs.client.newRequest(ctx, http.MethodGet, fmt.Sprintf("/api/v1/collections/%s", id))
	if err != nil {
		return
	}

	_, err = cs.client.doRequest(req, &c)

	return
}
