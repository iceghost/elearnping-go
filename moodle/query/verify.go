package query

import (
	"encoding/json"
)

type verifyToken struct{}

func NewVerifyQuery() Query[bool] {
	return verifyToken{}
}

func (fn verifyToken) Encode() (string, map[string]string) {
	return "core_course_get_enrolled_courses_by_timeline_classification",
		map[string]string{"classification": "past"}
}

func (fn verifyToken) Decode(decoder *json.Decoder) bool {
	var err struct {
		ErrorCode string `json:"errorcode"`
	}
	decoder.Decode(&err)
	return err.ErrorCode != "invalidtoken"
}
