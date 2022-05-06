package moodle

import (
	"encoding/json"
	"net/http"
	"go.uber.org/ratelimit"
)

type MoodleService interface {
	Exec(function string, args map[string]string) (*http.Response, error)
}

var rl = ratelimit.New(10)

func Exec[Result any](moodle MoodleService, api APIFunction[Result]) (*Result, error) {
	rl.Take()
	res, err := moodle.Exec(api.Function(), api.Arguments())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	decoder := json.NewDecoder(res.Body)
	result := api.Decode(decoder, moodle)
	return &result, err
}

type APIFunction[Res any] interface {
	Function() string
	Arguments() map[string]string
	Decode(*json.Decoder, MoodleService) Res
}
