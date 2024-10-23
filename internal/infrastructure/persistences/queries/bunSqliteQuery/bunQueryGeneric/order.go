package bunQueryGeneric

import (
	"reflect"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/order"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/convert"
	"github.com/uptrace/bun"
)

func ApplyOrder[DomainModel, BunModel any](
	query *bun.SelectQuery,
	dto func(DomainModel) BunModel,
	genericOrder order.Order[DomainModel],
) error {
	bunOrder := dto(genericOrder.By)

	orderByValue := reflect.ValueOf(bunOrder)
	orderByType := reflect.TypeOf(bunOrder)
	orderByKind := orderByType.Kind()

	if orderByKind != reflect.Struct {
		return coderr.NewInternalString("Invalid order value")
	}

	for i := range orderByType.NumField() {
		fieldValue := orderByValue.Field(i)
		fieldType := orderByType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		switch fieldValue.Kind() {
		case reflect.Int:
			value, ok := fieldValue.Interface().(int)
			if !ok {
				return coderr.NewInternalString("Invalid order value")
			}

			var orderString string
			if value == order.AscInt {
				orderString = "ASC"
			} else if value == order.DescInt {
				orderString = "DESC"
			} else {
				return coderr.NewInternalString("Invalid order value")
			}

			query = query.Order(convert.ToSnakeCase(fieldType.Name) + " " + orderString)

		case reflect.String:
			value, ok := fieldValue.Interface().(string)
			if !ok {
				return coderr.NewInternalString("Invalid order value")
			}

			var orderString string
			if value == order.AscString {
				orderString = "ASC"
			} else if value == order.DescString {
				orderString = "DESC"
			} else {
				return coderr.NewInternalString("Invalid order value")
			}

			query = query.Order(convert.ToSnakeCase(fieldType.Name) + " " + orderString)

		default:
			return coderr.NewInternalString("Invalid order value")
		}
	}

	return nil
}
