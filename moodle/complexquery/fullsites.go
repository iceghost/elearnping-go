package complexquery

import (
	"elearnping-go/moodle"
	"elearnping-go/moodle/query"
)

func CallFullSites(token string) (map[string][]moodle.Site, error) {
	catsites, err := query.Call(query.NewSitesQuery("inprogress"), token)
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
