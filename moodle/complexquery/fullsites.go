package complexquery

import (
	"elearnping-go/moodle"
	"elearnping-go/moodle/query"
)

// get sites, with group id information
func CallFullSites(token string) (map[moodle.Category][]moodle.Site, error) {
	catsites, err := query.Call(query.NewSitesQuery(moodle.INPROGRESS), token)
	if err != nil {
		return nil, err
	}

	groups, err := query.Call(query.NewGroupsQuery(), token)
	if err != nil {
		return nil, err
	}

	for cat := range catsites {
		for i := range catsites[cat] {
			catsites[cat][i].GroupId = groups[catsites[cat][i].Id]
		}
	}

	return catsites, nil
}
