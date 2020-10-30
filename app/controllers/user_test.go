package controllers_test

import (
	"bytes"
	"discord-clone-server/models"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) Test_EndToEnd_UserCreate() {
	s.T().Run("User Create", func(t *testing.T) { UserCreate(s) })
	s.T().Run("User Create - no first name - error", func(t *testing.T) { UserCreate__NoFirstNameError(s) })
	s.T().Run("User Create - no last name - error", func(t *testing.T) { UserCreate__NoLastNameError(s) })
	s.T().Run("User Create - no user name - error", func(t *testing.T) { UserCreate__NoUsernameError(s) })
	s.T().Run("User Create - username must be longer than 3 - error", func(t *testing.T) { UserCreate__UsernameMustBeLongerThan3Error(s) })
	s.T().Run("User Create - username must be shorter than 15 - error", func(t *testing.T) { UserCreate__UsernameMustBeShorterThan15Error(s) })
	s.T().Run("User Create - no email - error", func(t *testing.T) { UserCreate__NoEmailError(s) })
	s.T().Run("User Create - email must be valid - error", func(t *testing.T) { UserCreate__EmailMustBeValidError(s) })
	s.T().Run("User Create - no password - error", func(t *testing.T) { UserCreate__NoPasswordError(s) })
	s.T().Run("User Create - password must be longer than 8 - error", func(t *testing.T) { UserCreate__PasswordMustBeLongerThan8Error(s) })
	s.T().Run("User Create - password must be shorter than 36 - error", func(t *testing.T) { UserCreate__PasswordMustBeShorterThan36Error(s) })

}

func (s *e2eTestSuite) Test_EndToEnd_UserLogin() {
	user := models.User{
		FirstName: "Bobs",
		LastName:  "Bobbers",
		Username:  "bobbies",
		Email:     "bob@bobers.com",
		Password:  "password",
	}
	s.DB.Create(&user)
	s.T().Run("Login user -- with email", func(t *testing.T) { Login__WithEmail(s) })
	s.T().Run("Login user -- with username", func(t *testing.T) { Login__WithUsername(s) })
}

func UserCreate(s *e2eTestSuite) {
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

}

func UserCreate__NoFirstNameError(s *e2eTestSuite) {
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
func UserCreate__NoLastNameError(s *e2eTestSuite) {
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

func UserCreate__NoUsernameError(s *e2eTestSuite) {
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

func UserCreate__UsernameMustBeLongerThan3Error(s *e2eTestSuite) {
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

func UserCreate__UsernameMustBeShorterThan15Error(s *e2eTestSuite) {
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

func UserCreate__NoEmailError(s *e2eTestSuite) {
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

func UserCreate__EmailMustBeValidError(s *e2eTestSuite) {
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

func UserCreate__NoPasswordError(s *e2eTestSuite) {
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

func UserCreate__PasswordMustBeLongerThan8Error(s *e2eTestSuite) {
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

func UserCreate__PasswordMustBeShorterThan36Error(s *e2eTestSuite) {
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

func Login__WithUsername(s *e2eTestSuite) {
	requestBody, err := json.Marshal(map[string]string{
		"username": "bobbies",
		"password": "password",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/login", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 200, resp.StatusCode)
	type httpResp struct {
		User models.User `json:"user"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))

	assert.Equal(s.T(), "bobbies", r.User.Username)
	assert.Equal(s.T(), "", r.User.Password)
}

func Login__WithEmail(s *e2eTestSuite) {
	requestBody, err := json.Marshal(map[string]string{
		"email":    "bob@bobers.com",
		"password": "password",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/login", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	s.NoError(err)
	assert.Equal(s.T(), 200, resp.StatusCode)
	type httpResp struct {
		User models.User `json:"user"`
	}
	var r httpResp

	s.NoError(json.NewDecoder(resp.Body).Decode(&r))

	assert.Equal(s.T(), "bobbies", r.User.Username)
	assert.Equal(s.T(), "", r.User.Password)

	for _, cookie := range resp.Cookies() {
		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
	}

}
