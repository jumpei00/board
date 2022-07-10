package params

type CreateCommentDomainLayerParam struct {
	ThreadKey string
	Comment     string
	Contributor string
}

type EditCommentDomainLayerParam struct {
	Comment     string
}