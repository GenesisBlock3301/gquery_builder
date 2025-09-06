package builder

import (
	"fmt"
	"strings"
)

type Builder interface {
	Build() string
}

// QueryBuilder Implement coming Query
type QueryBuilder struct {
	queryType string
	tableName string
	fields    []string
	condition []string
}

func (qb *QueryBuilder) Select(fields ...string) *QueryBuilder {
	qb.queryType = "SELECT"
	qb.fields = fields
	return qb
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.tableName = table
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.condition = append(qb.condition, condition)
	return qb
}

func (qb *QueryBuilder) Build() string {
	switch qb.queryType {
	case "SELECT":
		fields := "*"
		if len(qb.fields) > 0 {
			fields = strings.Join(qb.fields, ",")
		}
		query := fmt.Sprintf("%s FROM %s %s", qb.queryType, qb.tableName, fields)
		if len(qb.condition) > 0 {
			query += fmt.Sprintf(" WHERE %s", strings.Join(qb.condition, " AND "))
		}
		return query
	default:
		return ""
	}

}

// Query Make query
type Query struct{}

func (q Query) Select(fields ...string) *QueryBuilder {
	return &QueryBuilder{
		queryType: "SELECT",
		fields:    fields,
	}
}
func (q Query) From(table string) *QueryBuilder {
	return &QueryBuilder{
		queryType: "FROM",
		tableName: table,
	}
}
func (q Query) Where(condition string) *QueryBuilder {
	return &QueryBuilder{
		queryType: "WHERE",
		condition: []string{condition},
	}
}
