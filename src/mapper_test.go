package gopattern

import (
	//"errors"
	//"fmt"
	//"slices"
	"log"
	"strconv"
	"testing"
)

// CheckMapperJob - release Job intrface
type CheckMapperJob struct {
	begin int
	end   int
	id int
}

func (c CheckMapperJob) Begin() int {
	return c.begin
}

func (c CheckMapperJob) End() int {
	return c.end
}

func (c CheckMapperJob) Id() int {
	return c.id
}

func (c *CheckMapperJob) SetId(pid int) {
	c.id = pid
}



type CheckMapper struct {
	minSize int
	numJob  int
	size    int
	mapJob  []Job
}

// Data2Map() release method Mapper interface
func (c *CheckMapper) Data2Map() {
	c.mapJob = make([]Job, 0, 20)

	n := c.size / c.minSize
	nRound := c.size % c.minSize
	k := c.minSize + nRound/n
	if n >= 1 {
		i := 0
		l := 0
		for i = range (n-1) {
			l = i*k
			var oJ Job = &CheckMapperJob{l, l + k,0}
			c.mapJob = append(c.mapJob, oJ)
		}
		var oJ Job = &CheckMapperJob{l, k+(nRound % n),0} // last job + size + remainder
		c.mapJob = append(c.mapJob, oJ)
	} else { // size < minSize
		var oJ Job = &CheckMapperJob{0, c.size,0}
		c.mapJob = append(c.mapJob, oJ)
	}

}

func (c CheckMapper) MapSize() int {
	return len(c.mapJob)
}

func (c *CheckMapper) SetId() int {
	return len(c.mapJob)
}

var tstPrmtr1 = map[string]int{
	"MinSize": 10000,
	"NumJob":  10,
	"Size":    10000,
	"MapSize": 1,
}

var tstPrmtr2 = map[string]int{"MinSize": 10000,
	"NumJob":  10,
	"Size":    30005,
	"MapSize": 3,
}

func Test_Mapper(t *testing.T) {

	t.Run("Mapper", func(t *testing.T) {
		var errMsg string = "ERROR in Mapper"
		var tData []map[string]int
		tData = append(tData, tstPrmtr1, tstPrmtr2)
		for i, d := range tData {

			var p Mapper = &CheckMapper{minSize: d["MinSize"], numJob: d["NumJob"], size: d["Size"]}
			p.Data2Map()
			if d["MapSize"] != p.MapSize() {
				t.Errorf("Expected error message for `%s`, got `%s`",
					strconv.Itoa(i), errMsg)
			}
			log.Print(strconv.Itoa(d["MinSize"])+" "+strconv.Itoa(d["NumJob"])+ " "+ strconv.Itoa(d["Size"]))
			log.Println("->len Mapper is ",p.MapSize())
		}
	})
	log.Println("Test_Mapper is Ok")

}
