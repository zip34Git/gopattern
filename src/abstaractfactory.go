package pattern

import (

)

type AbstractFactory interface {
	Create() error
	CreateFrom(prmtr map[string]string) error
}