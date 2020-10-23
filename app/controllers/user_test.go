package controllers_test

import (
	"bytes"
	"discord-clone-server/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) Test_EndToEnd_UserCreate() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "Bob",
		"last_name":  "Bobbers",
		"username":   "Bobbers123",
		"email":      "bob@bobbers.com",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 201, resp.StatusCode)
	type httpResp struct {
		User models.User `json:"user"`
	}
	var r httpResp

	s.NoError(err)

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	for _, cookie := range resp.Cookies() {
		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
	}
	assert.Equal(s.T(), uint(1), r.User.ID)
	assert.Equal(s.T(), "Bob", r.User.FirstName)
	assert.Equal(s.T(), "Bobbers", r.User.LastName)
	assert.Equal(s.T(), "Bobbers123", r.User.Username)
	assert.Equal(s.T(), "bob@bobbers.com", r.User.Email)

	// @todo: find out how to check if gin.Context has current user in session
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__NoFirstNameError() {
	requestBody, err := json.Marshal(map[string]string{
		"last_name": "Bobbers",
		"username":  "Bobbers123",
		"email":     "bob@bobbers.com",
		"password":  "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.FirstName' Error:Field validation for 'FirstName' failed on the 'required' tag", r.Error)
}
func (s *e2eTestSuite) Test_EndToEnd_UserCreate__NoLastNameError() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"username":   "Bobbers123",
		"email":      "bob@bobbers.com",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.LastName' Error:Field validation for 'LastName' failed on the 'required' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__NoUsernameError() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"email":      "bob@bobbers.com",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Username' Error:Field validation for 'Username' failed on the 'required' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__UsernameMustBeLongerThan3Error() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"username":   "bo", // 2
		"email":      "bob@bobbers.com",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Username' Error:Field validation for 'Username' failed on the 'min' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__UsernameMustBeShorterThan15Error() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"username":   "bobss name is 16", // 16
		"email":      "bob@bobbers.com",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Username' Error:Field validation for 'Username' failed on the 'max' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__NoEmailError() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "Bob",
		"last_name":  "Bobbers",
		"username":   "Bobbers123",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Email' Error:Field validation for 'Email' failed on the 'required' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__EmailMustBeValidError() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "Bob",
		"last_name":  "Bobbers",
		"username":   "Bobbers123",
		"email":      "notValidEmail",
		"password":   "testing123",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Email' Error:Field validation for 'Email' failed on the 'email' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__NoPasswordError() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"username":   "Bobbers123",
		"email":      "bob@bobbers.com",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Password' Error:Field validation for 'Password' failed on the 'required' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__PasswordMustBeLongerThan8Error() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"username":   "bobs",
		"email":      "bob@bobbers.com",
		"password":   "1234567", // 7
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Password' Error:Field validation for 'Password' failed on the 'min' tag", r.Error)
}

func (s *e2eTestSuite) Test_EndToEnd_UserCreate__PasswordMustBeShorterThan36Error() {
	requestBody, err := json.Marshal(map[string]string{
		"first_name": "bob",
		"last_name":  "Bobbers",
		"username":   "bobss",
		"email":      "bob@bobbers.com",
		"password":   "1234567891234567891234567891234567891", // 37
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/users", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 400, resp.StatusCode)
	type httpResp struct {
		Error string `json:"error"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.Equal(s.T(), "Key: 'userCreateParams.Password' Error:Field validation for 'Password' failed on the 'max' tag", r.Error)
}
