package pattern

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"testing"
)

// LogObserver - Log Update
type LogFileObserver struct {
	fname string
}

func (l LogFileObserver) Update() error {
	fmt.Println(l.fname)
	return nil
}

func (l LogFileObserver) UpdatePrmtr(prmtr map[string]string) error {
	return nil
}

func (l LogFileObserver) Close() {
}

func NewLogFileObserver(fname string) (*LogFileObserver, error) {
	l := &LogFileObserver{fname}

	return l, nil
}

// GogenObserverable - notify object
type GogenObserverable struct {
	lO []*Observer
	lM map[string][]*Observer
}

// RegisterObserver - method interface Observerable
func (g *GogenObserverable) RegisterObserver(o *Observer) error {

	if slices.Contains(g.lO, o) {
		return errors.New("Observer already exists")
	}

	g.lO = append(g.lO, o)
	return nil
}

// RemoveObserver - method interface Observerable
func (g *GogenObserverable) RemoveObserver(o *Observer) error {
	i := slices.Index(g.lO, o)
	if i >= 0 {
		g.lO = slices.Delete(g.lO, i, i+1)
	}
	return nil
}

// NotifyObservers - method interface Observerable
func (g GogenObserverable) NotifyObservers() error {
	for _, v := range g.lO {
		(*v).Update()
	}
	return nil
}

// Save(s string) - method Memento
// s - key of map saving array *Observers
func (g *GogenObserverable) Save(s string) error {
	var tsave = []*Observer{}
	tsave = append(tsave, g.lO...)
	g.lM[s] = tsave
	return nil
}

// Restore(s string) - method Memento
// s - key of map savers
func (g *GogenObserverable) Restore(s string) error {
	if _, ok := g.lM[s]; ok {
		g.lO = g.lM[s]
	} else {
		return errors.New("ключ не найден в map ")
	}
	return nil
}

func NewGogenObserverable() (*GogenObserverable, error) {
	G := &GogenObserverable{}
	G.lM = map[string][]*Observer{}
	return G, nil
}

func Test_Observer(t *testing.T) {
	var oL Observerable
	tL, err := NewGogenObserverable()
	oL = tL // check type on compiler time
	if err != nil {
		log.Fatal(err)
	}

	tstStr := []string{"pattern.log", "test.1", "test.2"}
	for _, v := range tstStr {
		var e Observer
		e, err := NewLogFileObserver(v)
		if err != nil {
			log.Fatal(errors.New("Open NewLogFileObser error"))
		}
		tL.RegisterObserver(&e)
	}
	oL.NotifyObservers()
	var m Memento = tL
	mkey := "full"
	m.Save(mkey)
	oL.RemoveObserver(tL.lO[0])
	oL.NotifyObservers()
	m.Restore(mkey)
	fmt.Println("->after restore")
	oL.NotifyObservers()
	var tOtst Observer
	tOtst, err = NewLogFileObserver("test.3")
	if err != nil {
		log.Fatal(errors.New("Open NewLogFileObser error"))
	}
	oL.RemoveObserver(&tOtst)
	oL.NotifyObservers()
}

