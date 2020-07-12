package game

import "fmt"

type MatchIsOverError struct {
	message string
}

func NewMatchIsOverError() error {
	return &MatchIsOverError{fmt.Sprintf("match has already a winner")}
}

func (e *MatchIsOverError) Error() string {
	return e.message
}

type ColumnIsFullError struct {
	message string
}

func NewColumnIsFullError(column int) error {
	return &ColumnIsFullError{fmt.Sprintf("column %v is full", column)}
}

func (e *ColumnIsFullError) Error() string {
	return e.message
}