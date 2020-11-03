package repositories

import (
	"discord-clone-server/models"
	"discord-clone-server/utils"
	"testing"
)

func Test_ServerRepo_Append(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)

	server := models.Server{Name: "New server"}
	db.Create(&server)

	user := models.User{FirstName: "Bob1", LastName: "bobbers", Username: "bobbers1", Password: "test", Email: "bob1@bob.com"}
	db.Create(&user)

	serverRepo := NewServerRepo(db)

	if err := serverRepo.Append(user, &server); err != nil {
		t.Fatalf("Error appending user to server: %v", err.Error())
	}

	var results models.User
	tx := db.Raw("SELECT * FROM `users` JOIN `server_users` ON `server_users`.`user_id` = `users`.`id` AND `server_users`.`server_id` = ? WHERE `users`.`id` = ? AND `users`.`deleted_at` IS NULL", server.ID, user.ID).Scan(&results)
	if tx.Error != nil {
		t.Fatalf("Error getting user from DB: %v", tx.Error)
	}

	if results.ID != user.ID {
		t.Fatalf("User was not associated to server")
	}

}

func Test_ServerRepo_Create(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	serverRepo := NewServerRepo(db)

	server := models.Server{Name: "New server"}
	err := serverRepo.Create(&server)
	if err != nil {
		t.Fatalf("Error creating server: %v", err.Error())
	}

	var results models.User
	tx := db.Raw("SELECT * FROM servers where servers.id = ?", server.ID).Scan(&results)
	if err := tx.Error; err != nil {
		t.Fatalf("Error getting user: %v", err.Error())
	}

	if results.ID == 0 {
		t.Fatalf("Creating server was not successfull")
	}

}

func Test_ServerRepo_Find(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	serverRepo := NewServerRepo(db)

	server := models.Server{Name: "New server"}
	err := serverRepo.Create(&server)
	if err != nil {
		t.Fatalf("Error creating server: %v", err.Error())
	}

	var results models.Server

	tx := db.Raw("SELECT * FROM servers WHERE servers.id = ?", server.ID).Scan(&results)

	if tx.Error != nil {
		t.Fatalf("Failed quering servers: %v", tx.Error)
	}

	if server.ID != results.ID {
		t.Fatalf("Failed to find server")
	}
}

func Test_ServerRepo_FindWithRoles(t *testing.T) {
	t.Parallel()
	dbName := utils.GetRandomString(t.Name(), 6)
	db := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, db, dbName)
	serverRepo := NewServerRepo(db)
	roleRepo := NewRoleRepo(db)
	adminRole, _ := roleRepo.CreateAdminRole()
	baseRole, _ := roleRepo.CreateBaseRole()

	server := models.Server{Name: "New server", Roles: []models.Role{adminRole, baseRole}}
	tx := db.Create(&server)
	if tx.Error != nil {
		t.Fatalf("Error creating server: %v", tx.Error)
	}

	var serverWithRoles models.Server
	err := serverRepo.FindWithRoles(server.ID, &serverWithRoles)
	if err != nil {
		t.Fatalf("Error fetching server with roles: %v", err.Error())
	}

	var results models.Server
	tx = db.Preload("Roles").Find(&results, server.ID)

	if tx.Error != nil {
		t.Fatalf("Failed quering servers: %v", tx.Error)
	}

	if server.ID != results.ID {
		t.Fatalf("Failed to find server")
	}
	if len(results.Roles) != 2 {
		t.Fatalf("Failed to get the correct amount of roles back")
	}

}
