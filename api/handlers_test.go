package api

import (
	"bytes"
	"encoding/json"
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

func (mock *GameServiceMock) Turn(column int, gameID string) error {
    args := mock.Called(column, gameID)
    return args.Error(0)
}

var csSwap game.GameService
var h TurnHandler 
var cookie *http.Cookie

func setup () {
    gameServiceMock := GameServiceMock{} 
    gameServiceMock.On("Turn", 4,"324234-555").Return(nil);
    gameServiceMock.On("Turn", 3,"324234-555").Return(fmt.Errorf("error"));
    h = NewTurnHandler(&gameServiceMock)
    cookie = &http.Cookie{Name: "gameId", Value: "324234-555"}
}

func TestTurnHandler(t *testing.T) {
    // Arrange
    setup()    
    body := struct {Column int}{4}
    bytesBody,_ := json.Marshal(body)
    req, _ := http.NewRequest("", "/turn", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr := httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
    
    // Arrange
    body.Column = -1
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "/turn", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if column number is negative"))

    // Arrange
    wrongBody := struct {Unknown int}{4}
    bytesBody,_ = json.Marshal(wrongBody)
    req, _ = http.NewRequest("", "/turn", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if body does not contain column field"))

    // Arrange
    body.Column = 3
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "/turn", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if error if game service returns error"))
}