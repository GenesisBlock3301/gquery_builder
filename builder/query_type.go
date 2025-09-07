package builder

// QueryType represents the type of SQL query
type QueryType int

const (
	SELECT QueryType = iota
	INSERT
	UPDATE
	DELETE
	CREATE
	DROP
)

// String converts QueryType to its SQL string representation
func (qt QueryType) String() string {
	switch qt {
	case SELECT:
		return "SELECT"
	case INSERT:
		return "INSERT"
	case UPDATE:
		return "UPDATE"
	case DELETE:
		return "DELETE"
	case CREATE:
		return "CREATE"
	case DROP:
		return "DROP"
	default:
		return ""
	}
}
