package repositories

import (
	"discord-clone-server/models"
	"discord-clone-server/utils"
	"testing"
)

func Test_UserRepo_Get(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	users := []models.User{
		{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"},
		{FirstName: "Bob2", LastName: "bobbers", Username: "bobbers2", Password: "test", Email: "bob2@bob.com"},
		{FirstName: "Bob3", LastName: "bobbers", Username: "bobbers3", Password: "test", Email: "bob3@bob.com"},
	}
	userRepo := NewUserRepo(db)

	utils.MakeTestUsers(t, db, users)

	dbUsers, err := userRepo.Get()
	if err != nil {
		t.Fatalf("Error getting users: %v", err.Error())
	}

	if len(dbUsers) != len(users) {
		t.Fatalf("DBusers does not match provided users: %v != %v", len(dbUsers), len(users))
	}
}

func Test_UserRepo_Create(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	user := models.User{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"}
	userRepo := NewUserRepo(db)

	err := userRepo.Create(&user)
	if err != nil {
		t.Fatalf("Error creating user: %v", err.Error())
	}
	var results models.User
	tx := db.Raw("SELECT * FROM users where users.id = ?", user.ID).Scan(&results)
	if err := tx.Error; err != nil {
		t.Fatalf("Error getting user: %v", err.Error())
	}

	if results.ID == 0 {
		t.Fatalf("Creating user was not successfull")
	}

}

func Test_UserRepo_Find_byID(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	user := models.User{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"}
	db.Create(&user)
	userRepo := NewUserRepo(db)

	dbUser, err := userRepo.Find(UserFindParams{ID: user.ID})
	if err != nil {
		t.Fatalf("Error executing Find: %v", err.Error())
	}

	if user.ID != dbUser.ID {
		t.Fatalf("Failed to find user")
	}
}

func Test_UserRepo_Find_byUsername(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	user := models.User{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"}
	db.Create(&user)
	userRepo := NewUserRepo(db)

	dbUser, err := userRepo.Find(UserFindParams{Username: user.Username})
	if err != nil {
		t.Fatalf("Error executing Find: %v", err.Error())
	}

	if user.ID != dbUser.ID {
		t.Fatalf("Failed to find user")
	}

}

func Test_UserRepo_Find_byEmail(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	user := models.User{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"}
	db.Create(&user)
	userRepo := NewUserRepo(db)

	dbUser, err := userRepo.Find(UserFindParams{Email: user.Email})
	if err != nil {
		t.Fatalf("Error executing Find: %v", err.Error())
	}

	if user.ID != dbUser.ID {
		t.Fatalf("Failed to find user")
	}
}
