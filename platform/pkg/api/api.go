package api

import (
	"context"
	"livecom/logger"
	"livecom/pkg/firebaseauth"
	"livecom/pkg/handlers"

	"net/http"

	"github.com/gin-gonic/gin"
)


func SetupWebApiRoutes(ctx context.Context, handlers handlers.Handlers ,mux *http.ServeMux) error {
    router := gin.Default()
	router.Use(authMiddleware(handlers))

	router.POST("/live", handlers.CreateLiveHandler)
    router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
    router.GET("/hello", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"userID":123,
		})
	})
	router.Any("/verify", handlers.Verify)
 
    mux.Handle("/webapi/", http.StripPrefix("/webapi", router))


    return nil
}



func authMiddleware(handlers handlers.Handlers) gin.HandlerFunc {
	return func(c *gin.Context) {
	// Check if the request is POST /users
		if c.Request.Method == "POST" && c.Request.URL.Path == "/users" {
			// Skip the action and proceed to the next handler
			c.Next()
			return
		}
		if c.Request.URL.Path == "/verify" {
			// Skip the action and proceed to the next handler
			c.Next()
			return
		}

		firetoken := c.GetHeader("Authorization")
 		firebase := firebaseauth.GetInstance()

		// Example: Use Firebase Auth to verify an ID token
		token, err := firebase.VerifyIDToken(firetoken)
		if err != nil {
			logger.Tf(nil,"Error verifying ID token: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":err.Error()})
			return
		}
		var firebase_uid string
		if uid, ok := token.Claims["user_id"].(string); ok {
			firebase_uid = uid
		}else{
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Error getting Firebase UID"})
			return
		}
		user, err:=handlers.Service.GetUserByFirebaseUID(c, firebase_uid)
		if err != nil {
			logger.Tf(nil,"Error finding user by firebase UID: %v\n", err)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error finding user by firebase UID: " + err.Error()})
			return
		}
		if user == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"User not found"})
			return
		}
		logger.Tf(c,"Authenticated user: %v\n", user)
		
		// Store the authenticated user in the gin context
		c.Set("authenticated_user", user)
		c.Next()
	}
}


// type AuthParams struct {
// 	AccessToken string `json:"accessToken" binding:"required"`
// }

// func handleWebApiService(ctx context.Context, handler *http.ServeMux) error {
// 	r := gin.Default()
// 	r.Use(authMiddleware())


// 	r.GET("/asd", func(c *gin.Context) {
// 		firetoken := c.GetHeader("FireAuth")
// 		// var params AuthParams
// 		// if err := c.ShouldBindJSON(&params); err != nil {
// 		// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		// 	return
// 		// }
// 		firebase := firebaseauth.GetInstance()

// 		// Example: Use Firebase Auth to verify an ID token
// 			token, err := firebase.VerifyIDToken(firetoken )
// 		if err != nil {
// 			fmt.Printf("Error verifying ID token: %v\n", err)
// 			c.JSON(401, gin.H{"error":err.Error()})
// 			return
// 		}
	

// 		c.JSON(200, token)
// 	})
	
// 	handler.Handle("/webapi/", http.StripPrefix("/webapi", r))

// return nil
// }
