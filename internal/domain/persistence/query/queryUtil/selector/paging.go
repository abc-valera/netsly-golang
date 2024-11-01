package selector

// Paging contains the data needed to perform pagination over the datasource.
type Paging struct {
	// Limit is the maximum number of records to be returned.
	// If no value is provided, all records will be returned.
	Limit uint
	// Offset is the number of records to skip before starting to return records.
	Offset uint
}
