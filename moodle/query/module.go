package query

import (
	"elearnping-go/moodle"
	"encoding/json"
	"fmt"
)

type moduleQuery struct {
	Id moodle.ModuleId
}

func NewModuleQuery(id moodle.ModuleId) CachableQuery[moodle.Module] {
	return moduleQuery{id}
}

func (req moduleQuery) Encode() (name string, payload map[string]string) {
	return "core_course_get_course_module", map[string]string{"cmid": fmt.Sprint(req.Id)}
}

func (req moduleQuery) CacheKey() string {
	return fmt.Sprintf("module:%d", req.Id)
}

func (req moduleQuery) Decode(decoder *json.Decoder) moodle.Module {
	type Response struct {
		Module moodle.Module `json:"cm"`
	}
	var response Response
	decoder.Decode(&response)
	return response.Module
}
