package domain

type NoPacksError struct{}

func (e *NoPacksError) Error() string {
	return "error: cannot start application without at least one available pack"
}

func NewNoPacksError() error {
	return &NoPacksError{}
}

type CannotCalculateOrderItemsError struct{}

func (e *CannotCalculateOrderItemsError) Error() string {
	return "error: cannot calculate order items"
}
