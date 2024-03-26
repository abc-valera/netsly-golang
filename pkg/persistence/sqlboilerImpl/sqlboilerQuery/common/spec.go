package common

import (
	"slices"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/spec"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func ToBoilerSelectParams(params spec.SelectParams) []qm.QueryMod {
	mods := make([]qm.QueryMod, 0, 3)

	if params.Order() == spec.OrderAsc {
		mods = append(mods, qm.OrderBy("created_at asc"))
	} else {
		mods = append(mods, qm.OrderBy("created_at desc"))
	}

	mods = append(mods, qm.Limit(params.Limit()))
	mods = append(mods, qm.Offset(params.Offset()))

	return mods
}

func ToBoilerSelectParamsPipe(params spec.SelectParams, mods ...qm.QueryMod) []qm.QueryMod {
	return slices.Concat(mods, ToBoilerSelectParams(params))
}
