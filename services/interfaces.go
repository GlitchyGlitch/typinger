package services

type ErrHandler interface {
	Error(string) error
}
