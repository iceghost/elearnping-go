package function

import (
	"elearnping-go/moodle"
	"encoding/json"
	"fmt"
)

type GetModuleFunction struct {
	Id moodle.ModuleId
}

func (req GetModuleFunction) Name() string {
	return "core_course_get_course_module"
}

func (req GetModuleFunction) Arguments() map[string]string {
	return map[string]string{"cmid": fmt.Sprint(req.Id)}
}

func (req GetModuleFunction) Decode(token string, decoder *json.Decoder) moodle.Module {
	type Response struct {
		Module moodle.Module `json:"cm"`
	}
	var response Response
	decoder.Decode(&response)
	return response.Module
}
