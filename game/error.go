package game

import "fmt"

type BoardDoesNotExistError struct {
	message string
}

func NewBoardDoesNotExistError(gameID string) error {
	return &BoardDoesNotExistError{fmt.Sprintf("board with gameID <%v> does not exist", gameID)}
}

func (e *BoardDoesNotExistError) Error() string {
	return e.message
}


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