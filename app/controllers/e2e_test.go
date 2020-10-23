package controllers_test

import (
	"discord-clone-server/app"
	_ "discord-clone-server/migrations"
	"discord-clone-server/utils"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/pressly/goose"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type e2eTestSuite struct {
	suite.Suite
	dbConnectionStr string
	port            int
	server          *httptest.Server
	dbName          string
	DB              *gorm.DB
	// dbMigration     *migrate.Migrate
}

func TestE2ETestSuite(t *testing.T) {
	fmt.Print("\n\n\n\nTestE2ETestSuite\n")
	suite.Run(t, &e2eTestSuite{})
}
func (s *e2eTestSuite) SetupSuite() {
	var services app.Services
	// var err error

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
	s.server = ts
}

func (s *e2eTestSuite) TearDownSuite() {
	fmt.Print("\n\n\nTestE2ETestSuite: teardown\n")
	utils.DropTestDB(s.T(), s.DB, s.dbName)
	s.server.Close()
	// p, _ := os.FindProcess(syscall.Getpid())
	// p.Signal(syscall.SIGINT)
}

func (s *e2eTestSuite) SetupTest() {
	sqlDB, err := s.DB.DB()
	s.NoError(err)
	goose.Run("up", sqlDB, ".", "")
	// if err := s.dbMigration.Up(); err != nil && err != migrate.ErrNoChange {
	// 	s.Require().NoError(err)
	// }
}

func (s *e2eTestSuite) TearDownTest() {
	fmt.Print("TestE2ETestSuite: teardown test\n")
	sqlDB, err := s.DB.DB()
	s.NoError(err)
	goose.Run("reset", sqlDB, ".", "")
	// s.NoError(s.dbMigration.Down())
}
