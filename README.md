# connectfour-backend
Back-end  application that can serve a complete connect-four match of either player-versus-player or player-versus-computer. The focus is on a easy-to-use REST API that minimizes the logic that is to be implemented in a front-end that uses the API.

## **Start server**

```console
daniel@r2d2: ~/projects/connect-four$ go run .
```

Starts a server at localhost:8080. The server stores the running matches in-memory so that all progress will be lost when restarting. The application uses cookies to re-indentify clients. The cookie contains a uuid that is used to save/retrieve the game board that describes the current state of the running match to/from the in-memory database. To start a match use a client that can manage cookies like cURL or postman.

## **API**

The all game boards hold by the backend have a fixed size of 6 rows and 7 columns. The first turn is always done by the red player.

### **Board**
----
  Used to retrieve the current game board. If no cookie is attached to the request a new game board will be returned and a new cookie (that contains the game id of the new board) will be added to the response. The game board can be requested as either text or json representation. Text: n=empty field, r=red chip, b=blue chip  JSON: 0=empty field, 1=blue, 2=red (or however you want to define it). 

**Request:**
```json
 GET /board HTTP/1.1
 Host: localhost:8080
 Cookie: gameID=ec4d28b0-ed4f-4498-a8de-dd14649e312c
 Content-Type: text/plain
```
**Response:**
```json
HTTP/1.1 200 OK
Content-Length: 90
Content-Type: text/plain; charset=utf-8

n n n n n n n
n n n n n n n
n n n n n n n
n n n n n n n
n n n b n n n
n n n r n n n
```

**Sample Call:**
 
 Gets a new game board and saves the new cookie to temp/cookies (file must be created before).
 
  ```console
  curl -c temp/cookies localhost:8080/board -H "Content-Type: text/plain"
  ```

 Gets an existing game board using an existing cookie from temp/cookies.

  ```console
  curl -b temp/cookies localhost:8080/board -H "Content-Type: text/plain"
  ```

**Request:**
```json
 GET /board HTTP/1.1
 Host: localhost:8080
 Cookie: gameID=a8b83d5e-39f7-47ec-acb9-5105b2a5c890
 Content-Type: application/json

```
**Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 123

{"Fields":[[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,1,0,0,0],[0,0,0,2,0,0,0]], "NextColor":2}

```

**Sample Call:**

  ```console
  curl -b temp/cookies localhost:8080/board -H "Content-Type: application/json"
  ```

### **Turn**
----
 Used to perform the next turn on the board. The request body must include the index (0-6) of the column the next chip is to be inserted at (see below). Use GET /board with content-type application/json as detailed above to retrieve the status of the current board and the next color. If no cookie is attached to the request or the gameID does not match a game board a http 400 is returned. Use GET /board to get a new board/cookie in that case.

**Request:**
```json
PATCH /turn HTTP/1.1
Host: localhost:8080
Cookie: gameID=a8b83d5e-39f7-47ec-acb9-5105b2a5c890
Content-Type: application/json
{"column": 3}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Length: 0
```

**Failed Response:**
```json
HTTP/1.1 404 Not Found
Content-Length: 62
Content-Type: text/plain; charset=utf-8

no board created, please perform a GET request on /board first
``` 

**Failed Response:**
```json
HTTP/1.1 409 Conflict
Content-Length: 16
Content-Type: text/plain; charset=utf-8

column 3 is full
``` 

**Failed Response:**
```json
HTTP/1.1 409 Conflict
Content-Length: 26
Content-Type: text/plain; charset=utf-8

match has already a winner
``` 

**Failed Response:**
```json
HTTP/1.1 400 Bad Request
Content-Length: 30
Content-Type: text/plain; charset=utf-8

column 8 is out of bounds: 0-6
``` 

**Sample Call:**

  ```console
  curl -b temp/cookies -X PATCH localhost:8080/turn -H "--Content-Type: application/json" --data-raw '{"column": 3}'
  ```

### **AI**
----
 Used to perform the next turn on the board using artificial intelligence. At the moment a monte carlo algorithm with a fixed number of repetitions = 500 is the only avaible choice. Returns the column where the next chip was inserted at. A subsequent GET /board request (see above) will return the updated game board. If no cookie is attached to the request or the gameID does not match a game board a http 400 is returned. Use GET /board to get a new board/cookie in that case.

**Request:**
```json
PATCH /ai/montecarlo HTTP/1.1
Host: localhost:8080
Cookie: gameID=270c91e2-7aa9-4bc8-b72d-9d40543b203d
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: 13

{"Column":3}
```

**Failed Response:**
```json
HTTP/1.1 404 Not Found
Content-Length: 62
Content-Type: text/plain; charset=utf-8

no board created, please perform a GET request on /board first
``` 

**Failed Response:**
```json
HTTP/1.1 409 Conflict
Content-Length: 26
Content-Type: text/plain; charset=utf-8

match has already a winner
``` 

**Sample Call:**

  ```console
  curl -b temp/cookies -X PATCH localhost:8080/ai/montecarlo
  ```


