package pattern

import (
// "fmt"
)

type Memento interface {
	Save(string) error
	Restore(string) error
}