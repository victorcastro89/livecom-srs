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

// Define constants for user roles
const (
    Owner  = "owner"
    Admin  = "admin"
    CoHost = "cohost"
)
var (
	ErrInvalidRole = errors.New("Invalid role")
)

// Function to validate a user role
func IsValidUserRole(role string) bool {
    switch role {
    case Owner, Admin, CoHost:
        return true
    default:
        return false
    }
}


func (s *Service) AddUserToAccount(ctx context.Context, userId pgtype.UUID , accountId pgtype.UUID, role string) (db.UserAccountRole, error) {
	return s.db.CreateUserAccountRoleRelation(ctx, db.CreateUserAccountRoleRelationParams{
		UserID: userId,
		AccountID: accountId,
		Role: pgtype.Text{String: role, Valid: true},
	})
}

func (s *Service) GetAccountsAndRolesByUser(ctx context.Context, userId pgtype.UUID) ([]db.GetAccountsAndRolesByUserIDRow, error) {
	return s.db.GetAccountsAndRolesByUserID(ctx, userId)
}
func (s *Service) GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error) {
    return s.db.GetUserByID(ctx, id)
}

func (s *Service) GetUserByFirebaseUID(ctx context.Context, firebaseUid string)(*db.GetUserWithRoleAndAccountByFirebaseUIDRow, error) {
	var uid pgtype.Text
	err:=uid.Scan(firebaseUid);
	if err != nil {
		return nil,err
	}
	u,err := s.db.GetUserWithRoleAndAccountByFirebaseUID(ctx, uid)
	if(err!=nil){
        return nil,err
    }
    return &u, nil
}


func (s *Service) GetUserAccountAndRoleRelation(ctx context.Context, userID pgtype.UUID) ([]db.UserAccountRole, error) {
	return s.db.GetUserAccountAndRoleRelation(ctx, userID)
}

func (s *Service) CreateOrGetUserAccountRole(ctx context.Context, token *auth.Token ,payload repo.CreateUserPayload) (	db.GetUserWithRoleAndAccountByIDRow, error) {
	if(IsValidUserRole(payload.Role)==false){
		return db.GetUserWithRoleAndAccountByIDRow{},ErrInvalidRole
	}
	logger.T(ctx,"CreateUserAccountRole ",token.UID ) 
	uexists,_:=s.GetUserByFirebaseUID(ctx, token.UID) 

	if(uexists!=nil){
		return db.GetUserWithRoleAndAccountByIDRow{
			UserID:uexists.UserID,
		FirebaseUid:uexists.FirebaseUid,
		Email:uexists.Email,
		EmailVerified:uexists.EmailVerified,
		FirstName:uexists.FirstName,
		LastName :uexists.LastName,
		DisplayName : uexists.DisplayName,
		PhotoUrl:  uexists.PhotoUrl,
		PhoneNumber: uexists.PhoneNumber ,
		Role : uexists.Role,
		AccountID: uexists.AccountID,
		AccountName:uexists.AccountName,
	
		} ,nil
	
	}
	tx,err :=s.dbCon.Begin(ctx);
	if(err!=nil){
		return 	db.GetUserWithRoleAndAccountByIDRow{},err
	}
	defer tx.Rollback(ctx)
	qtx := s.db.WithTx(tx)
	u, err := s.createUserTx(ctx,qtx, token, payload)
	if err != nil {
		return db.GetUserWithRoleAndAccountByIDRow{},err
	}
	accUid,err:=GeneratePsqlUUID();
	if err != nil {
		return db.GetUserWithRoleAndAccountByIDRow{},err
	}
	a,err:= qtx.CreateAccount(ctx, db.CreateAccountParams{
		AccountID:  *accUid,
		AccountName: pgtype.Text{String:payload.AccountName , Valid:true},
	})
	if err != nil {
		logger.E(ctx,"Create account error" ,err)
		return db.GetUserWithRoleAndAccountByIDRow{},err
	}

	userAccountRole,err := qtx.CreateUserAccountRoleRelation(ctx, db.CreateUserAccountRoleRelationParams{
		UserID:  u.UserID,
		AccountID: a.AccountID,
		Role:  pgtype.Text{String:payload.Role , Valid:true},
	})

	if err != nil {
		logger.E(ctx,"userAccountRole" ,err)
		return db.GetUserWithRoleAndAccountByIDRow{},err
	}
	err = tx.Commit(ctx)
	if err != nil {
		logger.E(ctx,"Commit error" ,err)
		return db.GetUserWithRoleAndAccountByIDRow{},err
	}

	return db.GetUserWithRoleAndAccountByIDRow{
	    UserID:u.UserID,
    FirebaseUid:u.FirebaseUid,
    Email:u.Email,
    EmailVerified:u.EmailVerified,
    FirstName:u.FirstName,
    LastName :u.LastName,
    DisplayName : u.DisplayName,
    PhotoUrl:  u.PhotoUrl,
    PhoneNumber: u.PhoneNumber ,
    Role : userAccountRole.Role,
    AccountID: a.AccountID,
    AccountName:a.AccountName,

	} ,nil

}

func (s *Service) createUserTx(ctx context.Context,qtx *db.Queries, token *auth.Token ,payload repo.CreateUserPayload)(*db.User, error) {

	params,err :=s.mergeTokenUsrPayload(token, payload)
	if(err!=nil){
        return nil,err
    }
    u,err :=qtx.CreateUser(ctx, *params)
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

func GeneratePsqlUUID() (*pgtype.UUID,error) {
	uid,err :=  uuid.NewRandom()
	pgUid := pgtype.UUID{}
	if err != nil {
		logger.E(nil,err)
			  return nil,err;
	}
	err = pgUid.Scan(uid.String()) 
	if err != nil {
		logger.E(nil,err)
		return nil,err;
	}
	return &pgUid, nil
}