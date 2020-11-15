package exceptions

import "fmt"

// Non-recoverable exeption - external service or database failure
type DependencyFailureException struct {
	OriginalError error
}

func (exception *DependencyFailureException) Error() string {
	return fmt.Sprintf("Dependency failure exception: %v", exception.OriginalError)
}

type InvalidArgumentException struct {
	OriginalError error
}

func (exception *InvalidArgumentException) Error() string {
	return fmt.Sprintf("Invalid Argument exception: %v", exception.OriginalError)
}

// Recoverable exception like database timeout
type RecoverableException struct {
	OriginalError error
}

func (exception *RecoverableException) Error() string {
	return fmt.Sprintf("Recoverable exception: %v", exception.OriginalError)
}
