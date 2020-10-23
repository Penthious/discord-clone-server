package controllers_test

// import (
// 	"bytes"
// 	"discord-clone-server/models"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"

// 	"github.com/stretchr/testify/assert"
// )

// func (s *e2eTestSuite) Test_EndToEnd_User1Create() {
// 	requestBody, err := json.Marshal(map[string]string{
// 		"first_name": "Bob",
// 		"last_name":  "Bobbers",
// 		"username":   "Bobbers123",
// 		"email":      "bob@bobbers.com",
// 		"password":   "testing123",
// 	})
// 	s.NoError(err)
// 	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
// 	s.NoError(err)
// 	assert.Equal(s.T(), 201, resp.StatusCode)
// 	type httpResp struct {
// 		User models.User `json:"user"`
// 	}
// 	var r httpResp

// 	s.NoError(err)

// 	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
// 	for _, cookie := range resp.Cookies() {
// 		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
// 	}
// 	assert.Equal(s.T(), uint(1), r.User.ID)
// 	assert.Equal(s.T(), "Bob", r.User.FirstName)
// 	assert.Equal(s.T(), "Bobbers", r.User.LastName)
// 	assert.Equal(s.T(), "Bobbers123", r.User.Username)
// 	assert.Equal(s.T(), "bob@bobbers.com", r.User.Email)

// 	// @todo: find out how to check if gin.Context has current user in session
// }
