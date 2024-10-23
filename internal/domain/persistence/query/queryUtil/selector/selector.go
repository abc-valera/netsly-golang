package selector

import (
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/order"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/paging"
)

type Selector[DomainModel any] struct {
	Paging  paging.Paging
	Filters []filter.Filter[DomainModel]
	Order   order.Order[DomainModel]
}

type Option[DomainModel any] func(*Selector[DomainModel])

func New[DomainModel any](options ...Option[DomainModel]) Selector[DomainModel] {
	var s Selector[DomainModel]
	for _, option := range options {
		option(&s)
	}
	return s
}

func WithLimit[DomainModel any](limit uint) Option[DomainModel] {
	return func(s *Selector[DomainModel]) {
		s.Paging.Limit = limit
	}
}

func WithOffset[DomainModel any](offset uint) Option[DomainModel] {
	return func(s *Selector[DomainModel]) {
		s.Paging.Offset = offset
	}
}

func WithFilter[DomainModel any](filterBy DomainModel) Option[DomainModel] {
	return func(s *Selector[DomainModel]) {
		s.Filters = append(s.Filters, filter.Filter[DomainModel]{By: filterBy})
	}
}

func WithRegexFilter[DomainModel any](filterBy DomainModel) Option[DomainModel] {
	return func(s *Selector[DomainModel]) {
		s.Filters = append(s.Filters, filter.Filter[DomainModel]{By: filterBy, IsRegex: true})
	}
}

func WithOrder[DomainModel any](orderBy DomainModel) Option[DomainModel] {
	return func(s *Selector[DomainModel]) {
		s.Order = order.Order[DomainModel]{By: orderBy}
	}
}
