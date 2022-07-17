package application

import (
	appParams "github.com/jumpei00/board/backend/app/application/params"
	"github.com/jumpei00/board/backend/app/domain"
	domainParams "github.com/jumpei00/board/backend/app/domain/params"
	"github.com/jumpei00/board/backend/app/domain/repository"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
)

type ThreadApplication interface {
	GetAllThread() (*[]domain.Thread, error)
	GetByThreadKey(threadKey string) (*domain.Thread, error)
	CreateThread(param *appParams.CreateThreadAppLayerParam) (*domain.Thread, error)
	EditThread(param *appParams.EditThreadAppLayerParam) (*domain.Thread, error)
	DeleteThread(param *appParams.DeleteThreadAppLayerParam) error
}

type threadApplication struct {
	userRepo    repository.UserRepository
	threadRepo  repository.ThreadRepository
	commentRepo repository.CommentRepository
}

func NewThreadApplication(ur repository.UserRepository, tr repository.ThreadRepository, cr repository.CommentRepository) *threadApplication {
	return &threadApplication{
		userRepo:    ur,
		threadRepo:  tr,
		commentRepo: cr,
	}
}

func (t *threadApplication) GetAllThread() (*[]domain.Thread, error) {
	threads, err := t.threadRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return threads, nil
}

func (t *threadApplication) GetByThreadKey(threadKey string) (*domain.Thread, error) {
	thread, err := t.threadRepo.GetByKey(threadKey)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (t *threadApplication) CreateThread(param *appParams.CreateThreadAppLayerParam) (*domain.Thread, error) {
	user, err := t.userRepo.GetByID(param.UserID)
	if err != nil {
		return nil, err
	}

	domainParam := domainParams.CreateThreadDomainLayerParam{
		Title:       param.Title,
		Contributor: user.Username,
	}

	newThread := domain.NewThread(&domainParam)

	thread, err := t.threadRepo.Insert(newThread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (t *threadApplication) EditThread(param *appParams.EditThreadAppLayerParam) (*domain.Thread, error) {
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	user, err := t.userRepo.GetByID(param.UserID)
	if err != nil {
		return nil, err
	}

	if thread.IsNotSameContributor(user.Username) {
		logger.Warning(
			"thread contibutor and edit requesting contributor is not same",
			"thread_contributor", thread.GetContributor(),
			"requesting_contributor", user.Username,
		)
		return nil, appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but edit requesting user is %s", thread.GetContributor(), user.Username,
		)
	}

	domainParam := domainParams.EditThreadDomainLayerParam{
		Title: param.Title,
	}

	editedThread, err := t.threadRepo.Update(thread.UpdateThread(&domainParam))
	if err != nil {
		return nil, err
	}

	return editedThread, nil
}

func (t *threadApplication) DeleteThread(param *appParams.DeleteThreadAppLayerParam) error {
	// スレッドが削除された場合は、それに紐づくコメントも全て削除させる必要がある
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return err
	}

	user, err := t.userRepo.GetByID(param.UserID)
	if err != nil {
		return err
	}

	if thread.IsNotSameContributor(user.Username) {
		logger.Warning(
			"thread contibutor and delete requesting contributor is not same",
			"thread_contributor", thread.GetContributor(),
			"requesting_contributor", user.Username,
		)
		return appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but delete requesting user is %s", thread.GetContributor(), user.Username,
		)
	}

	comments, err := t.commentRepo.GetAllByKey(param.ThreadKey)

	// スレッドに対応したコメントが無い場合は削除が必要ない
	if errors.Cause(err) == appError.ErrNotFound {
		if err := t.threadRepo.Delete(thread); err != nil {
			return err
		}
		return nil
	}

	// 何らかのエラーが発生した場合は何もせずリターンさせる
	if err != nil {
		return err
	}

	// コメントが正常に取得できている場合はスレッドとそれに紐付くコメントを削除しなければいけない
	if err := t.threadRepo.DeleteThreadAndComments(thread, comments); err != nil {
		return err
	}

	return nil
}
