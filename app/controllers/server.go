package controllers

import (
	"context"
	"discord-clone-server/app/services"
	"discord-clone-server/models"
	"discord-clone-server/repositories"
	"discord-clone-server/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type serverCreateParams struct {
	Name    string `json:"name" binding:"required,min=3,max=25"`
	Private bool   `json:"private"`
}

func ServerCreate(rs repositories.ServerRepo, rp repositories.PermissionRepo, ru repositories.UserRepo, rr repositories.RoleRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := services.GetUserContext(c)
		if err != nil {
			services.RespondBadRequestError(c, err, "Error getting user from context")
			return
		}
		var p serverCreateParams
		if err := c.ShouldBind(&p); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		permission, err := rp.Find(repositories.PermissionFindParams{Permission: "admin"})
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		server := models.Server{
			Name:    p.Name,
			Private: p.Private,
			User_ID: user.ID,
			Roles: []models.Role{
				{
					Name: "Admin",
					// Permissions: []*models.Permission{{Name: "Admin", Permission: "admin", Detail: "Full access"}},
					Permissions: []*models.Permission{&permission},
				},
			},
		}

		if err := rs.Create(&server); err != nil {
			services.RespondBadRequestError(c, err, "Error creating server")
			return
		}

		if err := rs.Append(&user, server); err != nil {
			services.RespondBadRequestError(c, err, "Error setting user to server")
			return
		}
		// attach server role admin to current user
		rr.AttachServerRoles([]models.ServerUserRole{
			{
				ServerID: server.ID,
				UserID:   user.ID,
				RoleID:   server.Roles[0].ID,
			},
		})

		services.RespondStatusCreated(c, "server", server)
		return
	}
}

type InviteUserParams struct {
	ServerID  int `json:"server_id" binding:"required_without=ChannelID"`
	ChannelID int `json:"channel_id" binding:"required_without=ServerID"`
	UserID    int `json:"user_id" binding:"required"`
}

func InviteUser(rs repositories.ServerRepo, r *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		// user, err := GetUserContext(c)
		// if err != nil {
		// 	RespondBadRequestError(c, err, "Error getting user from context")
		// 	return
		// }
		var p InviteUserParams
		if err := c.ShouldBind(&p); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		var server models.Server
		if p.ServerID != 0 {
			// find server
			rs.Find(p.ServerID, &server)
		} else {
			// find channel
		}
		// Get random key

		serverKey := utils.GetRandomString("invite_", 12)
		if err := r.Set(context.TODO(), "server_invite", serverKey, 10000).Err(); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		val2, err := r.Get(context.TODO(), "key").Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2", val2) // /servers/id/invite
		}
		// check current user exists on server

		// Check current user perms
		// Invite to server
		// create link to invite user

	}
}
