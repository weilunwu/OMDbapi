package OMDbapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

const (
	baseURL         = "http://omdbapi.com"
	defaultPlot     = "short"
	defaultResponse = "json"
	defaultYear     = ""
)

type searchResult struct {
	Title    string
	Response string
	Error    string
}

func query(title string) (*searchResult, error) {
	// set up parameters
	URL, err := url.Parse(baseURL)

	if err != nil {
		return nil, err
	}
	URL.Path += "/"

	// need factor these to deal with different inputs
	parameters := url.Values{}
	parameters.Add("t", title)
	parameters.Add("plot", defaultPlot)
	parameters.Add("r", defaultResponse)
	parameters.Add("y", defaultYear)
	URL.RawQuery = parameters.Encode()

	// make query
	resp, err := http.Get(URL.String())

	if err != nil {
		return nil, err
	}

	// decode json response
	result := &searchResult{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
