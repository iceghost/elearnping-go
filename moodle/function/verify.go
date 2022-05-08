package function

import (
	"encoding/json"
)

type VerifyTokenFunction struct{}

func (fn VerifyTokenFunction) Name() string {
	return "core_course_get_enrolled_courses_by_timeline_classification"
}

func (fn VerifyTokenFunction) Arguments() map[string]string {
	return map[string]string{"classification": "past"}
}

func (fn VerifyTokenFunction) Decode(_ string, decoder *json.Decoder) bool {
	var err struct {
		ErrorCode string `json:"errorcode"`
	}
	decoder.Decode(&err)
	return err.ErrorCode != "invalidtoken"
}
