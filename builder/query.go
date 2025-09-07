package builder

// Query factory
type Query struct{}

// Select creates a new QueryBuilder
func (q Query) Select(fields ...interface{}) *QueryBuilder {
	return &QueryBuilder{
		queryType: SELECT,
		fields:    fields,
	}
}

func (q Query) From(table string) *QueryBuilder {
	return &QueryBuilder{
		queryType: SELECT,
		tableName: table,
	}
}

func (q Query) Update(table string) *QueryBuilder {
	return &QueryBuilder{
		queryType: UPDATE,
		tableName: table,
	}
}

func (q Query) InsertInto(table string) *QueryBuilder {
	return &QueryBuilder{
		queryType: INSERT,
		tableName: table,
	}
}

func (q Query) DeleteFrom(table string) *QueryBuilder {
	return &QueryBuilder{
		queryType: DELETE,
		tableName: table,
	}
}

func (q Query) CreateTable(table string) *CreateQueryBuilder {
	return (&CreateQueryBuilder{}).CreateTable(table)
}

func (q Query) DropTable(tableName string) *DropQueryBuilder {
	return (&DropQueryBuilder{}).DropTable(tableName)
}

func (q Query) DropDatabase(databaseName string) *DropQueryBuilder {
	return (&DropQueryBuilder{}).DropDatabase(databaseName)
}
