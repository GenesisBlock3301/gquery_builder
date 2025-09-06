package main

import (
	"fmt"
	"github.com/GenesisBlock3301/gquery_builder/builder"
)

func main() {
	q := builder.Query{}
	var b builder.Builder = q.From("customers").Select("*").Where("id = 1")
	fmt.Println(b.Build())
}
