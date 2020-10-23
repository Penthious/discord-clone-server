package repositories

import (
	"discord-clone-server/models"
	"discord-clone-server/utils"
	"testing"
)

func Test_UserRepo_Get(t *testing.T) {
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

/*
... repos/sets_test.go
func TestGlobalSets(t *testing.T) {
	dbName := utils.TestDBName("negatives")
	gormDB := utils.InitTestDB(t, dbName)

	.... some test code
}

... utils/test_utils_test.go

import	mySQL "github.com/go-sql-driver/mysql"


func GetMysqlDSN(prefix string) string {
	config :=  &mySQL.Config{
		User:                 os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_USER")),
		Passwd:               os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_PASSWORD")),
		Addr:                 os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_ADDR")),
		DBName:               os.Getenv(fmt.Sprintf("%s%s", prefix, "DB_NAME")),
		Net:                  "tcp",
		Timeout:              5 * time.Second,
		ReadTimeout:          5 * time.Second,
		WriteTimeout:         5 * time.Second,
		MultiStatements:      true,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	return config.FormatDSN()
}

func TestDBName(prefix string) string {
	length := 6
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	runes := make([]rune, length)
	for i := 0; i < length; i++ {
		runes[i] = rune(charset[rand.Intn(len(charset))])
	}

	return fmt.Sprintf("%s_%s", prefix, string(runes))
}

func TestGlobalSets(t *testing.T) {
	dbName := utils.TestDBName("negatives")
	gormDB := utils.InitTestDB(t, dbName)
	defer utils.DropTestDB(t, gormDB, dbName)

	setsRepo := NewSetRepo(gormDB, "sets", SetsConfig{})

	sets := []models.Set{
		{Name: "Test Set 1", Description: "sample description", Channel: "SEO", IsGlobal: true, UserID: 1},
		{Name: "Test Set 2", Description: "sample description", Channel: "SEO", IsGlobal: true, UserID: 1},
		{Name: "Not Global Test Set", Description: "sample description", Channel: "SEO", IsGlobal: false, UserID: 1},
	}
	MakeTestSets(t, gormDB, sets)

	keywords := []models.Keyword{
		{Keyword: "copied keyword1", SetID: sets[0].ID},
		{Keyword: "copied keyword2", SetID: sets[0].ID},
		{Keyword: "copied keyword3", SetID: sets[0].ID},
	}
	MakeTestKeywords(t, gormDB, keywords)

	negatives := []models.Negative{
		{String: "sample negative", SetID: sets[0].ID},
		{String: "sample negative 2", SetID: sets[0].ID},
	}
	MakeTestNegatives(t, gormDB, negatives)

	buckets := []models.Bucket{
		{Name: "Sample Bucket", SetID: sets[0].ID},
	}
	MakeTestBuckets(t, gormDB, buckets)

	bucketValues := []models.BucketValue{
		{SearchFor: "sample search for", GroupAs: "sample group as", BucketID: buckets[0].ID},
	}
	MakeTestBucketValues(t, gormDB, bucketValues)

	t.Run("GetGlobalSets", func(t *testing.T) {
		globalSets, err := setsRepo.GlobalSets()
		require.NoError(t, err)
		assert.Equal(t, len(globalSets), 2, "global sets length retrieved should be 2")
	})

*/
