package params

type CreateThreadAppLayerParam struct {
	Title  string
	UserID string
}

type EditThreadAppLayerParam struct {
	ThreadKey string
	Title     string
	UserID    string
}

type DeleteThreadAppLayerParam struct {
	ThreadKey string
	UserID    string
}
