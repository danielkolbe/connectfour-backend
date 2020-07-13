package game

import "fmt"

//BoardDoesNotExistError implements the error interface
type BoardDoesNotExistError struct {
	message string
}

func NewBoardDoesNotExistError(gameID string) error {
	return &BoardDoesNotExistError{fmt.Sprintf("board with gameID <%v> does not exist", gameID)}
}

func (e *BoardDoesNotExistError) Error() string {
	return e.message
}

//MatchIsOverError implements the error interface
type MatchIsOverError struct {
	message string
}

func NewMatchIsOverError() error {
	return &MatchIsOverError{fmt.Sprintf("match has already a winner")}
}

func (e *MatchIsOverError) Error() string {
	return e.message
}

//ColumnIsFullError implements the error interface
type ColumnIsFullError struct {
	message string
}

func NewColumnIsFullError(column int) error {
	return &ColumnIsFullError{fmt.Sprintf("column %v is full", column)}
}

func (e *ColumnIsFullError) Error() string {
	return e.message
}

//ColumnIsOutOfBoundsError implements the error interface
type ColumnIsOutOfBoundsError struct {
	message string
}

func NewColumnIsOutOfBoundsError(column int) error {
	return &ColumnIsOutOfBoundsError{fmt.Sprintf("column %v is out of bounds: 0-%v", column, nCols-1)}
}

func (e *ColumnIsOutOfBoundsError) Error() string {
	return e.message
}