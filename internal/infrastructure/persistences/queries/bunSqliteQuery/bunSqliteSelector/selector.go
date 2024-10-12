package bunSqliteSelector

import (
	"reflect"
	"slices"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/convert"
	"github.com/uptrace/bun"
)

func NewSelectQuery(db bun.IDB, s selector.Selector) *bun.SelectQuery {
	// Apply Limit and Offset
	query := db.NewSelect().Limit(int(s.Limit)).Offset(int(s.Offset))

	// Apply OrderBy and Order
	if s.OrderBy != nil {
		orderByValue := reflect.ValueOf(s.OrderBy)
		orderByType := reflect.TypeOf(s.OrderBy)
		orderByKind := orderByType.Kind()

		if orderByKind == reflect.Struct {
			for i := range orderByType.NumField() {
				fieldValue := orderByValue.Field(i)
				fieldType := orderByType.Field(i)

				if fieldValue.IsZero() {
					continue
				}

				fieldKind := fieldValue.Kind()
				allowedFieldKinds := []reflect.Kind{
					reflect.Int,
					reflect.Float32,
					reflect.String,
				}

				if slices.Contains(allowedFieldKinds, fieldKind) {
					orderString := "DESC"
					if s.Order {
						orderString = "ASC"
					}

					query.Order(convert.ToSnakeCase(fieldType.Name) + " " + orderString)
					break
				}
			}
		}
	}

	// Apply Filter
	if s.FilterBy != nil {
		filterByValue := reflect.ValueOf(s.OrderBy)
		filterByType := reflect.TypeOf(s.OrderBy)
		filterByKind := filterByType.Kind()

		if filterByKind == reflect.Struct {
			for i := range filterByType.NumField() {
				fieldValue := filterByValue.Field(i)
				fieldType := filterByType.Field(i)
				fieldKind := fieldValue.Kind()

				if fieldValue.IsZero() {
					continue
				}

				if fieldKind == reflect.String {
					query.Where(convert.ToSnakeCase(fieldType.Name)+" REGEXP ?", s.Filter)
					break
				}
			}
		}
	}

	return query
}
