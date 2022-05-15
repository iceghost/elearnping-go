package query

import (
	"elearnping-go/moodle"
	"encoding/json"
)

type sitesQuery struct {
	Classification moodle.Classification
}

// get enrolled sites
//
// also note that site's group ids are not known, due to only one trip to server
// if you want to know about group ids, use elearnping-go/moodle/complexquery
func NewSitesQuery(classification moodle.Classification) Query[map[moodle.Category][]moodle.Site] {
	return sitesQuery{classification}
}

func (req sitesQuery) Encode() (string, map[string]string) {
	return "core_course_get_enrolled_courses_by_timeline_classification",
		map[string]string{"classification": string(req.Classification)}
}

func (req sitesQuery) Decode(decoder *json.Decoder) map[moodle.Category][]moodle.Site {
	type Response struct {
		Sites []struct {
			moodle.Site
			Category moodle.Category `json:"coursecategory"`
		} `json:"courses"`
	}
	var response Response
	decoder.Decode(&response)
	output := make(map[moodle.Category][]moodle.Site)
	for _, catsite := range response.Sites {
		output[catsite.Category] = append(output[catsite.Category], catsite.Site)
	}
	return output
}
