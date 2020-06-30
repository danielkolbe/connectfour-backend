package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/danielkolbe/connectfour/game"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type GameServiceMock struct {
    mock.Mock
}

func (mock GameServiceMock) Turn(column int, gameId string) error {
    fmt.Printf("column: %v gameId: %v", column, gameId)
    args := mock.Called(column, gameId)
    
    return args.Error(0)
}

var csSwap game.GameService
var h TurnHandler 
var cookie *http.Cookie

func setup () {
    gameServiceMock := GameServiceMock{} 
    gameServiceMock.On("Turn", 4,"324234-555").Return(nil);
    h = NewTurnHandler(gameServiceMock)
    cookie = &http.Cookie{Name: "gameId", Value: "324234-555"}
}

func TestTurnHandler(t *testing.T) {
    // Arrange
    setup ()    
    req, _ := http.NewRequest("", "/turn?column=4", nil)
    req.AddCookie(cookie)
    rr := httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("Wrong http status code returned"))
}