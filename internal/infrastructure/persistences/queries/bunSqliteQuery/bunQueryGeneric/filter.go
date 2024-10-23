package bunQueryGeneric

import (
	"reflect"
	"slices"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/convert"
	"github.com/uptrace/bun"
)

func ApplyFilter[DomainModel, BunModel any](
	query *bun.SelectQuery,
	dto func(DomainModel) BunModel,
	genericFilter filter.Filter[DomainModel],
) error {
	bunByModel := dto(genericFilter.By)

	filterByValue := reflect.ValueOf(bunByModel)
	filterByType := reflect.TypeOf(bunByModel)
	filterByKind := filterByType.Kind()

	if filterByKind != reflect.Struct {
		return coderr.NewInternalString("Invalid filter value")
	}

	for i := 0; i < filterByType.NumField(); i++ {
		fieldValue := filterByValue.Field(i)
		fieldType := filterByType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		if !fieldType.IsExported() {
			return coderr.NewInternalString("Invalid filter value")
		}

		forbiddenKinds := []reflect.Kind{
			reflect.Uintptr,
			reflect.Complex64,
			reflect.Complex128,
			reflect.Array,
			reflect.Chan,
			reflect.Func,
			reflect.Interface,
			reflect.Map,
			reflect.Pointer,
			reflect.Slice,
			reflect.Struct,
			reflect.UnsafePointer,
		}

		if slices.Contains(forbiddenKinds, fieldValue.Kind()) {
			return coderr.NewInternalString("Invalid filter value")
		}

		query = query.Where(convert.ToSnakeCase(fieldType.Name)+" = ?", fieldValue.Interface())
	}

	return nil
}
