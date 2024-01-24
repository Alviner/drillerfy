package querier

// Querier is the interface that wraps the basic methods to create a dialect
// specific query.
type Querier interface {
	CreateDatabase(name, temlate string) string
	DeleteDatabase(name string) string
	DisconnectFromDatabase(name string) string
}
