package service

import (
	"github.com/electronlabs/vibes-api/actions/model"
	"github.com/electronlabs/vibes-api/actions/repository"
)

// Service struct handles action business logic tasks.
type Service struct {
	repository repository.Repository
}

// GetActions contains the business logic to retrieve all actions
func (svc *Service) GetActions() ([]model.Action, error) {
	return svc.repository.GetAll()
}

// New creates a new service struct
func New(repository repository.Repository) *Service {
	return &Service{repository: repository}
}
