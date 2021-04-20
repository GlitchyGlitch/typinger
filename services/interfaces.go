package services

type tokenController interface {
	ParseAuthorization(string) (map[string]interface{}, error)
	Token(string) (string, error)
}
