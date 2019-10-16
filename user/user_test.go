package user

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	//given
	request, _ := http.NewRequest("GET", "/?name=test&lastName=test1&age=5", nil)
	response := httptest.NewRecorder()
	//when
	Get().ServeHTTP(response, request)

	//then
	var user User
	json.NewDecoder(response.Body).Decode(&user)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	assert.Equal(t, "test", user.FirstName, "test first name expected")
	assert.Equal(t, "test1", user.LastName, "test last name expected")
	assert.Equal(t, 5, user.Age, "test age expected")
}
