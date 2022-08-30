package go_itunes_api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	APIPath    = "https://itunes.apple.com"
	SearchPath = "/search"
	US         = "US"
)

const DefaultHTTPTimeout = 60 * time.Second

type SearchRequest struct {

	// The text string you want to search for. For example: jack+johnson.
	//
	// Required.
	Term string

	// The two-letter country code for the store you want to search. The search uses
	// the default store front for the specified country. For example: US. The
	// default is US. See http://en.wikipedia.org/wiki/ISO_3166-1_alpha-2 for a list
	// of ISO Country Codes.
	//
	// Required.
	Country string

	// The media type you want to search for. For example: movie. The default is all.
	//
	// Optional.
	Media MediaType

	// The type of results you want returned, relative to the specified media type.
	// For example: movieArtist for a movie media type search. The default is the
	// track entity associated with the specified media type.
	//
	// Optional.
	Entity Entity

	// The attribute you want to search for in the stores, relative to the specified
	// media type. For example, if you want to search for an artist by name specify
	// entity=allArtist&attribute=allArtistTerm. In this example, if you search for
	// term=maroon, iTunes returns “Maroon 5” in the search results, instead of all
	// artists who have ever recorded a song with the word “maroon” in the title. The
	// default is all attributes associated with the specified media type.
	//
	// Optional.
	Attribute Attribute

	// The name of the Javascript callback function you want to use when returning
	// search results to your website.
	//
	// Required for cross-site searches.
	Callback string

	// The number of search results you want the iTunes Store to return. For example:
	// 25. The default is 50.
	//
	// Optional.
	Limit int

	// The language, English or Japanese, you want to use when returning search
	// results. Specify the language using the five-letter codename. For example:
	// en_us. The default is en_us (English).
	//
	// Optional.
	Lang string

	// The search result key version you want to receive back from your search. The
	// default is 2.
	//
	// Optional.
	Version string

	// A flag indicating whether or not you want to include explicit content in your
	// search results. The default is Yes.
	//
	// Optional.
	Explicit bool
}

func (r *SearchRequest) Validate() error {
	if r.Term == "" {
		return fmt.Errorf("term is required")
	}
	if r.Country == "" {
		return fmt.Errorf("country is required")
	}
	return nil
}

func (r SearchRequest) GetURLValues() url.Values {
	v := url.Values{}
	v.Set("term", r.Term)
	v.Set("country", r.Country)
	if r.Media != "" {
		v.Set("media", string(r.Media))
	}
	if r.Entity != "" {
		v.Set("entity", string(r.Entity))
	}
	if r.Attribute != "" {
		v.Set("attribute", string(r.Attribute))
	}
	if r.Callback != "" {
		v.Set("callback", r.Callback)
	}
	if r.Limit != 0 {
		v.Set("limit", strconv.Itoa(r.Limit))
	}
	if r.Lang != "" {
		v.Set("lang", r.Lang)
	}
	if r.Version != "" {
		v.Set("version", r.Version)
	}
	if r.Explicit {
		v.Set("explicit", "Yes")
	} else {
		v.Set("explicit", "No")
	}
	return v
}

type SearchResponse struct {
	Results Results
}

type LookupRequest struct{}

type LookupResponse struct {
	Results Results
}

type API interface {
	Search(ctx context.Context, r SearchRequest) (SearchResponse, error)
	Lookup(ctx context.Context, r LookupRequest) (LookupResponse, error)
}

type ClientOption struct {
	Host    string
	Timeout time.Duration
}

type Client struct {
	client *http.Client
	host   string
}

func NewClient(opt ClientOption) (*Client, error) {
	if opt.Host == "" {
		opt.Host = APIPath
	}
	if opt.Timeout == 0 {
		opt.Timeout = DefaultHTTPTimeout
	}
	return &Client{
		client: &http.Client{
			Transport: new(http.Transport),
			Timeout:   opt.Timeout,
		},
		host: opt.Host,
	}, nil
}

func (c *Client) Search(ctx context.Context, r SearchRequest) (SearchResponse, error) {
	err := r.Validate()
	if err != nil {
		return SearchResponse{}, err
	}
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		fmt.Sprint(c.host, SearchPath, "?", r.GetURLValues().Encode()),
		nil,
	)
	if err != nil {
		return SearchResponse{}, err
	}
	return c.do(req)
}

func (c *Client) Lookup(ctx context.Context, r LookupRequest) (LookupResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (c *Client) do(req *http.Request) (SearchResponse, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return SearchResponse{}, err
	}
	defer func() { _ = resp.Body.Close() }()
	err = checkResponse(resp)
	if err != nil {
		return SearchResponse{}, err
	}
	var res Results
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return SearchResponse{}, err
	}
	return SearchResponse{Results: res}, nil
}

func checkResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}
	return fmt.Errorf("unexpected response: %s", resp.Status)
}
