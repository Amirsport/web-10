package api

type Usecase interface {
	FetchQuery() (string, error)
	SetQuery(msg string) error
}
