package params


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
	Title       string
}
