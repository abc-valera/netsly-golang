package filter

type Filter[DomainModel any] struct {
	// By must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the one that will be used as a filter.
	//
	// It's value will be used as a filter.
	By DomainModel
	// IsRegex flag reports if
	//  - the REGEXP operation will be performed.
	//    Example: SELECT * FROM table WHERE field REGEXP 'value';
	//  - Or the == operation will be performed.
	// 	  Example: SELECT * FROM table WHERE field = 'value';
	IsRegex bool
}

type Option[DomainModel any] func(*Filter[DomainModel])

func New[DomainModel any](options ...Option[DomainModel]) []Filter[DomainModel] {
	filters := make([]Filter[DomainModel], 0)
	for _, option := range options {
		filter := Filter[DomainModel]{}
		option(&filter)
		filters = append(filters, filter)
	}
	return filters
}

func By[DomainModel any](filterBy DomainModel) Option[DomainModel] {
	return func(f *Filter[DomainModel]) {
		f.By = filterBy
	}
}

func ByRegex[DomainModel any](filterBy DomainModel) Option[DomainModel] {
	return func(f *Filter[DomainModel]) {
		f.By = filterBy
		f.IsRegex = true
	}
}
