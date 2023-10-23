package api

import (
	"livecom/pkg/domain/users"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler  *users.UserHandler) *gin.Engine {
    router := gin.Default()



    router.GET("/users/:id", userHandler.GetUser)

    // Add other routes here

    return router
}

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
