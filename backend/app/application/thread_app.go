package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/jumpei00/board/backend/app/params"
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
}

func NewThreadApplication(tr repository.ThreadRepository) *threadApplication {
	return &threadApplication{
		threadRepo: tr,
	}
}

func (t *threadApplication) GetAllThread() (*[]domain.Thread, error) {
	threads, err := t.threadRepo.GetAll()
	if err != nil {
		return threads, err
	}

	return threads, nil
}

func (t *threadApplication) GetByThreadKey(threadKey string) (*domain.Thread, error) {
	thread, err := t.threadRepo.GetByKey(threadKey)
	if err != nil {
		return thread, err
	}

	return thread, nil
}

func (t *threadApplication) CreateThread(param *params.CreateThreadAppLayerParam) (*domain.Thread, error) {
	domainParam := params.CreateThreadDomainLayerParam{
		Title:       param.Title,
		Contributor: param.Title,
	}

	newThread := domain.NewThread(&domainParam)

	thread, err := t.threadRepo.Insert(newThread)
	if err != nil {
		return thread, err
	}

	return thread, nil
}

func (t *threadApplication) EditThread(param *params.EditThreadAppLayerParam) (*domain.Thread, error) {
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return thread, err
	}

	if thread.IsNotSameContributor(param.Contributor) {
		logger.Warning("thread contibutor is %s, but edit thread requesting user is %s", thread.GetContributor(), param.Contributor)
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
		return editedThread, err
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
		logger.Warning("thread contibutor is %s, but delete thread requesting user is %s", thread.GetContributor(), param.Contributor)
		return appError.NewErrBadRequest(
			appError.Message().NotSameContributor,
			"contibutor is %s, but delete requesting user is %s", thread.GetContributor(), param.Contributor,
		)
	}

	// TODO: comments側が実装された後でこちらを実装する
	if err := t.threadRepo.Delete(thread); err != nil {
		return err
	}

	return nil
}
