package actions

// Service struct handles action business logic tasks.
type Service struct {
	repository Repository
}

// ListActions contains the business logic to retrieve all actions
func (svc *Service) ListActions() ([]Action, error) {
	return svc.repository.ListActions()
}

// Get contains the business logic to retrieve specific action
func (svc *Service) Get(actionId string) (Action, error) {
	return svc.repository.Get(actionId)
}

// New creates a new service struct
func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}
