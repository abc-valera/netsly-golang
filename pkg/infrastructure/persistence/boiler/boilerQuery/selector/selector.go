package selector

import (
	"slices"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selector"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ToBoilerSelector(params selector.Selector) []qm.QueryMod {
	mods := make([]qm.QueryMod, 0, 3)

	if params.Order == selector.OrderAsc {
		mods = append(mods, qm.OrderBy("created_at asc"))
	} else {
		mods = append(mods, qm.OrderBy("created_at desc"))
	}

	mods = append(mods, qm.Limit(int(params.Limit)))
	mods = append(mods, qm.Offset(int(params.Offset)))

	return mods
}

func ToBoilerSelectorPipe(params selector.Selector, mods ...qm.QueryMod) []qm.QueryMod {
	return slices.Concat(mods, ToBoilerSelector(params))
}
