package query

import (
	"elearnping-go/moodle"
	"encoding/json"
	"net/url"
)

type CachableQuery[Res any] interface {
	Query[Res]
	CacheKey() string
}

type Query[Res any] interface {
	Encode() (name string, payload map[string]string)
	Decode(decoder *json.Decoder) Res
}

var endpoint = url.URL{
	Scheme: "http",
	Host:   "e-learning.hcmut.edu.vn",
	Path:   "/webservice/rest/server.php",
}

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
