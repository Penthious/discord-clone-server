package controllers_test

import (
	"bytes"
	"discord-clone-server/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) Test_EndToEnd_ServerCreate() {
	// Create User
	user := models.User{
		FirstName: "Bobs",
		LastName:  "Bobbers",
		Username:  "bobbies",
		Email:     "bob@bobers.com",
		Password:  "password",
	}
	s.DB.Create(&user)

	// Login user
	requestBody, err := json.Marshal(map[string]string{
		"email":    "bob@bobers.com",
		"password": "password",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/login", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	assert.Equal(s.T(), 200, resp.StatusCode)
	client := &http.Client{}

	// Prepare the http request for server
	requestBody, err = json.Marshal(map[string]string{
		"name": "New Server",
	})
	s.NoError(err)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/servers/", s.server.URL), bytes.NewBuffer(requestBody))
	s.NoError(err)
	req.Header.Add("Content-Type", "application/json")

	// Set login cookie to current request
	for _, cookie := range resp.Cookies() {
		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
		req.AddCookie(&http.Cookie{Name: cookie.Name, Value: cookie.Value})
	}

	// Send off request
	resp, err = client.Do(req)
	s.NoError(err)

	assert.Equal(s.T(), 201, resp.StatusCode)
	type httpResp struct {
		Server models.Server `json:"server"`
	}
	var r httpResp

	// Assert that we only have one server created
	var dbServers []models.Server
	err = s.DB.Find(&dbServers).Error
	s.NoError(err)
	assert.Equal(s.T(), 1, len(dbServers))

	// Assert that we created 2 roles
	var dbRoles []models.Role
	err = s.DB.Find(&dbRoles).Error
	s.NoError(err)
	assert.Equal(s.T(), 2, len(dbRoles))

	// Assert that the created roles are attached to the created server
	// @todo: add in a way to get all roles attached to a server

	// Assert that role admin was added to this user
	userRoles, err := s.Services.RoleRepo.GetUserServerRoles(user.ID, dbServers[0].ID)
	s.NoError(err)
	assert.Equal(s.T(), 1, len(userRoles))
	assert.Equal(s.T(), "Admin", userRoles[0].Name)

	// Assert that we still have the cookie on the resp
	for _, cookie := range resp.Cookies() {
		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
	}

	// Assert that response is what we expect
	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	// assert.Equal(s.T(), uint(1), r.User.ID)
	// assert.Equal(s.T(), "Bob", r.User.FirstName)
	// assert.Equal(s.T(), "Bobbers", r.User.LastName)
	// assert.Equal(s.T(), "Bobbers123", r.User.Username)
	// assert.Equal(s.T(), "bob@bobbers.com", r.User.Email)

	// @todo: find out how to check if gin.Context has current user in session
}
