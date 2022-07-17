package application

import (
	appParams "github.com/jumpei00/board/backend/app/application/params"
	"github.com/jumpei00/board/backend/app/domain"
	domainParams "github.com/jumpei00/board/backend/app/domain/params"
	"github.com/jumpei00/board/backend/app/domain/repository"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
)

type CommentApplication interface {
	GetAllByThreadKey(threadKey string) (*[]domain.Comment, error)
	CreateComment(param *appParams.CreateCommentAppLayerParam) (*domain.Comment, error)
	EditComment(param *appParams.EditCommentAppLayerParam) (*domain.Comment, error)
	DeleteComment(param *appParams.DeleteCommentAppLayerParam) error
}

type commentApplication struct {
	userRepo    repository.UserRepository
	threadRepo  repository.ThreadRepository
	commentRepo repository.CommentRepository
}

func NewCommentApplication(ur repository.UserRepository, tr repository.ThreadRepository, cr repository.CommentRepository) *commentApplication {
	return &commentApplication{
		userRepo:    ur,
		threadRepo:  tr,
		commentRepo: cr,
	}
}

func (c *commentApplication) GetAllByThreadKey(threadKey string) (*[]domain.Comment, error) {
	comments, err := c.commentRepo.GetAllByKey(threadKey)
	if err != nil {
		return comments, err
	}

	return comments, nil
}

func (c *commentApplication) CreateComment(param *appParams.CreateCommentAppLayerParam) (*domain.Comment, error) {
	// 対象のスレッドを取得
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	user, err := c.userRepo.GetByID(param.UserID)
	if err != nil {
		return nil, err
	}

	domainParam := domainParams.CreateCommentDomainLayerParam{
		ThreadKey:   param.ThreadKey,
		Contributor: user.Username,
		Comment:     param.Comment,
	}

	comment := domain.NewComment(&domainParam)

	newComment, err := c.commentRepo.Insert(comment)
	if err != nil {
		return nil, err
	}

	// スレッドの更新日時とコメント数を更新する
	thread.UpdateLatestUpdatedDate()
	thread.CountupCommentSum()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return nil, err
	}

	return newComment, nil
}

func (c *commentApplication) EditComment(param *appParams.EditCommentAppLayerParam) (*domain.Comment, error) {
	// 対象のスレッドを取得
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	// 対象のコメントを取得
	comment, err := c.commentRepo.GetByKey(param.CommentKey)
	if err != nil {
		return nil, err
	}

	user, err := c.userRepo.GetByID(param.UserID)
	if err != nil {
		return nil, err
	}

	// 同じ投稿者でなければ編集することはできない
	if comment.IsNotSameContritubor(user.Username) {
		logger.Warning(
			"comment conributor and edit contributor is not same",
			"comment_contributor", comment.GetContributor(),
			"requesting_contiributor", user.Username,
		)
		return nil, appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but edit requesting user is %s", comment.GetContributor(), user.Username,
		)
	}

	domainParam := domainParams.EditCommentDomainLayerParam{
		Comment: param.Comment,
	}

	updatedComment := comment.UpdateComment(&domainParam)

	insertedComment, err := c.commentRepo.Insert(updatedComment)
	if err != nil {
		return nil, err
	}

	// スレッドの更新時刻を更新する
	thread.UpdateLatestUpdatedDate()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return nil, err
	}

	return insertedComment, nil
}

func (c *commentApplication) DeleteComment(param *appParams.DeleteCommentAppLayerParam) error {
	// スレッドはあるかどうか確認
	thread, err := c.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return err
	}

	comment, err := c.commentRepo.GetByKey(param.CommentKey)
	if err != nil {
		return err
	}

	user, err := c.userRepo.GetByID(param.UserID)
	if err != nil {
		return err
	}

	// 同じ投稿者でなければ削除することはできない
	if comment.IsNotSameContritubor(user.Username) {
		logger.Warning(
			"comment conributor and delete contributor is not same",
			"comment_contributor", comment.GetContributor(),
			"requesting_contiributor", user.Username,
		)
		return appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but delete requesting user is %s", comment.GetContributor(), user.Username,
		)
	}

	if err := c.commentRepo.Delete(comment); err != nil {
		return err
	}

	// スレッドの更新時刻とコメント数を更新
	thread.UpdateLatestUpdatedDate()
	thread.CountdownCommentSum()
	if _, err := c.threadRepo.Update(thread); err != nil {
		return err
	}

	return nil
}
