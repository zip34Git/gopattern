package pattern

import (

)

type Job interface {
	Begin() int
	End()	int
	Id() int
	SetId(int)
}

type ResultJob interface {
	Result() int
}
// pattern Mapper - mapping data on workers
type Mapper interface {
	Data2Map()
	MapSize() int
}

