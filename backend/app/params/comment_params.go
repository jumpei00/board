package params

type CreateCommentAppLayerParam struct {
	ThreadKey   string
	Comment     string
	Contributor string
}

type EditCommentAppLayerParam struct {
	ThreadKey   string
	CommentKey  string
	Comment     string
	Contributor string
}

type DeleteCommentAppLayerParam struct {
	ThreadKey   string
	CommentKey  string
	Contributor string
}

type CreateCommentDomainLayerParam struct {
	ThreadKey string
	Comment     string
	Contributor string
}

type EditCommentDomainLayerParam struct {
	Comment     string
}