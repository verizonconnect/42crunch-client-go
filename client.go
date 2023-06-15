package crunchclient

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultTimeout   = 10 * time.Second
	DefaultUserAgent = "github.com/verizonconnect/42crunch-client-go"
)

type contextKey string

type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	userAgent  string
	debug      bool

	Collections CollectionsService
	API         ApiService
}

func NewClient(baseURL string, options ...ClientOption) (*Client, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("no api base url provided")
	}

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}

	client := Client{
		baseURL: u,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
		userAgent: DefaultUserAgent,
		debug:     false,
	}

	for _, option := range options {
		if optionErr := option(&client); optionErr != nil {
			return nil, optionErr
		}
	}

	client.Collections = CollectionsService{client: &client}
	client.API = ApiService{client: &client}

	return &client, nil
}

func (c Client) newRequest(ctx context.Context, method, path string, options ...requestOption) (*http.Request, error) {
	u, err := c.baseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.userAgent)

	for _, option := range options {
		if err = option(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}

type requestOption func(*http.Request) error

func (c Client) doRequest(req *http.Request, v interface{}) (resp apiResponse, err error) {
	if c.debug {
		reqDump, _ := httputil.DumpRequestOut(req, true)
		log.Printf("sending request:\n>>>>>>\n%s\n>>>>>>\n", string(reqDump))
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if c.debug {
		resDump, _ := httputil.DumpResponse(res, true)
		log.Printf("received response:\n<<<<<<\n%s\n<<<<<<\n", string(resDump))
	}

	err = checkResponseForError(res)
	if err != nil {
		return
	}

	if v != nil {
		switch vt := v.(type) {
		case *string:
			if content, readErr := io.ReadAll(res.Body); readErr == nil {
				*vt = strings.TrimSpace(string(content))
			} else {
				err = readErr
				return
			}
		default:
			err = json.NewDecoder(res.Body).Decode(v)
			if err != nil {
				return
			}
		}
	}

	resp, err = c.newAPIResponse(res)

	return resp, err
}

type apiResponse struct {
	*http.Response
	TotalCount int
}

func (c Client) newAPIResponse(res *http.Response) (a apiResponse, err error) {
	a = apiResponse{Response: res}

	totalCount, ok := a.Header["X-Total-Count"]
	if ok && len(totalCount) > 0 {
		totalCountVal, convErr := strconv.Atoi(totalCount[0])
		if convErr != nil {
			err = convErr
			return
		}
		a.TotalCount = totalCountVal
	}

	return
}

type ClientOption func(*Client) error
