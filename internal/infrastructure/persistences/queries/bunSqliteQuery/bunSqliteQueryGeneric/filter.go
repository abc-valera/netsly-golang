package bunSqliteQueryGeneric

import (
	"reflect"
	"slices"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/convert"
	"github.com/uptrace/bun"
)

func ApplyFilter[DomainModel, BunModel any](
	query *bun.SelectQuery,
	dto func(DomainModel) BunModel,
	genericFilter selector.Filter[DomainModel],
) error {
	bunByModel := dto(genericFilter.By)

	modelType := reflect.TypeOf(bunByModel)

	if modelType.Kind() != reflect.Struct {
		return coderr.NewInternalString("bunModel must be a struct")
	}

	modelValue := reflect.ValueOf(bunByModel)

	if !modelValue.IsValid() {
		return coderr.NewInternalString("reflect.Value must be valid")
	}

	for i := 0; i < modelType.NumField(); i++ {
		fieldValue := modelValue.Field(i)
		fieldKind := fieldValue.Kind()
		fieldType := modelType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		if !fieldType.IsExported() {
			return coderr.NewInternalString("Filter fields must be exported")
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

		if slices.Contains(forbiddenKinds, fieldKind) {
			return coderr.NewInternalString("Field of kind " + fieldKind.String() + " can't be used as filter")
		}

		if genericFilter.IsRegex {
			query = query.Where(convert.ToSnakeCase(fieldType.Name)+" REGEXP ?", fieldValue.Interface())
		} else {
			query = query.Where(convert.ToSnakeCase(fieldType.Name)+" = ?", fieldValue.Interface())
		}
	}

	return nil
}
