package query

import (
	"elearnping-go/moodle"
	"encoding/json"
)

type sitesQuery struct {
	Classification string
}

func NewSitesQuery(classification string) Query[map[string][]moodle.Site] {
	return sitesQuery{classification}
}

func (req sitesQuery) Encode() (string, map[string]string) {
	return "core_course_get_enrolled_courses_by_timeline_classification",
		map[string]string{"classification": req.Classification}
}

func (req sitesQuery) Decode(decoder *json.Decoder) map[string][]moodle.Site {
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
