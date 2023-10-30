package handlers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"livecom/logger"
	"livecom/pkg/db"
	"livecom/pkg/repo"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a live entity
// @Description Create a live entity with the given payload
// @ID create-live
// @Accept json
// @Produce json
// @Param Authorization header string true "<access_token>"
// @Param payload body repo.CreateLivePayload true "CreateLivePayload object"
// @Success 200 {object} db.Live "Successfully created"
// @Failure 400 {object} repo.RequestError "Bad request" '{"error": "string"}'
// @Failure 403 {object} repo.RequestError "Forbidden"
// @Failure 500 {object} repo.RequestError "Internal Server Error"
// @Router /webapi/live [post]
func (h *Handlers) CreateLiveHandler(c *gin.Context) {
	val , exists := c.Get("authenticated_user")
	if !exists {
		c.JSON(http.StatusInternalServerError, repo.RequestError{Err: "Error getting Gin context authenticated_user from handlers.CreateLiveHandler"}) 
        return
        }
	user, ok := val.(*db.User)
		if !ok {
			// Handle error: Unexpected type for user
			c.JSON(http.StatusInternalServerError, repo.RequestError{Err:"Failed to process authenticated user"})
			return
		}

	// Just checkin User ID 
	// TO DO Refactory with user access	
	if(user.UserID.Valid){	
    // Bind JSON payload to CreateUserPayload struct
	var livePayload repo.CreateLivePayload
	if err := c.ShouldBindJSON(&livePayload); err != nil {
		c.JSON(http.StatusBadRequest, repo.RequestError{Err: err.Error()})
		return
	}	
	live,err:=h.Service.CreateLive(c, *user, livePayload)	
	if err == nil{
		c.JSON(http.StatusOK, live)
	}else{
		logger.W(c, err.Error())
		c.JSON(http.StatusBadRequest, repo.RequestError{Err: err.Error()})
	}
	}else{
		c.JSON(http.StatusForbidden, repo.RequestError{Err: "Not allow to acess this resource"})
		return
	}
	}

// @Summary Get a live by ID
// @Description Retrieve live details by ID
// @ID get-live
// @Produce json
// @Consume json
// @Param id path string true "Live ID"
// @Param Authorization header string true "<access_token>"
// @Param Content-Type header string true "application/json"
// @Success 200 {object} db.CreateLiveRow "Successfully retrieved live"
// @Failure 400 {object} repo.RequestError "Bad request"
// @Failure 403 {object} repo.RequestError "Forbidden"
// @Failure 500 {object} repo.RequestError "Internal Server Error"
// @Router /webapi/live/{id} [get]
	func (h *Handlers) GetLiveHandler(c *gin.Context) {
		user:=h.GetUserFromContext(c)
		id := c.Param("id")
		i, err := strconv.ParseInt(id, 10, 32) // Base 10, and bit size 32 for int32
		if err != nil {
			c.JSON(http.StatusInternalServerError, repo.RequestError{Err: err.Error()}) 
			return
		}

		h.Service.GetLiveByID(c, *user, int32(i))
		return 

	}
// func (h *Handlers) DeleteLiveHandler(w http.ResponseWriter, r *http.Request) {
// 	idStr := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	err = h.Service.DeleteLive(int32(id))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusNoContent)
// }

// // Similarly, implement handlers for the other functions like GetLiveByID, GetLiveWithStatusByID, etc.

// func (h *Handlers) GetLiveByIDHandler(w http.ResponseWriter, r *http.Request) {
// 	idStr := mux.Vars(r)["id"]
// 	id, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		http.Error(w, "Invalid ID format", http.StatusBadRequest)
// 		return
// 	}

// 	result, err := h.Service.GetLiveByID(int32(id))
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	json.NewEncoder(w).Encode(result)
// }


// @Summary Verify if a live can be streamed
// @Description Method to SRS callback to verify if a live can be streamed, update status
// @ID srs-verify
// @Accept json
// @Produce json
// @Param payload body interface{} true  "Verify object"
// @Success 200 {object} repo.VerifyResponse "Authorized"
// @Failure 400 {object} repo.RequestError "Bad request" '{"error": "string"}'
// @Failure 403 {object} repo.RequestError "Forbidden"
// @Failure 500 {object} repo.RequestError "Internal Server Error"
// @Router /webapi/verify [post]
func (h *Handlers) Verify(c *gin.Context) {
 		//Read the request body
		 bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		 if err != nil {
			 log.Printf("Error reading body: %v", err)
			 c.AbortWithStatus(http.StatusInternalServerError)
			 return
		 }
 
		 // Print the request body
		 fmt.Println(string(bodyBytes))
 
		 // Reset the request body to its original state for other handlers/middlewares
		 c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		c.Header("Content-Type", "application/json")
	
		c.JSON(200, repo.VerifyResponse{
			Code:   0,
			Data:   nil,
		
		})
}
