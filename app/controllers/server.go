package controllers

import (
	"discord-clone-server/models"
	"discord-clone-server/repositories"

	"github.com/gin-gonic/gin"
)

type serverCreateParams struct {
	Name    string `json:"name" binding:"required,min=3,max=25"`
	Private bool   `json:"private"`
}

func ServerCreate(rs repositories.ServerRepo, rp repositories.PermissionRepo, ru repositories.UserRepo, rr repositories.RoleRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := GetUserContext(c)
		if err != nil {
			RespondBadRequestError(c, err, "Error getting user from context")
			return
		}
		var p serverCreateParams
		if err := c.ShouldBind(&p); err != nil {
			RespondBadRequestError(c, err, err.Error())
			return
		}

		permission, err := rp.Find(repositories.PermissionFindParams{Permission: "admin"})
		if err != nil {
			RespondBadRequestError(c, err, err.Error())
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
			RespondBadRequestError(c, err, "Error creating server")
			return
		}

		if err := rs.Append(&user, server); err != nil {
			RespondBadRequestError(c, err, "Error setting user to server")
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

		RespondStatusCreated(c, "server", server)
		return
	}
}

func InviteUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
