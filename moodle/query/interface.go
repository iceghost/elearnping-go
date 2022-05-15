package query

import (
	"elearnping-go/moodle"
	"encoding/json"
	"net/url"
)

// interface for cachable queries (e.g module informations, updates of the same group)
type CachableQuery[Res any] interface {
	Query[Res]
	CacheKey() string
}

// an interface for moodle web service api calls
// see https://docs.moodle.org/dev/Web_service_API_functions for list of functions
// and https://github.com/moodle/moodle for payloads
type Query[Res any] interface {
	Encode() (name string, payload map[string]string)
	Decode(decoder *json.Decoder) Res
}

var endpoint = url.URL{
	Scheme: "http",
	Host:   "e-learning.hcmut.edu.vn",
	Path:   "/webservice/rest/server.php",
}

// make url for moodle web service api calls
//
// using format:
// http://e-learning.hcmut.edu.vn/webservice/rest/server.php
// 		?wstoken={{token}}&wsfunction={{function}}&moodlewsrestformat=json
func URL[Res any](q Query[Res], token string) url.URL {
	name, payload := q.Encode()

	newEndpoint := endpoint
	args := url.Values{}
	args.Add("wstoken", token)
	args.Add("wsfunction", name)
	args.Add("moodlewsrestformat", "json")
	for key, val := range payload {
		args.Add(key, val)
	}
	newEndpoint.RawQuery = args.Encode()
	return newEndpoint
}

// call moodle web service api through Query[Res] interface
func Call[Res any](q Query[Res], token string) (Res, error) {
	if body, err := moodle.Fetch(URL(q, token)); err != nil {
		var zero Res
		return zero, err
	} else {
		defer body.Close()
		decoder := json.NewDecoder(body)
		return q.Decode(decoder), nil
	}
}
