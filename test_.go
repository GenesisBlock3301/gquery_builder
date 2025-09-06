package main

import (
	"fmt"
	"strings"
)

type QueryBuilder struct {
	table   string
	columns []string
	where   string
	orderBy []string
}

func (qb *QueryBuilder) Select(columns ...string) *QueryBuilder {
	qb.columns = columns
	return qb
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.table = table
	return qb
}

func (qb *QueryBuilder) Where(condition string) *QueryBuilder {
	qb.where = condition
	return qb
}

func (qb *QueryBuilder) OrderBy(columns ...string) *QueryBuilder {
	qb.orderBy = columns
	return qb
}

func (qb *QueryBuilder) Build() string {
	return fmt.Sprintf("SELECT %s FROM %s", strings.Join(qb.columns, ", "), qb.table, qb.where)
}
