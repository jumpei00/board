package application

import (
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/jumpei00/board/backend/app/domain/repository"
	"github.com/jumpei00/board/backend/app/params"
)

type ThreadApplication interface {
	GetAllThread() ([]*domain.Thread, error)
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

func (t *threadApplication) GetAllThread() ([]*domain.Thread, error) {
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
		Contributor: param.Title,
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

	if thread.IsNotSameContritubor(param.Contributor) {
		return nil, err
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
	thread, err := t.threadRepo.GetByKey(param.ThreadKey)
	if err != nil {
		return err
	}

	if thread.IsNotSameContritubor(param.Contributor) {
		return err
	}

	if err := t.threadRepo.Delete(thread); err != nil {
		return err
	}

	return nil
}
