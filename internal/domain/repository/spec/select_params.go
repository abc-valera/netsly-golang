package spec

// SelectParams represents query data for specifying select details.
type SelectParams struct {
	OrderBy string // OrderField is field by which sorting will be performed (usually is 'created_at')
	Order   string // Order is order of sorting ('acs' or 'desc')
	Limit   uint   // Limit limits number of returned units
	Offset  uint   // Offset sets an offset for returned units
}
