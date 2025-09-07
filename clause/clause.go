package clause

type Clause int

const (
	FROM Clause = iota
	WHERE
	ORDERBY
	GROUPBY
	LIMIT
	TABLE
	DATABASE
	INTO
	VALUES
)

func (c Clause) String() string {
	switch c {
	case FROM:
		return "FROM"
	case WHERE:
		return "WHERE"
	case ORDERBY:
		return "ORDER BY"
	case GROUPBY:
		return "GROUP BY"
	case LIMIT:
		return "LIMIT"
	case TABLE:
		return "TABLE"
	case DATABASE:
		return "DATABASE"
	case INTO:
		return "INTO"
	case VALUES:
		return "VALUES"
	default:
		return ""
	}
}
