package board

import (
	"encoding/json"
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

func (mock *GameServiceMock) Board(gameID string) game.Board {
	args := mock.Called(gameID)
	return args.Get(0).(game.Board)
}

func (mock *GameServiceMock) TurnAI(gameID string, ai game.AI) (int, error) {
	fmt.Println("The number you have dialed is not available.")
	return -1, nil
}

func (mock *GameServiceMock) Turn(column int, gameID string) error {
	fmt.Println("The number you have dialed is not available.")
	return nil
}

func (mock *GameServiceMock) Winner(gameID string) (string, error) {
	fmt.Println("The number you have dialed is not available.")
	return "" , nil
}

func (mock *GameServiceMock) Reset(gameID string) (error) {
	fmt.Println("The number you have dialed is not available.")
	return nil
}



var h http.Handler

func setup() {
	gameServiceMock := GameServiceMock{}
	gameID := func(w http.ResponseWriter, req *http.Request) string {
		c, _ := req.Cookie("gameID")
		return c.Value
	}
	gameServiceMock.On("Board", "324234-555").Return(game.Board{})
	gameServiceMock.On("Board", "please panic!").Panic("panic!")
	h = NewHandler(&gameServiceMock, gameID)
}

func TestHandler(t *testing.T) {
	// Arrange
	setup()
	req, _ := http.NewRequest("", "", nil)
	req.Header.Set("Content-type", "application/json")
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "324234-555"})
	rr := httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	body, _ := ioutil.ReadAll(rr.Body)
	var board game.Board
	err := json.Unmarshal(body, &board)
	// Assert
	require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
    require.Equal(t, nil, err, fmt.Sprintf("should return fields of board as json if content type is application/json"))
    
    // Arrange
	req, _ = http.NewRequest("", "", nil)
	req.Header.Set("Content-type", "")
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "324234-555"})
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	body, _ = ioutil.ReadAll(rr.Body)
	fmt.Println()
	// Assert
	require.Equal(t, http.StatusOK, rr.Code, fmt.Sprintf("should return http 200 if request is valid"))
	require.Equal(t,
		"n n n n n n n \n" +
		"n n n n n n n \n" + 
		"n n n n n n n \n" + 
		"n n n n n n n \n" + 
		"n n n n n n n \n" +
		"n n n n n n n \n", string(body), fmt.Sprint("should return fields of board as text if content type is NOT application/json"))

	// Arrange
	req, _ = http.NewRequest("", "", nil)
	req.AddCookie(&http.Cookie{Name: "gameID", Value: "please panic!"})
	rr = httptest.NewRecorder()
	// Act
	h.ServeHTTP(rr, req)
	// Assert
	require.Equal(t, http.StatusInternalServerError, rr.Code, fmt.Sprintf("should return http 500 in case of panic"))
}
