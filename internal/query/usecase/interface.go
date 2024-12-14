package usecase

type Provider interface {
	SelectQuery() (string, error)
	CheckQueryExitByMsg(string) (bool, error)
	InsertQuery(string) error
}
