package users

import (
	"fmt"
	"livecom/pkg/db"
	"livecom/pkg/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
    UserService *Service
}

type CreateUserPayload struct {
	Email     string `json:"email" binding:"required"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}


func NewUserHandler(s *Service) *UserHandler {
    return &UserHandler{UserService: s}
}

func (h *UserHandler) GetUser(c *gin.Context) {
     userUUID := c.Param("id")
	 uid, err := uuid.Parse(userUUID);
 
     if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	// Convert google/uuid to pgx/v5/pgtype.UUID

    pgUUID:= helpers.ConvertToPgUUID(uid);
    user, err := h.UserService.GetUserByID(c, pgUUID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
  


    var payload CreateUserPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    uid:= uuid.New();
    pgUUID:= helpers.ConvertToPgUUID(uid);
    // Translate CreateUserPayload to CreateUserParams
    userParams := db.CreateUserParams{
        UserID:  pgUUID ,
        Email:     payload.Email,

    }
    
    fmt.Println(userParams)
    user, err := h.UserService.CreateUser(c,userParams)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": err})
        return
    }
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}