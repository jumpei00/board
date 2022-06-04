package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	"github.com/jumpei00/board/backend/app/params"
)

type CommentApplication interface {
	GetAllByThreadKey(threadKey string) ([]*domain.Comment, error)
	CreateComment(param *params.CreateCommentAppLayerParam) ([]*domain.Comment, error)
	EditComment(param *params.EditCommentAppLayerParam) ([]*domain.Comment, error)
	DeleteComment(param *params.DeleteCommentAppLayerParam) ([]*domain.Comment, error)
}

type commentApplication struct {
	threadRepo  repository.ThreadRepository
	commentRepo repository.CommentRepository
}

func NewCommentApplication(tr repository.ThreadRepository, cr repository.CommentRepository) *commentApplication {
	return &commentApplication{
		threadRepo: tr,
		commentRepo: cr,
	}
}

func (c *commentApplication) GetAllByThreadKey(threadKey string) ([]*domain.Comment, error) {
	comments, err := c.commentRepo.GetAllByKey(threadKey)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentApplication) CreateComment(param *params.CreateCommentAppLayerParam) ([]*domain.Comment, error) {
	// コメントをつけようとしているスレッドはあるか？
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	domainParam := params.CreateCommentDomainLayerParam{
		ThreadKey: param.ThreadKey,
		Contributor: param.Contributor,
		Comment: param.Comment,
	}

	newComment := domain.NewComment(&domainParam)

	if _, err := c.commentRepo.Insert(newComment); err != nil {
		return nil, err
	}

	// スレッドの更新日時とコメント数を更新する
	thread.UpdateLatestUpdateDate()
	thread.CountupSumComment()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return nil, err
	}

	comments, err := c.commentRepo.GetAllByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentApplication) EditComment(param *params.EditCommentAppLayerParam) ([]*domain.Comment, error) {
	// スレッドはあるかどうか確認
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	comment, err := c.commentRepo.GetByKey(param.CommentKey)
	if err != nil {
		return nil, err
	}

	// 同じ投稿者でなければ編集することはできない
	// TODO: errorを変える
	if comment.IsNotSameContritubor(param.Contributor) {
		return nil, err
	}

	domainParam := params.EditCommentDomainLayerParam{
		Comment: param.Comment,
	}

	updatedComment := comment.UpdateComment(&domainParam)
	if _, err := c.commentRepo.Insert(updatedComment); err != nil {
		return nil, err
	}

	// スレッドの更新時刻を更新する
	thread.UpdateLatestUpdateDate()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return nil, err
	}

	comments, err := c.commentRepo.GetAllByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (c *commentApplication) DeleteComment(param *params.DeleteCommentAppLayerParam) ([]*domain.Comment, error) {
	// スレッドはあるかどうか確認
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	comment, err := c.commentRepo.GetByKey(param.CommentKey)
	if err != nil {
		return nil, err
	}

	// 同じ投稿者でなければ削除することはできない
	// TODO: errorを変える
	if comment.IsNotSameContritubor(param.Contributor) {
		return nil, err
	}

	if err := c.commentRepo.Delete(comment); err != nil {
		return nil, err
	}

	// スレッドの更新時刻とコメント数を更新
	thread.UpdateLatestUpdateDate()
	thread.CountdownSumComment()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return nil, err
	}

	comments, err := c.commentRepo.GetAllByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	return comments, nil
}
