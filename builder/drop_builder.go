package builder

import (
	"fmt"
	"gquery_builder/clause"
)

// DropQuery Builder
type DropQueryBuilder struct {
	objectType string
	name       string
}

func (db *DropQueryBuilder) DropTable(table string) *DropQueryBuilder {
	db.objectType = clause.TABLE.String()
	db.name = table
	return db
}

func (db *DropQueryBuilder) DropDatabase(database string) *DropQueryBuilder {
	db.objectType = clause.DATABASE.String()
	db.name = database
	return db
}

func (db *DropQueryBuilder) Build() string {
	return fmt.Sprintf("%s %s %s", DROP.String(), db.objectType, db.name)
}
