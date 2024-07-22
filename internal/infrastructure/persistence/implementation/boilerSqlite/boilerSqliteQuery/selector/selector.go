package selector

import (
	"slices"

	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query/selector"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ToBoilerSelector(s selector.Selector) []qm.QueryMod {
	mods := make([]qm.QueryMod, 0, 3)

	if s.Order == selector.OrderAsc {
		mods = append(mods, qm.OrderBy("created_at asc"))
	} else {
		mods = append(mods, qm.OrderBy("created_at desc"))
	}

	mods = append(mods, qm.Limit(int(s.Limit)))
	mods = append(mods, qm.Offset(int(s.Offset)))

	return mods
}

func ToBoilerSelectorPipe(selector selector.Selector, mods ...qm.QueryMod) []qm.QueryMod {
	return slices.Concat(mods, ToBoilerSelector(selector))
}
