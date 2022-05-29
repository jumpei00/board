package params

import "time"

type CreateThreadAppLayerParam struct {
	Title       string
	Contributor string
}

type EditThreadAppLayerParam struct {
	ThreadKey   string
	Title       string
	Contributor string
}

type DeleteThreadAppLayerParam struct {
	ThreadKey   string
	Contributor string
}

type CreateThreadDomainLayerParam struct {
	Title       string
	Contributor string
}

type EditThreadDomainLayerParam struct {
	ThreadKey   string
	Title       string
	Contributor string
	UpdateDate    time.Time
	Views       int
	SumComment  int
}
