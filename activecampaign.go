package activecampaign

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sethgrid/pester"
)

var (
	ErrNoURLProvided            = errors.New("please provide your api url")
	ErrNoAuthenticationProvided = errors.New("please provide an authentication method")
)

// ActiveCampaign will be the main
type ActiveCampaign struct {
	Client *pester.Client

	url    string
	apiKey string
	output string
}

func New(url, apiKey string) (*ActiveCampaign, error) {
	if url == "" {
		return nil, ErrNoURLProvided
	}

	if strings.HasSuffix(url, "/") {
		url = strings.TrimSuffix(url, "")
	}

	client := pester.New()
	client.MaxRetries = 10
	client.Backoff = pester.LinearBackoff
	client.RetryOnHTTP429 = true

	ac := ActiveCampaign{
		Client: client,
		output: "json",
		url:    url,
	}

	if apiKey != "" {
		ac.apiKey = apiKey
	} else {
		return nil, ErrNoAuthenticationProvided
	}

	return &ac, nil
}

type POF struct {
	Pagination *Pagination
	Ordering   []Ordering
	Filtering  []Filtering
}
type Pagination struct {
	Limit  int
	Offset int
}
type Ordering struct {
	Key   string
	Order string
}
type Filtering struct {
	Key   string
	Value string
}

func (a *ActiveCampaign) send(ctx context.Context, method, api string, pof *POF, body io.Reader) (*http.Response, error) {
	u, err := url.Parse(a.url + "/api/3/" + api)
	if err != nil {
		return nil, &Error{Op: "send", Err: err}
	}

	if pof != nil {
		query := u.Query()
		if pof.Pagination != nil {
			query.Set("limit", strconv.Itoa(pof.Pagination.Limit))
			query.Set("offset", strconv.Itoa(pof.Pagination.Offset))
		}
		for _, v := range pof.Ordering {
			query.Add(fmt.Sprintf("orders[%s]", v.Key), v.Order)
		}
		for _, v := range pof.Filtering {
			query.Add(fmt.Sprintf("filters[%s]", v.Key), v.Value)
		}
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), body)
	if err != nil {
		return nil, &Error{Op: "send", Err: err}
	}
	req.Header.Set("Api-Token", a.apiKey)

	//b, _ := httputil.DumpRequest(req, true)
	//fmt.Println(string(b))
	res, err := a.Client.Do(req)
	if err != nil {
		return nil, &Error{Op: "send", Err: err}
	}
	// b, _ = httputil.DumpResponse(res, true)
	// fmt.Println(string(b))

	return res, nil
}

func (a *ActiveCampaign) CredentialsTest() bool {
	return true
}
