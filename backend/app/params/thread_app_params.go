package params

type CreateThreadAppLayerParam struct {
	Title       string
	Contributer string
}

type EditThreadAppLayerParam struct {
	ThreadKey   string
	Title       string
	Contributer string
}

type DeleteThreadAppLayerParam struct {
	ThreadKey   string
	Contributer string
}
