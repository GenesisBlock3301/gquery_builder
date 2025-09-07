package builder

import (
	"fmt"
	"gquery_builder/clause"
	"strings"
)

// QueryBuilder builds queries
type QueryBuilder struct {
	queryType QueryType
	tableName string
	fields    []interface{}
	condition []string
	values    []interface{}
}

// Select sets fields
func (qb *QueryBuilder) Select(fields ...interface{}) *QueryBuilder {
	qb.queryType = SELECT
	qb.fields = fields
	return qb
}

// From sets table name
func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.tableName = table
	return qb
}

// Where adds conditions
func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.condition = append(qb.condition, condition)
	return qb
}

func (qb *QueryBuilder) Into(values ...interface{}) *QueryBuilder {
	for _, value := range values {
		switch value.(type) {
		case string:
			qb.values = append(qb.values, fmt.Sprintf("'%s'", value))
		default:
			qb.values = append(qb.values, fmt.Sprintf("%v", value))
		}
	}
	return qb
}

// Build generates SQL
func (qb *QueryBuilder) Build() string {
	switch qb.queryType {
	case SELECT:
		fields := "*"
		if len(qb.fields) > 0 {
			quotedFields := make([]string, len(qb.fields))
			for i, field := range qb.fields {
				quotedFields[i] = fmt.Sprintf("\"%s\"", field)
			}
			fields = strings.Join(quotedFields, ", ")
		}
		query := fmt.Sprintf("%s %s %s \"%s\"", qb.queryType, fields, clause.FROM.String(), qb.tableName)
		if len(qb.condition) > 0 {
			query += fmt.Sprintf(" %s %s", clause.WHERE.String(), strings.Join(qb.condition, " AND "))
		}
		return query
	case INSERT:
		values := make([]string, len(qb.values))
		if len(qb.values) > 0 {
			for i, v := range qb.values {
				values[i] = fmt.Sprintf("%s", v)
			}
		}
		return fmt.Sprintf("%s %s %s %s (%s)", INSERT.String(), clause.INTO.String(),
			qb.tableName, clause.VALUES.String(), strings.Join(values, ", "))
	case UPDATE:
		return fmt.Sprintf("%s %s", UPDATE.String(), qb.tableName)
	case DELETE:
		return fmt.Sprintf("%s %s", DELETE.String(), qb.tableName)
	default:
		return ""
	}
}
