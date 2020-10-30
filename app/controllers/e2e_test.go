package controllers_test

import (
	"bytes"
	"discord-clone-server/app"
	_ "discord-clone-server/migrations"
	"discord-clone-server/models"
	"discord-clone-server/seeder"
	"discord-clone-server/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pressly/goose"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	server          *httptest.Server
	Services        app.Services
	dbName          string
	DB              *gorm.DB
}

func TestE2ETestSuite(t *testing.T) {
	fmt.Print("\n\n\n\nTestE2ETestSuite\n")
	suite.Run(t, &e2eTestSuite{})
}
func (s *e2eTestSuite) SetupSuite() {
	var services app.Services

	s.dbName = utils.GetRandomString(s.T().Name(), 6)
	services.DB = utils.InitTestDB(s.T(), s.dbName)
	s.DB = services.DB
	services.UserRepo = app.InitUserRepo(services.DB)
	services.ServerRepo = app.InitServerRepo(services.DB)
	services.RoleRepo = app.InitRoleRepo(services.DB)
	services.PermissionRepo = app.InitPermissionRepo(services.DB)
	services.Redis = utils.NewTestRedis()
	r, err := app.InitRouter(services)
	s.Require().NoError(err)
	ts := httptest.NewServer(r)
	s.Services = services
	s.server = ts

}

func (s *e2eTestSuite) TearDownSuite() {
	utils.DropTestDB(s.T(), s.DB, s.dbName)
	db, err := s.DB.DB()
	s.NoError(err)
	db.Close()

	s.server.Close()
}

func (s *e2eTestSuite) SetupTest() {
	sqlDB, err := s.DB.DB()
	s.NoError(err)
	goose.Run("up", sqlDB, ".", "")
	seeder.PermissionsSeeder(s.DB)
}

func (s *e2eTestSuite) TearDownTest() {
	sqlDB, err := s.DB.DB()
	s.NoError(err)
	goose.Run("reset", sqlDB, ".", "")
}

func GetLoginCookie(s *e2eTestSuite, user models.User) *http.Cookie {
	requestBody, err := json.Marshal(map[string]string{
		"email":    user.Email,
		"password": "password",
	})
	s.NoError(err)

	resp, err := http.Post(fmt.Sprintf("%s/login", s.server.URL), "application/json", bytes.NewBuffer(requestBody))
	assert.Equal(s.T(), 200, resp.StatusCode)

	for _, cookie := range resp.Cookies() {
		assert.Equal(s.T(), "discord_clone_session", cookie.Name)
		if cookie.Name == "discord_clone_session" {
			return &http.Cookie{Name: cookie.Name, Value: cookie.Value}
		}
	}
	return &http.Cookie{}
}

func CreateTestServer(s *e2eTestSuite, name string, private bool, userID uint) models.Server {
	adminRole, err := s.Services.RoleRepo.CreateAdminRole()
	if err != nil {
		s.NoError(err)
	}

	baseRole, err := s.Services.RoleRepo.CreateBaseRole()
	if err != nil {
		s.NoError(err)
	}
	server := models.Server{
		Name:    name,
		Private: private,
		User_ID: userID,
		Roles: []models.Role{
			adminRole,
			baseRole,
		},
	}
	s.DB.Create(&server)
	return server

}