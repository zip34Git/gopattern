// Set of program pattern for gogen
package pattern

import (
// "fmt"
)

// Observer
type Observer interface {
	Update() error
	//UpdatePrmtr(map[string]string) error
}

type Observerable interface {
	RegisterObserver(*Observer) error
	RemoveObserver(*Observer) error
	NotifyObservers() error
}

