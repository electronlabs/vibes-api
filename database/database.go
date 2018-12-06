package database

// Database interface defines basic behavior to interact with a database.
type Database interface {
	Connect(uri string) (interface{}, error)
	Disconnect(client interface{}) error
}
