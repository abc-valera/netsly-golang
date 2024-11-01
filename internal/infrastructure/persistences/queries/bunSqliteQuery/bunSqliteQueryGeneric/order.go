package bunSqliteQueryGeneric

import (
	"reflect"
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/selector"
	"github.com/abc-valera/netsly-golang/internal/domain/util/coderr"
	"github.com/abc-valera/netsly-golang/internal/domain/util/convert"
	"github.com/uptrace/bun"
)

func ApplyOrder[DomainModel, BunModel any](
	query *bun.SelectQuery,
	dto func(DomainModel) BunModel,
	genericOrder selector.Order[DomainModel],
) error {
	bunOrder := dto(genericOrder.By)

	orderByType := reflect.TypeOf(bunOrder)

	if orderByType.Kind() != reflect.Struct {
		return coderr.NewInternalString("Order.By must be a struct")
	}

	orderByValue := reflect.ValueOf(bunOrder)

	if !orderByValue.IsValid() {
		return coderr.NewInternalString("Order.By must be valid")
	}

	for i := range orderByType.NumField() {
		fieldValue := orderByValue.Field(i)
		fieldType := orderByType.Field(i)

		if fieldValue.IsZero() {
			continue
		}

		if !fieldType.IsExported() {
			return coderr.NewInternalString("Order field must be exported")
		}

		var orderString string

		switch value := fieldValue.Interface().(type) {
		case int:
			switch value {
			case selector.AscInt:
				orderString = "ASC"
			case selector.DescInt:
				orderString = "DESC"
			default:
				return coderr.NewInternalString("Invalid order field value")
			}
		case float32, float64:
			switch value {
			case selector.AscFloat:
				orderString = "ASC"
			case selector.DescFloat:
				orderString = "DESC"
			default:
				return coderr.NewInternalString("Invalid order field value")
			}
		case string:
			switch value {
			case selector.AscString:
				orderString = "ASC"
			case selector.DescString:
				orderString = "DESC"
			default:
				return coderr.NewInternalString("Invalid order field value")
			}
		case time.Time:
			switch value {
			case selector.AscTime:
				orderString = "ASC"
			case selector.DescTime:
				orderString = "DESC"
			default:
				return coderr.NewInternalString("Invalid order field value")
			}
		default:
			return coderr.NewInternalString("Invalid order field type")
		}

		query = query.Order(convert.ToSnakeCase(fieldType.Name) + " " + orderString)
	}

	return nil
}
