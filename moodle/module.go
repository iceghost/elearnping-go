package moodle

import (
	"encoding/json"
	"fmt"
)

type ModuleId uint

type Module struct {
	Id         ModuleId `json:"id"`
	Name       string   `json:"name"`
	PluginName string   `json:"modname"`
}

type GetModule struct {
	Id ModuleId
}

func (req GetModule) Function() string {
	return "core_course_get_course_module"
}

func (req GetModule) Arguments() map[string]string {
	return map[string]string{"cmid": fmt.Sprint(req.Id)}
}

func (req GetModule) Decode(decoder *json.Decoder, _ MoodleService) Module {
	type Response struct {
		Module Module `json:"cm"`
	}
	var response Response
	decoder.Decode(&response)
	return response.Module
}
