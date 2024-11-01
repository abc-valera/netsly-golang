package selector

// Filter is used as an abstraction over the datasource filtering mechanism.
//
// It allows both 'REGEXP' and '==' operations.
type Filter[Model any] struct {
	// By must be a struct, more precisely a domain/model,
	// with all the fields unitialized,
	// except the ones that will be used as filters.
	By Model
	// IsRegex flag reports if
	//  - the REGEXP operation will be performed.
	//    Example: SELECT * FROM table WHERE field REGEXP 'value';
	//  - Or the == operation will be performed.
	// 	  Example: SELECT * FROM table WHERE field = 'value';
	IsRegex bool
}
