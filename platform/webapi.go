package main

import (
	"context"
	"fmt"
	"net/http"

	"livecom/pkg/firebaseauth"

	"github.com/gin-gonic/gin"
)



func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" || token != "Bearer YOUR_VALID_TOKEN" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		c.Next()
	}
}


type AuthParams struct {
	AccessToken string `json:"accessToken" binding:"required"`
}

func handleWebApiService(ctx context.Context, handler *http.ServeMux) error {
	r := gin.Default()
	r.Use(authMiddleware())


	r.GET("/asd", func(c *gin.Context) {
		firetoken := c.GetHeader("FireAuth")
		// var params AuthParams
		// if err := c.ShouldBindJSON(&params); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// 	return
		// }
		firebase := firebaseauth.GetInstance()

		// Example: Use Firebase Auth to verify an ID token
			token, err := firebase.VerifyIDToken(firetoken )
		if err != nil {
			fmt.Printf("Error verifying ID token: %v\n", err)
			c.JSON(401, gin.H{"error":err.Error()})
			return
		}
	

		c.JSON(200, token)
	})
	
	handler.Handle("/webapi/", http.StripPrefix("/webapi", r))

return nil
}
