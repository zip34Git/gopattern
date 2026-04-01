package gopattern

import (
// "fmt"
)

// WorkerPoll pattern, sigle data map to multiplay workers SDSI - MDMI
// Runner() error
// RunnerFrom(prmtr map[string]string) error
type WorkerPoll interface {
	Runner() error
	RunnerFrom(prmtr map[string]string) error
}