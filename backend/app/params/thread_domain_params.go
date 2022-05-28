package params

import "time"

type ThreadCreateDomainLayerParam struct {
	Title       string
	Contributor string
}

type ThreadEditDomainLayerParam struct {
	ThreadKey   string
	Title       string
	Contributor string
	PostDate    time.Time
	Views       int
	SumComment  int
}
