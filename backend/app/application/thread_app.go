package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
	"github.com/pkg/errors"
)

type ThreadApplication interface {
	GetAllThread() (*[]domain.Thread, error)
	GetByThreadKey(threadKey string) (*domain.Thread, error)
	CreateThread(param *params.CreateThreadAppLayerParam) (*domain.Thread, error)
	EditThread(param *params.EditThreadAppLayerParam) (*domain.Thread, error)
	DeleteThread(param *params.DeleteThreadAppLayerParam) error
}

type threadApplication struct {
	threadRepo repository.ThreadRepository
	commentRepo repository.CommentRepository
}

func NewThreadApplication(tr repository.ThreadRepository, cr repository.CommentRepository) *threadApplication {
	return &threadApplication{
		threadRepo: tr,
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

func (t *threadApplication) CreateThread(param *params.CreateThreadAppLayerParam) (*domain.Thread, error) {
	domainParam := params.CreateThreadDomainLayerParam{
		Title:       param.Title,
		Contributor: param.Contributor,
	}

	newThread := domain.NewThread(&domainParam)

	thread, err := t.threadRepo.Insert(newThread)
	if err != nil {
		return nil, err
	}

	return thread, nil
}

func (t *threadApplication) EditThread(param *params.EditThreadAppLayerParam) (*domain.Thread, error) {
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return nil, err
	}

	if thread.IsNotSameContributor(param.Contributor) {
		logger.Warning(
			"thread contibutor and edit requesting contributor is not same",
			"thread_contributor", thread.GetContributor(),
			"requesting_contributor", param.Contributor,
		)
		return nil, appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but edit requesting user is %s", thread.GetContributor(), param.Contributor,
		)
	}

	domainParam := params.EditThreadDomainLayerParam{
		Title:       param.Title,
	}

	editedThread, err := t.threadRepo.Update(thread.UpdateThread(&domainParam))
	if err != nil {
		return nil, err
	}

	return editedThread, nil
}

func (t *threadApplication) DeleteThread(param *params.DeleteThreadAppLayerParam) error {
	// スレッドが削除された場合は、それに紐づくコメントも全て削除させる必要がある
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return err
	}

	if thread.IsNotSameContributor(param.Contributor) {
		logger.Warning(
			"thread contibutor and delete requesting contributor is not same",
			"thread_contributor", thread.GetContributor(),
			"requesting_contributor", param.Contributor,
		)
		return appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but delete requesting user is %s", thread.GetContributor(), param.Contributor,
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
