package tfl

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type (
	//Client
	Client struct {
		baseURL string
		client  *http.Client
	}

	// Place where bikes live
	Place struct {
		Type                 string            `json:"$type"`
		ID                   string            `json:"id"`
		URL                  string            `json:"url"`
		CommonName           string            `json:"commonName"`
		PlaceType            string            `json:"placeType"`
		AdditionalProperties []AdditionalProp  `json:"additionalProperties"`
		Children             []json.RawMessage `json:"children"`
		ChildrenUrls         []json.RawMessage `json:"childrenUrls"`
		Lat                  float64           `json:"lat"`
		Lon                  float64           `json:"lon"`
	}
	// AdditionalProp data
	AdditionalProp struct {
		Type            string    `json:"$type"`
		Category        string    `json:"category"`
		Key             string    `json:"key"`
		SourceSystemKey string    `json:"sourceSystemKey"`
		Value           string    `json:"value"`
		Modified        time.Time `json:"modified"`
	}
	// PlaceIndex to lookup based on name
	PlaceIndex struct {
		CommonName string  `json:"commonName"`
		ID         string  `json:"id"`
		URL        string  `json:"url"`
		Lat        float64 `json:"lat"`
		Lon        float64 `json:"lon"`
	}
	// Snapshot of current state
	Snapshot struct {
		CommonName   string     `json:"commonName,omitempty"`
		UpdatedAt    time.Time  `json:"updatedAt,omitempty"`
		TerminalName string     `json:"terminalName,omitempty"`
		Installed    bool       `json:"installed,omitempty"`
		Locked       bool       `json:"locked,omitempty"`
		InstallDate  *time.Time `json:"installDate,omitempty"`
		RemovalDate  *time.Time `json:"removalDate,omitempty"`
		Temporary    bool       `json:"temporary,omitempty"`
		Bikes        int        `json:"bikes,omitempty"`
		EmptyDocks   int        `json:"emptyDocks,omitempty"`
		Docks        int        `json:"docks,omitempty"`
	}
	//ClientOptions for client
	ClientOptions func(c *Client)
)

func WithBase(url string) ClientOptions {
	return func(c *Client) {
		c.baseURL = url
	}
}

// NewBikeClient to query TFL api
func NewBikeClient(client *http.Client, options ...ClientOptions) Client {
	ret := Client{
		baseURL: "https://api.tfl.gov.uk",
		client:  client,
	}
	for _, opt := range options {
		opt(&ret)
	}
	return ret
}

// ViewIndex for all sites
func (c Client) ViewIndex() ([]Place, error) {
	var results []Place

	resp, err := c.getRequest("bikepoint")
	if err != nil {
		return nil, err
	}

	err = readJSON(resp.Body, &results)
	if err != nil {
		return nil, err
	}
	return results, nil
}

// ViewDocking details
func (c Client) ViewDocking(id string) (*Place, error) {
	var results Place

	resp, err := c.getRequest(fmt.Sprintf("Place/BikePoints_%s", id))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("Not found")
	}
	err = readJSON(resp.Body, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}

func (c Client) getRequest(path string) (*http.Response, error) {
	r, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", c.baseURL, path), nil)
	if err != nil {
		return nil, errors.WithMessage(err, "cannot create request")
	}
	return c.client.Do(r)
}

func readJSON(response io.ReadCloser, result interface{}) error {
	var err error
	defer func() {
		err = response.Close()
	}()

	buf, err := ioutil.ReadAll(response)
	if err != nil {
		return errors.WithMessage(err, "unable to read body")
	}

	err = json.Unmarshal(buf, result)
	if err != nil {
		return errors.WithMessage(err, "bad json response")
	}
	return err
}
