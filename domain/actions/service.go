package actions

// Service struct handles action business logic tasks.
type Service struct {
	repository Repository
}

// List contains the business logic to retrieve all actions
func (svc *Service) List() ([]Action, error) {
	return svc.repository.List()
}

// New creates a new service struct
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
