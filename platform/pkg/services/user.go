package services

import (
	"context"
	"errors"
	"fmt"
	"livecom/logger"
	"livecom/pkg/db"
	"livecom/pkg/repo"
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)


func (s *Service) GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error) {
    return s.db.GetUserByID(ctx, id)
}

func (s *Service) GetUserByFirebaseUID(ctx context.Context, firebaseUid string)(*db.User, error) {
	var uid pgtype.Text
	err:=uid.Scan(firebaseUid);
	if err != nil {
		return nil,err
	}
	u,err := s.db.GetUserByFirebaseUID(ctx, uid)
	if(err!=nil){
        return nil,err
    }
    return &u, nil
}
func (s *Service) CreateUser(ctx context.Context, token *auth.Token ,payload repo.CreateUserPayload)(*db.User, error) {

	params,err :=s.mergeTokenUsrPayload(token, payload)
	if(err!=nil){
        return nil,err
    }
    u,err :=s.db.CreateUser(ctx, *params)
    if(err!=nil){
        return nil,err
    }
    return &u, nil
}


func (s *Service) mergeTokenUsrPayload(token *auth.Token, payload repo.CreateUserPayload)(*db.CreateUserParams, error) {
	validate := validator.New()

		
		err:=validate.Struct(payload)
	
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			return nil,fmt.Errorf("validation error: %s", validationErrors.Error())
	}

	
	var userParams db.CreateUserParams

	if uid, ok := token.Claims["user_id"].(string); ok {
		userParams.FirebaseUid.Scan(uid);
       

	}
	if email, ok := token.Claims["email"].(string); ok {
		userParams.Email = email
	}
    if email, ok := token.Claims["email_verified"].(string); ok {
		userParams.EmailVerified.Scan(email) 
	}
	if name, ok := token.Claims["name"].(string); ok {
		userParams.DisplayName.Scan(name)
	}
	if phoneNumber, ok := token.Claims["phone_number"].(string); ok {
		userParams.PhoneNumber.Scan(phoneNumber)
	}
	if PhotoUrl, ok := token.Claims["photo"].(string); ok {
		userParams.PhotoUrl.Scan(PhotoUrl)
	}
	// Overwrite JWT payload with CreateUserPayload if provided
	if payload.FirstName != nil && strings.TrimSpace(*payload.FirstName) != "" {
	
        userParams.FirstName.Scan(*payload.FirstName);
	}
	if payload.LastName != nil && strings.TrimSpace(*payload.LastName) != "" {
		userParams.LastName.Scan(*payload.LastName);
	}
	if payload.PhoneNumber != nil && strings.TrimSpace(*payload.PhoneNumber) != "" {
		userParams.PhoneNumber.Scan(*payload.PhoneNumber)
	}
	if payload.PhotoUrl != nil && strings.TrimSpace(*payload.PhotoUrl) != "" {
		userParams.PhotoUrl.Scan(*payload.PhotoUrl)
	}

        // Generate a new UUID and convert it to PostgreSQL UUID format
        uid,err :=  uuid.NewRandom()
   
        if err != nil {
            logger.E(nil,err)
                  return nil,err;
        }
         err = userParams.UserID.Scan(uid.String()) 
         if err != nil {
            logger.E(nil,err)
            return nil, fmt.Errorf("Failed Generating/Converting UUID");
                      
        }

    //logger.T(nil,"Creating user new user: %d",userParams)
	return &userParams, nil
}


func validatePayload(payload db.CreateUserParams) error {
	// Add validation logic for CreateUserPayload fields based on API best practices
	// For example, you can validate the format of the phone number, check the length of the name, etc.

	// Sample validation: Check if first name is not too long
	if payload.FirstName.String == "" && len(payload.FirstName.String) > 50 && len(payload.FirstName.String) < 50 {
		return errors.New("first name is too long")
	}

	// Add more validation logic as needed

	return nil
}

func UUIDToString(uuid pgtype.UUID) (string) {

	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:])
}
