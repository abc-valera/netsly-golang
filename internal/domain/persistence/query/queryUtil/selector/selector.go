// selector package contains abstractions over the common database operations.
package selector

type Selector[Model any] struct {
	Paging  Paging
	Filters []Filter[Model]
	Orders  []Order[Model]
}

type Option[Model any] func(*Selector[Model])

func New[Model any](options ...Option[Model]) Selector[Model] {
	s := Selector[Model]{
		Filters: make([]Filter[Model], 0),
		Orders:  make([]Order[Model], 0),
	}

	for _, option := range options {
		option(&s)
	}

	return s
}

func WithLimit[Model any](limit uint) Option[Model] {
	return func(s *Selector[Model]) {
		s.Paging.Limit = limit
	}
}

func WithOffset[Model any](offset uint) Option[Model] {
	return func(s *Selector[Model]) {
		s.Paging.Offset = offset
	}
}

func WithFilter[Model any](filter Model) Option[Model] {
	return func(s *Selector[Model]) {
		s.Filters = append(s.Filters, Filter[Model]{By: filter})
	}
}

func WithRegex[Model any](regex Model) Option[Model] {
	return func(s *Selector[Model]) {
		s.Filters = append(s.Filters, Filter[Model]{By: regex, IsRegex: true})
	}
}

func WithOrder[Model any](order Model) Option[Model] {
	return func(s *Selector[Model]) {
		s.Orders = append(s.Orders, Order[Model]{By: order})
	}
}
