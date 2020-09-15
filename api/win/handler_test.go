package win

import (
	"fmt"
	"github.com/danielkolbe/connectfour/game"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type GameServiceMock struct {
	mock.Mock
}

func (mock *GameServiceMock) Winner(gameID string) (string, error) {
	args := mock.Called(gameID)
	return args.String(0), args.Error(1)
}

func (mock *GameServiceMock) TurnAI(gameID string, ai game.AI) (int, error) {
	fmt.Println("The number you have dialed is not available.")
	return -1, nil
}

func (mock *GameServiceMock) Turn(column int, gameID string) error {
	fmt.Println("The number you have dialed is not available.")
	return nil
}

func (mock *GameServiceMock) Board(gameID string) game.Board {
	fmt.Println("The number you have dialed is not available.")
	return game.Board{}
}

func (mock *GameServiceMock) Reset(gameID string) error {
	fmt.Println("The number you have dialed is not available.")
	return nil
}

var h http.Handler
var cookie *http.Cookie

func setup() {
	gameServiceMock := GameServiceMock{}
	gameID := func(w http.ResponseWriter, req *http.Request) string {
		c, _ := req.Cookie("gameID")
		return c.Value
	}
	gameServiceMock.On("Winner", "324234-555").Return("b", nil)
	gameServiceMock.On("Winner", "unknown").Return("b", game.NewBoardDoesNotExistError("unknown"))
	gameServiceMock.On("Winner", "panic").Panic("panic!")
	h = NewHandler(&gameServiceMock, gameID)
}

func TestHandler(t *testing.T) {
	// Arrange
	setup()
	req, _ := http.NewRequest("", "", nil)
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "324234-555"})
	rr := httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
	require.Equal(t, "b", bodyString, fmt.Sprintf("should return b if winner is blue"))

	// Arrange
	setup()
	req, _ = http.NewRequest("", "", nil)
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "unknown"})
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	bodyBytes, _ = ioutil.ReadAll(rr.Body)
	bodyString = string(bodyBytes)
	require.Equal(t, http.StatusNotFound, rr.Code, fmt.Sprintf("should return http 404 if board does not exist"))
	require.Equal(t, "no board created, please perform a GET request on /board first", bodyString, fmt.Sprintf("should add the correct error to response"))

	// Arrange
	setup()
	req, _ = http.NewRequest("", "", nil)
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "panic"})
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	bodyBytes, _ = ioutil.ReadAll(rr.Body)
	bodyString = string(bodyBytes)
	require.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("should return http 500 in case of panic"))
}
