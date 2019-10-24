package user

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	//given
	router := gin.Default()
	router.GET("/", Get())
	response := httptest.NewRecorder()

	//when
	req, _ := http.NewRequest("GET", "/?name=test&lastName=test1&age=5", nil)
	router.ServeHTTP(response, req)

	//then
	var user User
	err := json.NewDecoder(response.Body).Decode(&user)
	assert.Equal(t, err, nil, "error should be nil")
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "test", user.FirstName, "test first name expected")
	assert.Equal(t, "test1", user.LastName, "test last name expected")
	assert.Equal(t, 5, user.Age, "test age expected")
}
