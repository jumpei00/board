package params

import "time"

type ThreadCreateDomainLayerParam struct {
	Title       string
	Contributer string
}

type ThreadEditDomainLayerParam struct {
	ThreadKey   string
	Title       string
	Contributer string
	PostDate    time.Time
	Views       int
	SumComment  int
}
