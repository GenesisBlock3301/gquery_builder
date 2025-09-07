package builder

import (
	"fmt"
	"gquery_builder/clause"
)

// CreateQuery Builder
type CreateQueryBuilder struct {
	tableName string
}

func (cb *CreateQueryBuilder) CreateTable(table string) *CreateQueryBuilder {
	cb.tableName = table
	return cb
}

func (cb *CreateQueryBuilder) Build() string {
	return fmt.Sprintf("%s %s %s", CREATE.String(), clause.TABLE.String(), cb.tableName)
}
