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
func (h *Handlers) CreateUser(c *gin.Context) {
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
    user, err := h.Service.CreateUser(c, token, payload)
    if err != nil {
		logger.E(nil,err)
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Return JSON response with success message and user details
    c.JSON(http.StatusOK, user)
}

func (h *Handlers) GetUserFromContext(c *gin.Context) (*db.User) {
	val , exists := c.Get("authenticated_user")
	if !exists {
		c.JSON(http.StatusInternalServerError, repo.RequestError{Err: "Error getting Gin context authenticated_user from handlers.CreateLiveHandler"}) 
        return nil
        }
		user, ok := val.(*db.User)
		if !ok {
			// Handle error: Unexpected type for user
			c.JSON(http.StatusInternalServerError, repo.RequestError{Err:"Failed to process authenticated user"})
			return nil
		}
        return user
}