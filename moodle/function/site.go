package function

import (
	"elearnping-go/moodle"
	"encoding/json"
)

type GetSitesFunction struct {
	Classification string
}

func (req GetSitesFunction) Name() string {
	return "core_course_get_enrolled_courses_by_timeline_classification"
}

func (req GetSitesFunction) Arguments() map[string]string {
	return map[string]string{"classification": req.Classification}
}

func (req GetSitesFunction) Decode(_ string, decoder *json.Decoder) map[string][]moodle.Site {
	type Response struct {
		Sites []struct {
			moodle.Site
			Category string `json:"coursecategory"`
		} `json:"courses"`
	}
	var response Response
	decoder.Decode(&response)
	output := make(map[string][]moodle.Site)
	for _, catsite := range response.Sites {
		output[catsite.Category] = append(output[catsite.Category], catsite.Site)
	}
	return output
}
