package function

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"go.uber.org/ratelimit"
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

var rl = ratelimit.New(3)
var logger = log.New(os.Stdout, "[base fetch]", log.Ldate|log.Ltime)

func (fn BaseFunction[Res]) Fetch() (io.ReadCloser, error) {
	rl.Take()
	url := fn.URL()
	logger.Printf("Fetch %s %v", fn.iFn.Name(), fn.iFn.Arguments())
	res, err := http.Get(url.String())
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
