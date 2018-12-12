package actions

// Repository provides an abstraction on top of the action data source
type Repository interface {
	List() ([]Action, error)
}
