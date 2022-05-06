package moodle

import (
	"net/http"
	"net/url"
)

var endpoint = url.URL{
	Scheme: "http",
	Host:   "e-learning.hcmut.edu.vn",
	Path:   "/webservice/rest/server.php",
}

type BaseMoodleService struct {
	Token string
}

func (service *BaseMoodleService) Exec(function string, args map[string]string) (*http.Response, error) {
	newEndpoint := endpoint
	query := url.Values{}
	query.Add("wstoken", service.Token)
	query.Add("wsfunction", function)
	query.Add("moodlewsrestformat", "json")
	for key, val := range args {
		query.Add(key, val)
	}
	newEndpoint.RawQuery = query.Encode()
	return http.Get(newEndpoint.String())
}
