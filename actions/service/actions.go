package service

import (
	"github.com/electronlabs/vibes-api/actions/model"
	"github.com/electronlabs/vibes-api/actions/repository"
)

type ActionsService struct {
	repository repository.ActionsRepository
}

func (s *ActionsService) GetActions() ([]model.Action, error) {
	return s.repository.GetAll()
}

func NewActionsService(repository repository.ActionsRepository) *ActionsService {
	return &ActionsService{repository: repository}
}
