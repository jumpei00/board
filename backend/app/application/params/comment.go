package params

type CreateCommentAppLayerParam struct {
	ThreadKey string
	Comment   string
	UserID    string
}

type EditCommentAppLayerParam struct {
	ThreadKey  string
	CommentKey string
	Comment    string
	UserID     string
}

type DeleteCommentAppLayerParam struct {
	ThreadKey  string
	CommentKey string
	UserID     string
}
