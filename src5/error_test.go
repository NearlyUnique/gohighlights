package interfaces

import (
	"fmt"
	"testing"
)

func Test_custom_error(t *testing.T) {
	err := eatPorridge(42)

	if err != nil {
		t.Errorf("eatPorridge() failed: %v", err)
	}
}

func eatPorridge(id int) error {
	// if ok
	// return nil //no error
	// but we are going to fail
	return &MyDomainError{
		ID:  id,
		Cat: Hot,
	}
}

type (
	MyDomainError struct {
		ID  int
		Cat Category
	}
	Category string
)

const (
	Cold      Category = "too_cold"
	Hot                = "too_hot"
	JustRight          = "just_right"
)

func (e *MyDomainError) Error() string {
	return fmt.Sprintf("cannot consume, #%d [%s]", e.ID, e.Cat)
}
