package controllers_test

import (
	"bytes"
	"discord-clone-server/app/services"
	"discord-clone-server/models"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func (s *e2eTestSuite) Test_EndToEnd_ServerCreate() {
	// setup
	user := models.User{
		FirstName: "Bobs",
		LastName:  "Bobbers",
		Username:  "bobbies",
		Email:     "bob@bobers.com",
		Password:  "password",
	}
	s.DB.Create(&user)

	// run test
	s.T().Run("Server Create", func(t *testing.T) { ServerCreate(s) })
	s.T().Run("Server Create - must be logged in - error", func(t *testing.T) { ServerCreate_MustBeLoggedIn(s) })
}

func (s *e2eTestSuite) Test_EndToEnd_InviteUser() {
	// setup
	mainUser := models.User{
		FirstName: "Bobs",
		LastName:  "Bobbers",
		Username:  "bobbies",
		Email:     "bob@bobers.com",
		Password:  "password",
	}
	invitedUser := models.User{
		FirstName: "dillian",
		LastName:  "jones",
		Username:  "invited",
		Email:     "someotheruser@bobers.com",
		Password:  "password",
	}
	s.DB.Create(&mainUser)
	s.DB.Create(&invitedUser)
	server := CreateTestServer(s, "Test Name", false, mainUser.ID)

	// Append mainUser to server
	s.NoError(s.Services.ServerRepo.Append(mainUser, &server))

	// assign mainUser to role admin
	_, err := s.Services.RoleRepo.AttachServerRoles([]models.ServerUserRole{
		{
			ServerID: server.ID,
			UserID:   mainUser.ID,
			RoleID:   server.Roles[0].ID,
		},
	})
	s.NoError(err)

	// run test
	s.T().Run("Invite User", func(t *testing.T) { InviteUser(s) })

}

func ServerCreate(s *e2eTestSuite) {
	var user models.User
	s.DB.Find(&user, 1)

	client := &http.Client{}

	// Prepare the http request for server
	requestBody, err := json.Marshal(map[string]string{
		"name": "New Server",
	})
	s.NoError(err)
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/servers/", s.server.URL), bytes.NewBuffer(requestBody))
	s.NoError(err)
	req.Header.Add("Content-Type", "application/json")

	// Set login cookie to current request
	req.AddCookie(GetLoginCookie(s, user))

	// Send off request
	resp, err := client.Do(req)
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
	assert.Equal(s.T(), false, dbServers[0].Private)

	// Assert that we created 2 roles
	var dbRoles []models.Role
	err = s.DB.Find(&dbRoles).Error
	s.NoError(err)
	assert.Equal(s.T(), 2, len(dbRoles))

	// Assert that the created roles are attached to the created server
	roles, err := s.Services.RoleRepo.GetServerRoles(&dbServers[0])
	s.NoError(err)
	assert.Equal(s.T(), 2, len(roles))
	assert.Equal(s.T(), "Admin", roles[0].Name)
	assert.Equal(s.T(), "Base", roles[1].Name)

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

func ServerCreate_MustBeLoggedIn(s *e2eTestSuite) {
	requestBody, err := json.Marshal(map[string]string{
		"name": "New Server",
	})
	s.NoError(err)
	resp, err := http.Post(fmt.Sprintf("%s/servers/", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	assert.Equal(s.T(), 401, resp.StatusCode)
}

func InviteUser(s *e2eTestSuite) {

	var user models.User
	var server models.Server
	s.DB.Find(&user).Where("username = ?", "invited")
	s.DB.Find(&server, 1)

	client := &http.Client{}

	// Prepare the http request for server invite
	requestBody, err := json.Marshal(map[string]uint{
		"server_id": server.ID,
		"user_id":   user.ID,
	})
	s.NoError(err)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/servers/invite", s.server.URL), bytes.NewBuffer(requestBody))
	s.NoError(err)
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(GetLoginCookie(s, user))

	// Send off request
	resp, err := client.Do(req)
	s.NoError(err)

	assert.Equal(s.T(), 200, resp.StatusCode)

	type httpResp struct {
		Message struct {
			Invite string
		}
	}
	var rsi services.RedisServerInvite

	var r httpResp
	// Assert that response is what we expect
	s.NoError(json.NewDecoder(resp.Body).Decode(&r))
	assert.NotEqual(s.T(), "", r.Message.Invite)
	err = services.GetRedisKey(r.Message.Invite, s.Services.Redis, &rsi)
	s.NoError(err)
	assert.Equal(s.T(), services.SERVER_INVITE, rsi.Key)
	assert.Equal(s.T(), user.ID, rsi.UserID)
	assert.Equal(s.T(), server.ID, rsi.ServerID)
}

