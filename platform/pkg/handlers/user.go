package handlers

import (
	"livecom/logger"
	"livecom/pkg/db"
	"livecom/pkg/firebaseauth"
	"livecom/pkg/helpers"
	"livecom/pkg/repo"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handlers) Test(c *gin.Context)  {
  u,err:= h.Service.GetUserByFirebaseUID(c, "test")
  logger.T(c,"user ",u)
  logger.T(c,"error ",err)
  c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
  return
}


func (h *Handlers) GetAccountsAndRolesByUser(c *gin.Context) {
    var accounts []db.GetAccountsAndRolesByUserIDRow
    user := h.GetUserFromContext(c)
    if user == nil {
         return
    }
    logger.T(c,"user ",user.UserID)
    acc, err := h.Service.GetAccountsAndRolesByUser(c, user.UserID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Account not found"})
        return
    }
    for _,account := range acc{
        if(h.UserHaveAcessToResource(c, GetAccounts, account.AccountID , nil)){
            accounts = append(accounts, account)
        }
    }
    if accounts == nil {
        c.JSON(http.StatusForbidden, gin.H{"error": "Not allow to acess this resource"})
        return
    }

    c.JSON(http.StatusOK, accounts)
}
type RequestBody struct {
    Role      string `json:"role"`
    UserID    string `json:"user_id"`
    AccountID string `json:"account_id"`
}
func (h *Handlers) AddUserToAccount(c *gin.Context) {
    var requestBody RequestBody
    if err := c.ShouldBindJSON(&requestBody); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
 result,err:= h.Service.AddUserToAccount(c, helpers.ConvertToPgUUID(uuid.MustParse(requestBody.UserID)), helpers.ConvertToPgUUID(uuid.MustParse(requestBody.AccountID)), requestBody.Role)
 if err != nil {
    logger.E(c, "Error adding user to account: %v",err)
    c.JSON(http.StatusBadRequest, err.Error())
    return 
 }
 c.JSON(http.StatusOK, result)

}
// @Summary Get a user by ID
// @Description Retrieve user details by ID
// @ID get-user
// @Produce json
// @Consume json
// @Param id path string true "User ID"
// @Param Authorization header string true "<access_token>"
// @Param Content-Type header string true "application/json"
// @Success 200 {object} db.User "Successfully retrieved user"
// @Failure 400 {object} repo.RequestError "Bad request"
// @Failure 403 {object} repo.RequestError "Forbidden"
// @Failure 500 {object} repo.RequestError "Internal Server Error"
// @Router /users/{id} [get]
func (h *Handlers) GetUser(c *gin.Context) {
    id := c.Param("id")

    userID, err := uuid.Parse(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
        return
    }

    pgID := helpers.ConvertToPgUUID(userID)

    user, err := h.Service.GetUserByID(c, pgID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

// @Summary Create a user
// @Description Create a new user
// @ID create-user
// @Accept json
// @Produce json
// @Param Authorization header string true "<access_token>"
// @Param payload body repo.CreateUserPayload true "CreateUserPayload  object"
// @Success 200 {object} db.User "Successfully created"
// @Failure 400 {object} repo.RequestError "Bad request" '{"error": "string"}'
// @Failure 403 {object} repo.RequestError "Forbidden"
// @Failure 500 {object} repo.RequestError "Internal Server Error"
// @Router /webapi/users [post]
func (h *Handlers) CreateOrGetUserWithAccountAndRole(c *gin.Context) {
    // Bind JSON payload to CreateUserPayload struct
    var payload repo.CreateUserPayload
    if err := c.ShouldBindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    //Use Firebase Auth to verify an ID token   
	firetoken := c.GetHeader("Authorization")
	firebase := firebaseauth.GetInstance()
	token, err := firebase.VerifyIDToken(firetoken )
    if err != nil {
		logger.T(nil,err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		
		return
	}
 

    
  

    // Create user using the service
    user, err := h.Service.CreateOrGetUserAccountRole(c, token, payload)
    if err != nil {
		logger.E(nil,err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Return JSON response with success message and user details
    c.JSON(http.StatusOK, user)
}
