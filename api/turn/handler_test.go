package turn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func (mock *GameServiceMock) Board(gameID string) game.Board {
    fmt.Println("The number you have dialed is not available.")
    return game.Board{}
}

func (mock *GameServiceMock) Winner(gameID string) (string, error) {
	fmt.Println("The number you have dialed is not available.")
	return "" , nil
}


func (mock *GameServiceMock) Reset(gameID string) error {
    fmt.Println("The number you have dialed is not available.")
    return nil
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
    gameServiceMock.On("Turn", 10,"324234-555").Return(game.NewColumnIsOutOfBoundsError(10));
    gameServiceMock.On("Turn", 0,"324234-555").Return(game.NewBoardDoesNotExistError("324234-555"));
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
    bodyBytes, _ := ioutil.ReadAll(rr.Body)
    bodyString := string(bodyBytes)
    require.Equal(t, http.StatusConflict, rr.Code, fmt.Sprintf("should return http 409 if game service returns an ColumnIsFullError"))
    require.Equal(t, "column 5 is full", bodyString, fmt.Sprintf("should add the correct error to response"))
   
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
    bodyBytes, _ = ioutil.ReadAll(rr.Body)
    bodyString = string(bodyBytes)
    require.Equal(t, http.StatusConflict, rr.Code, fmt.Sprintf("should return http 409 if ame service returns an MatchIsOverError"))
    require.Equal(t, "match has already a winner", bodyString, fmt.Sprintf("should add the correct error to response"))
    
    // Arrange
    body.Column = -1
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    bodyBytes, _ = ioutil.ReadAll(rr.Body)
    bodyString = string(bodyBytes)
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if column number is negative"))
    require.Equal(t, "missing or negative column property in post body", bodyString, fmt.Sprintf("should add the correct error to response"))

    // Arrange
    wrongBody := struct {Unknown int}{4}
    bytesBody,_ = json.Marshal(wrongBody)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    bodyBytes, _ = ioutil.ReadAll(rr.Body)
    bodyString = string(bodyBytes)
    require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if body does not contain column field"))
    require.Equal(t, "missing or negative column property in post body", bodyString, fmt.Sprintf("should add the correct error to response"))

    // Arrange
    body.Column = 3
    bytesBody,_ = json.Marshal(body)
    req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
    req.AddCookie(cookie)
    rr = httptest.NewRecorder()
    // Act
    h.ServeHTTP(rr, req)
    // Assert
    bodyBytes, _ = ioutil.ReadAll(rr.Body)
    bodyString = string(bodyBytes)
    require.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("should return http 500 if game service returns an unknown error"))
    require.Equal(t, "sorry for that", bodyString, fmt.Sprintf("should add the correct error to response"))

     // Arrange
     body.Column = 10
     bytesBody,_ = json.Marshal(body)
     req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
     req.AddCookie(cookie)
     rr = httptest.NewRecorder()
     // Act
     h.ServeHTTP(rr, req)
     // Assert
     bodyBytes, _ = ioutil.ReadAll(rr.Body)
     bodyString = string(bodyBytes)
     require.Equal(t, http.StatusBadRequest, rr.Code, fmt.Sprintf("should return http 400 if game service returns an ColumnIsOutOfBoundsError"))
     require.Equal(t, "column 10 is out of bounds: 0-6", bodyString, fmt.Sprintf("should add the correct error to response"))

     // Arrange
     body.Column = 2
     bytesBody,_ = json.Marshal(body)
     req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
     req.AddCookie(cookie)
     rr = httptest.NewRecorder()
     // Act
     h.ServeHTTP(rr, req)
     // Assert
     bodyBytes, _ = ioutil.ReadAll(rr.Body)
     bodyString = string(bodyBytes)
     require.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("should return http 500 in case of panic"))
     require.Equal(t, "sorry for that\n", bodyString, fmt.Sprintf("should add the correct error to response"))

     // Arrange
     body.Column = 0
     bytesBody,_ = json.Marshal(body)
     req, _ = http.NewRequest("", "", bytes.NewReader(bytesBody))
     req.AddCookie(cookie)
     rr = httptest.NewRecorder()
     // Act
     h.ServeHTTP(rr, req)
     // Assert
     bodyBytes, _ = ioutil.ReadAll(rr.Body)
     bodyString = string(bodyBytes)
     require.Equal(t, http.StatusNotFound, rr.Code, fmt.Sprintf("should return http 400 if board does not exist"))
     require.Equal(t, "no board created, please perform a GET request on /board first", bodyString, fmt.Sprintf("should add the correct error to response"))
}