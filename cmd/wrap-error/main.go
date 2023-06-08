package main

import (
	"errors"
	"fmt"
)

type myerror struct {
	err error
}

func NewErr(err error) error {
	if err == nil {
		return nil
	}
	return &myerror{err}
}
func (e myerror) Error() string {
	if e.err == nil {
		return "retryable: <nil>"
	}
	return "retryable: " + e.err.Error()
}
func (e myerror) Unwrap() error {
	return e.err
}

func main() {
	ee := errors.New("最底层error")
	w := fmt.Errorf("wrap了一个错误: %w", ee)
	e := NewErr(w)
	fmt.Println("剥离wrap前:", w)
	fmt.Println("剥离wrap后:", errors.Unwrap(w))
	fmt.Println("再次剥离wrap后:", errors.Unwrap(errors.Unwrap(w)))
	eee := &myerror{}
	if errors.As(e, &eee) {
		fmt.Println("my error:", e)
		fmt.Println("my error 剥离后:", eee.Unwrap())
	}

}
