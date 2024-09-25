package userErrors

type BusinessError struct {
	Message string
	Status  int
}

func (e *BusinessError) Error() string {
	return e.Message
}
