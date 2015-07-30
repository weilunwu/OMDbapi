package OMDbapi

import "net/http"

const API_ENDPOINT = "http://omdbapi.com"

func query(title string) (*http.Response, error) {
	return http.Get(API_ENDPOINT + "?t=" + title + "&r=json")
}
