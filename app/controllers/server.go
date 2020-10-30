package controllers

import (
	"discord-clone-server/app/services"
	"discord-clone-server/models"
	"discord-clone-server/repositories"
	"discord-clone-server/utils"
	"errors"
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

		adminRole, err := rr.CreateAdminRole()
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		baseRole, err := rr.CreateBaseRole()
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		server := models.Server{
			Name:    p.Name,
			Private: p.Private,
			User_ID: user.ID,
			Roles: []models.Role{
				adminRole,
				baseRole,
			},
		}

		if err := rs.Create(&server); err != nil {
			services.RespondBadRequestError(c, err, "Error creating server")
			return
		}

		if err := rs.Append(user, &server); err != nil {
			services.RespondBadRequestError(c, err, "Error setting user to server")
			return
		}
		// attach server role admin to current user
		_, err = rr.AttachServerRoles([]models.ServerUserRole{
			{
				ServerID: server.ID,
				UserID:   user.ID,
				RoleID:   server.Roles[0].ID,
			},
		})
		if err != nil {
			services.RespondBadRequestError(c, err, "Error attaching role to user")
			return
		}

		services.RespondStatusCreated(c, "server", server)
		return
	}
}

type InviteUserParams struct {
	ServerID  uint `json:"server_id" binding:"required_without=ChannelID"`
	ChannelID uint `json:"channel_id" binding:"required_without=ServerID"`
	UserID    uint `json:"user_id" binding:"required"`
}

func InviteUser(rs repositories.ServerRepo, rp repositories.PermissionRepo, r *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := services.GetUserContext(c)
		if err != nil {
			services.RespondBadRequestError(c, err, "Error getting user from context")
			return
		}

		var p InviteUserParams
		if err := c.ShouldBind(&p); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		var server models.Server
		if p.ServerID != 0 {
			err = rs.Find(p.ServerID, &server)
			if err != nil {
				services.RespondBadRequestError(c, err, err.Error())
				return
			}
		} else {
			// find channel
		}

		err = rs.UserExistsOnServer(&server, user)
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		perms, err := rp.GetUserServerPermissions(user.ID, server.ID)
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}

		requiredPerms, err := rp.InviteUserPermission()
		if err != nil {
			services.RespondBadRequestError(c, err, "Permission's not found")
			return
		}

		err = rp.CanAccess(requiredPerms, perms)
		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return

		}

		serverInviteKey := utils.GetRandomString("", 12)
		ro := services.RedisServerInvite{Key: services.SERVER_INVITE, ServerID: server.ID, UserID: p.UserID}
		// serverInviteValue := fmt.Sprintf("serverID=%v|invitedUserID=%v|key=%v", server.ID, p.UserID, services.SERVER_INVITE)
		if err := services.SetRedisKey(serverInviteKey, r, ro); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}
		// if err := r.Set(context.Background(), serverInviteKey, ro, 24*time.Hour).Err(); err != nil {
		// 	services.RespondBadRequestError(c, err, err.Error())
		// 	return
		// }

		type inviteResponse struct {
			Invite string `json:"invite"`
		}
		services.RespondStatusOK(c, "message", inviteResponse{Invite: serverInviteKey})
		return
	}
}

type JoinServerParams struct {
	ServerID  uint   `json:"server_id" binding:"required_without=InviteKey"`
	InviteKey string `json:"invite_key" binding:"required_without=ServerID"`
}

func JoinServer(rs repositories.ServerRepo, rr repositories.RoleRepo, r *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := services.GetUserContext(c)
		if err != nil {
			services.RespondBadRequestError(c, err, "Error getting user from context")
			return
		}

		var p JoinServerParams
		if err := c.ShouldBind(&p); err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}
		var server models.Server

		if p.InviteKey != "" {
			// key, err := services.GetRedisKey(p.InviteKey, r)
			// if err != nil {
			// 	services.RespondBadRequestError(c, err, err.Error())
			// 	return
			// }
			// data := strings.Split(key, "|")
			// if len(data) != 3 {
			// 	services.RespondBadRequestError(c, errors.New("Server|User|Key was not provided"), "Mismatch data")
			// 	return
			// }

			var rsi services.RedisServerInvite

			if err := services.GetRedisKey(p.InviteKey, r, &rsi); err != nil {
				services.RespondBadRequestError(c, errors.New("Issue getting struct from redis"), "Mismatch data")
				return
			}
			// type redisOutput struct {
			// 	Key      string
			// 	ServerID uint
			// 	UserID   uint
			// }
			// ro := redisOutput{}
			// for _, s := range data {
			// 	if strings.Contains(s, "serverID") {
			// 		t := strings.Split(s, "=")[1]
			// 		v, err := strconv.ParseUint(t, 10, 32)
			// 		if err != nil {
			// 			services.RespondBadRequestError(c, err, err.Error())
			// 			return
			// 		}
			// 		ro.ServerID = uint(v)

			// 	} else if strings.Contains(s, "userID") {
			// 		t := strings.Split(s, "=")[1]
			// 		v, err := strconv.ParseUint(t, 10, 32)
			// 		if err != nil {
			// 			services.RespondBadRequestError(c, err, err.Error())
			// 			return
			// 		}
			// 		ro.UserID = uint(v)

			// 	} else if strings.Contains(s, "key") {
			// 		ro.Key = strings.Split(s, "=")[1]
			// 	}
			// }
			if rsi.Key != services.SERVER_INVITE {
				services.RespondBadRequestError(c, errors.New("Key was not provided"), "Mismatch data")
				return
			}
			if rsi.UserID != user.ID {
				services.RespondBadRequestError(c, errors.New("User tried to join server that he was not invited to"), "Mismatch data")
				return
			}

			if err := rs.Find(rsi.ServerID, &server); err != nil {
				services.RespondBadRequestError(c, err, err.Error())
				return
			}
		} else {
			if err := rs.Find(p.ServerID, &server); err != nil {
				services.RespondBadRequestError(c, err, err.Error())
				return
			}
			if server.Private == true {
				services.RespondBadRequestError(c, fmt.Errorf("UserID %v, is trying to access serverid %v but does not have access", user.ID, server.ID), "Mismatch data")
				return
			}
		}

		if err := rs.Append(user, &server); err != nil {
			services.RespondBadRequestError(c, err, "Error setting user to server")
			return
		}
		_, err = rr.AttachServerRoles([]models.ServerUserRole{
			{
				ServerID: server.ID,
				UserID:   user.ID,
				RoleID:   server.Roles[1].ID, //server base role
			},
		})

		if err != nil {
			services.RespondBadRequestError(c, err, err.Error())
			return
		}
		services.RespondStatusOK(c, "server", server)
		return

	}
}
