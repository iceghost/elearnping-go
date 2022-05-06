package moodle

import (
	"encoding/json"
)

type SiteId uint

type Site struct {
	Id   SiteId `json:"id"`
	Name string `json:"fullname"`
}

type GetSites struct {
	Classification string
}

func (req GetSites) Function() string {
	return "core_course_get_enrolled_courses_by_timeline_classification"
}

func (req GetSites) Arguments() map[string]string {
	return map[string]string{"classification": req.Classification}
}

func (req GetSites) Decode(decoder *json.Decoder, _ MoodleService) map[string][]Site {
	type Response struct {
		Sites []struct {
			Site
			Category string `json:"coursecategory"`
		} `json:"courses"`
	}
	var response Response
	decoder.Decode(&response)
	output := make(map[string][]Site)
	for _, catsite := range response.Sites {
		output[catsite.Category] = append(output[catsite.Category], catsite.Site)
	}
	return output
}
