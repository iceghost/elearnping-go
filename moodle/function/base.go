package function

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type Function[Res any] interface {
	Name() string
	Arguments() map[string]string
	Decode(token string, decoder *json.Decoder) Res
}

type BaseFunction[Res any] struct {
	token string
	iFn   Function[Res]
}

func NewFunction[Res any](token string, iFn Function[Res]) BaseFunction[Res] {
	return BaseFunction[Res]{token, iFn}
}

var endpoint = url.URL{
	Scheme: "http",
	Host:   "e-learning.hcmut.edu.vn",
	Path:   "/webservice/rest/server.php",
}

func (fn BaseFunction[Res]) URL() *url.URL {
	newEndpoint := endpoint
	query := url.Values{}
	query.Add("wstoken", fn.token)
	query.Add("wsfunction", fn.iFn.Name())
	query.Add("moodlewsrestformat", "json")
	for key, val := range fn.iFn.Arguments() {
		query.Add(key, val)
	}
	newEndpoint.RawQuery = query.Encode()
	return &newEndpoint
}

func (fn BaseFunction[Res]) Fetch() (io.ReadCloser, error) {
	res, err := http.Get(fn.URL().String())
	if err != nil {
		return nil, err
	}
	return res.Body, nil
}

func (fn BaseFunction[Res]) Call() (Res, error) {
	body, err := fn.Fetch()
	if err != nil {
		var zero Res
		return zero, err
	}
	defer body.Close()
	decoder := json.NewDecoder(body)
	return fn.iFn.Decode(fn.token, decoder), nil
}