package repositories

import (
	"discord-clone-server/models"
	"fmt"

	"gorm.io/gorm"
)

func NewUserRepo(db *gorm.DB) UserRepo {
	return userRepo{
		DB: db,
	}
}

type UserRepo interface {
	Create(*models.User) error
	Get() ([]models.User, error)
	Find(params UserFindParams) (models.User, error)
	AttachServerRoles([]models.ServerUserRole) (models.User, error)
}

type userRepo struct {
	DB *gorm.DB
}

func (r userRepo) Create(user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r userRepo) Get() ([]models.User, error) {
	var users []models.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

type UserFindParams struct {
	ID       uint
	Email    string
	Username string
}

func (r userRepo) Find(p UserFindParams) (models.User, error) {
	var user models.User
	if p.ID != 0 {
		tx := r.DB.First(&user, p.ID)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	} else if p.Email != "" {
		tx := r.DB.Where("email = ?", p.Email).First(&user)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	} else if p.Username != "" {
		tx := r.DB.Where("username = ?", p.Username).First(&user)
		if err := tx.Error; err != nil {
			return models.User{}, tx.Error
		}
	}

	return user, nil
}

func (r userRepo) AttachServerRoles(sur []models.ServerUserRole) (models.User, error) {

	var user models.User
	for _, s := range sur {
		fmt.Print(s)
		r.DB.Exec("INSERT INTO `server_user_roles` (`server_id`, `user_id`, `role_id`) VALUES (?, ?, ?)", s.ServerID, s.UserID, s.RoleID)

	}
	return user, nil
}

func (r userRepo) GetUserServerRoles(userID uint, serverID uint) (models.User, error) {
	query := `SELECT * FROM server_users as su
JOIN servers AS s
	ON su.server_id = s.id
JOIN server_user_roles AS sur
	ON s.id = sur.server_id
JOIN roles AS r
	ON sur.role_id = r.id
WHERE user_id = ?
AND server_id = ?
`
	r.DB.Exec(query, userID, serverID)
	// user -> server_users -> server -> server_user -> server_user_roles -> roles
	// user1 -> servers
	// server 1 -> roles
	// {user:
	// 	servers: [
	// 		{users: {
	// 			roles: []
	// 		}}
	// 	]
	// }
	// roles
	//  {user:
	// 	servers: [
	// 		{

	// 			roles: [],
	// 			channels: [
	// 				messages: [],

	// 			],
	// 		}
	// 	]
	// }
	var user models.User
	for _, s := range sur {
		fmt.Print(s.ServerID)
		fmt.Print(s.UserID)
		r.DB.Exec(fmt.Sprintf(""))

	}
	// 	`
	// SELECT
	//   s AS server
	//   ,u AS user
	//   ,r AS role
	//   ,r.permissions AS role_permissions -- { }
	// FROM
	//   servers s
	//   JOIN servers_users_su ON s.id = su.server_id
	//   JOIN users u ON su.user_id = u.id -- NEED THIS IF YOU WANT TO FILTER TO A USER NAME OR EMAIL
	//   JOIN server_user_roles sur ON -- JOIN ON USER AND SERVER IDS
	//     su.user_id = sur.user_id
	//     AND su.server_id = sur.server_id
	//   JOIN roles r ON
	//     sur.server_id = r.server_id
	//     AND sur.role_id = r.id
	// WHERE
	//   s.name = 'Server 1'
	//   AND u.name = 'Tomas'
	// ;
	// 	`
	return user, nil
}
