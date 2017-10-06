package greeter

import (
	"fmt"
)

// Service greets name
type Service interface {
	Greet(name string) string
}

type greet string

func (g greet) Greet(name string) string {
	return fmt.Sprintf(string(g), name)
}

// New service
func New() Service {
	return greet("Hello %s!")
}
