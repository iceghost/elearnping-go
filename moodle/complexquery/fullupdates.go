package complexquery

import (
	"elearnping-go/moodle"
	"elearnping-go/moodle/cache"
	"elearnping-go/moodle/query"
	"time"
)

func CallFullUpdates(token string, site moodle.Site, since time.Time) (moodle.SiteUpdate, error) {
	var zero moodle.SiteUpdate

	updates, err := cache.Call(query.NewUpdatesQuery(site, since), token)

	if err != nil {
		return zero, nil
	}

	for i := range updates.Updates {
		module, err := cache.Call(query.NewModuleQuery(updates.Updates[i].Module.Id), token)
		if err != nil {
			return zero, nil
		}
		updates.Updates[i].Module = module
	}

	return updates, nil
}
