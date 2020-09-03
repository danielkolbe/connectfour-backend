package ai

import (
	"bytes"
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

func (mock *GameServiceMock) TurnAI(gameID string, ai game.AI) (int, error) {
	args := mock.Called(gameID, ai)
	return args.Int(0), args.Error(1)
}

func (mock *GameServiceMock) Turn(column int, gameID string) error {
	fmt.Println("The number you have dialed is not available.")
	return nil
}

func (mock *GameServiceMock) Board(gameID string) game.Board {
	fmt.Println("The number you have dialed is not available.")
	return game.Board{}
}

func (mock *GameServiceMock) Winner(gameID string) (string, error) {
	fmt.Println("The number you have dialed is not available.")
	return "", nil
}

func (mock *GameServiceMock) Reset(gameID string) error {
	fmt.Println("The number you have dialed is not available.")
	return nil
}

type AIMock struct {
	mock.Mock
}

func (mock *AIMock) NextTurn(b *game.Board) (int, error) {
	fmt.Println("The number you have dialed is not available.")
	return -1, nil
}

var h http.Handler
var cookie *http.Cookie

func setup() {
	gameServiceMock := GameServiceMock{}
	aiMock := AIMock{}
	gameID := func(w http.ResponseWriter, req *http.Request) string {
		c, _ := req.Cookie("gameID")
		return c.Value
	}
	gameServiceMock.On("TurnAI", "324234-111", &aiMock).Return(3, nil)
	gameServiceMock.On("TurnAI", "324234-222", &aiMock).Return(-1, game.NewMatchIsOverError("match has already a winner"))
	gameServiceMock.On("TurnAI", "324234-333", &aiMock).Return(-1, fmt.Errorf("error"))
	gameServiceMock.On("TurnAI", "324234-444", &aiMock).Panic("panic!")
	gameServiceMock.On("TurnAI", "324234-555", &aiMock).Return(-1, game.NewBoardDoesNotExistError("324234-555"))
	h = NewHandler(&gameServiceMock, &aiMock, gameID)
}

func TestHandler(t *testing.T) {
	// Arrange
	setup()
	req, _ := http.NewRequest("", "", bytes.NewReader(nil))
	cookie = &http.Cookie{Name: "gameID", Value: "324234-111"}
	req.AddCookie(cookie)
	rr := httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
	body, _ := ioutil.ReadAll(rr.Body)
	require.Equal(t, "{\"Column\":3}\n", string(body), fmt.Sprintf("should return http 200 if request is valid"))

	// Arrange
	req, _ = http.NewRequest("", "", bytes.NewReader(nil))
	cookie = &http.Cookie{Name: "gameID", Value: "324234-222"}
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	bodyBytes, _ := ioutil.ReadAll(rr.Body)
	bodyString := string(bodyBytes)
	require.Equal(t, http.StatusConflict, rr.Code, fmt.Sprintf("should return http 409 if game service returns an MatchIsOverError"))
	require.Equal(t, "match has already a winner", bodyString, fmt.Sprintf("should add the correct error to response"))

	// Arrange
	cookie = &http.Cookie{Name: "gameID", Value: "324234-333"}
	req, _ = http.NewRequest("", "", bytes.NewReader(nil))
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
	cookie = &http.Cookie{Name: "gameID", Value: "324234-444"}
	req, _ = http.NewRequest("", "", bytes.NewReader(nil))
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
	cookie = &http.Cookie{Name: "gameID", Value: "324234-555"}
	req, _ = http.NewRequest("", "", bytes.NewReader(nil))
	req.AddCookie(cookie)
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	bodyBytes, _ = ioutil.ReadAll(rr.Body)
	bodyString = string(bodyBytes)
	require.Equal(t, http.StatusNotFound, rr.Code, fmt.Sprintf("should return http 404 if board does not exist"))
	require.Equal(t, "no board created, please perform a GET request on /board first", bodyString, fmt.Sprintf("should add the correct error to response"))
}
