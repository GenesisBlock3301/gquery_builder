package builder

import "fmt"

// Builder interface
type Builder interface {
	Build() string
}

// PrintSQL accepts any builder
func PrintSQL(b Builder) {
	fmt.Println(b.Build())
}
