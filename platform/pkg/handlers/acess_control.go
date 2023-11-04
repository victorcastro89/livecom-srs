package handlers

import (
	"livecom/logger"
	"livecom/pkg/db"
	"livecom/pkg/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

const (
    GetLives  = "getLives"
    GetLive  = "getLive"
	GetAccounts  = "getAccounts"
 
)
func (h *Handlers) GetUserFromContext(c *gin.Context) (*db.User) {
	val , exists := c.Get("authenticated_user")
	if !exists {
		c.JSON(http.StatusInternalServerError, repo.RequestError{Err: "Error getting Gin context authenticated_user from handlers.CreateLiveHandler"}) 
        return nil
        }
	logger.T(c,"authenticated_user %+v ",val)	
		user, ok := val.(*db.User)
		if !ok {
			// Handle error: Unexpected type for user
			c.JSON(http.StatusInternalServerError, repo.RequestError{Err:"Failed to process authenticated user"})
			return nil
		}
        return user
}



func (h *Handlers) UserHaveAcessToResource(c *gin.Context, resourceName string, accountId pgtype.UUID,  itemToCheck interface{}) (bool) {
    u:=h.GetUserFromContext(c)
    if u == nil {
        return false
    }
	switch resourceName {
        case GetLives:
            return h.userHasAcessToGetLives(c,*u,accountId )
        case GetLive:    
        return h.userHasAcessToGetLive(c,*u,accountId )
		case GetAccounts:
			return h.userHasAcessToGetAccounts(c,*u)
        default:
            return false
    }
}
func (h *Handlers) userHasAcessToGetAccounts(c *gin.Context,u db.User)(bool) {
	roles:= h.getAccountRolesFromContext(c);
	for _, role := range roles {
		if role.UserID == u.UserID && ( role.Role.String == "admin" ||  role.Role.String == "owner" )   {
			return true
		}
	}
	return false
}
func (h *Handlers) userHasAcessToGetLives(c *gin.Context,u db.User,accountId pgtype.UUID )(bool) {
    roles:= h.getAccountRolesFromContext(c);
    for _, role := range roles {
        if role.UserID == u.UserID && role.AccountID == accountId && ( role.Role.String == "admin" ||  role.Role.String == "owner" )   {
            return true
        }
    }
    return false
}


func (h *Handlers) userHasAcessToGetLive(c *gin.Context,u db.User,accountId pgtype.UUID)(bool) {
    roles:= h.getAccountRolesFromContext(c);
    for _, role := range roles {
        if role.UserID == u.UserID && role.AccountID == accountId && ( role.Role.String == "admin" ||  role.Role.String == "owner" )   {
            return true
        }
    }
    return false
}
func (h *Handlers) getAccountRolesFromContext(c *gin.Context) ([]db.UserAccountRole) {
	res , exists := c.Get("accounts_roles")
	if !exists {
		c.JSON(http.StatusInternalServerError, repo.RequestError{Err: "Error getting Gin context accounts_roles from GetAccountRolesFromContext"}) 
        return nil
        }
		roles, ok := res.([]db.UserAccountRole)
		if !ok {
			// Handle error: Unexpected type for user
			c.JSON(http.StatusForbidden, repo.RequestError{Err:"User do not have access to this resource"})
			return nil
		}
        return roles
}