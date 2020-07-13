package turn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/require"
    "github.com/danielkolbe/connectfour/game"
)

type GameServiceMock struct {
    mock.Mock
}

func (mock *GameServiceMock) Turn(column int, gameID string) error {
    args := mock.Called(column, gameID)
    return args.Error(0)
}

func (mock *GameServiceMock) Board(gameID string) game.Board {
    fmt.Println("Dummy method, please don't call")
    return game.Board{}
}

func (mock *GameServiceMock) Winner(gameID string) (string, error) {
	fmt.Println("Dummy method, please don't call")
	return "" , nil
}

var h http.Handler
var cookie *http.Cookie

func setup () {
    gameServiceMock := GameServiceMock{}
    gameID := func(w http.ResponseWriter, req *http.Request) string {
        c, _ := req.Cookie("gameID")
        return c.Value
    }
    gameServiceMock.On("Turn", 4,"324234-555").Return(nil);
    gameServiceMock.On("Turn", 3,"324234-555").Return(fmt.Errorf("error"));
    gameServiceMock.On("Turn", 5,"324234-555").Return(game.NewColumnIsFullError(5));
    gameServiceMock.On("Turn", 6,"324234-555").Return(game.NewMatchIsOverError());
    gameServiceMock.On("Turn", 2,"324234-555").Panic("panic!")
    h = NewHandler(&gameServiceMock, gameID)
    cookie = &http.Cookie{Name: "gameID", Value: "324234-555"}
}

func TestHandler(t *testing.T) {
    // Arrange
    setup()    
    body := struct {Column int}{4}
    bytesBody,_ := json.Marshal(body)
    req, _ := http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr := httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
    
    // Arrange
    setup()    
    body = struct {Column int}{5}
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusConflict, rr.Code, fmt.Sprintf("should return http 409 if column is full"))
    
    // Arrange
    setup()    
    body = struct {Column int}{6}
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusConflict, rr.Code, fmt.Sprintf("should return http 409 if column is full"))

    // Arrange
    body.Column = -1
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if column number is negative"))

    // Arrange
    wrongBody := struct {Unknown int}{4}
    bytesBody,_ = json.Marshal(wrongBody)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if body does not contain column field"))

    // Arrange
    body.Column = 3
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if error if game service returns error"))

     // Arrange
     body.Column = 2
     bytesBody,_ = json.Marshal(body)
     req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
     req.AddCookie(cookie)
     rr = httptest.NewRecorder()
     // Act
     h.ServeHTTP(rr, req)
     // Assert
     require.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("should return http 500 in case of panic"))
}