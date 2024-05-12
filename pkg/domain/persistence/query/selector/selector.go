package selector

type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

type Selector struct {
	Order  Order
	Limit  uint
	Offset uint
}
