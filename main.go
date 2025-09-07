package main

import (
	"fmt"
	"gquery_builder/builder"
)

type Query = builder.Query
type Builder = builder.Builder

func main() {
	q := Query{}
	var b Builder = q.From("users").Select("name", "age", "year").Where("id = 1")
	builder.PrintSQL(b)

	b = q.CreateTable("users")
	builder.PrintSQL(b)

	b = q.DropDatabase("users_database")
	builder.PrintSQL(b)

	b = q.CreateTable("users_table")
	builder.PrintSQL(b)

	b = q.InsertInto("users").Into("sifat", "age", "year")
	fmt.Println(b.Build())

}
