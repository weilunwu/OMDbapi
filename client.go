package OMDbapi

import (
	"encoding/json"
	"fmt"
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
	Title      string
	Year       string
	Rated      string
	Released   string
	Runtime    string
	Genre      string
	Director   string
	Writer     string
	Actors     string
	Plot       string
	Language   string
	Country    string
	Awards     string
	Poster     string
	Metascore  string
	IMDBRating string
	IMDBVotes  string
	IMDBID     string
	Type       string
	Response   string
	Error      string
}

// QueryByTitle requires the details by giving a movie title
func QueryByTitle(title string) (*searchResult, error) {
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

func (sr *searchResult) String() string {
	return fmt.Sprintf("#%s: %s (%s) Type: %s", sr.IMDBID, sr.Title, sr.Year, sr.Type)
}
