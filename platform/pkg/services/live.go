package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"livecom/logger"
	"livecom/pkg/config"
	"livecom/pkg/db"
	"livecom/pkg/repo"

	"github.com/jackc/pgx/v5/pgtype"
)
func (s *Service) CreateLive(ctx context.Context, usr db.User, arg repo.CreateLivePayload) (db.CreateLiveRow, error) {
	randId,err := generateRandomID(20);
	if err != nil {
		return db.CreateLiveRow{}, err
	}

	live :=db.CreateLiveParams{
	
	}
	live.UserID = usr.UserID
	live.Title = arg.Title
	live.Description.Scan(*arg.Description)
	live.ScheduledStartTime.Scan(*arg.ScheduledStartTime)
	live.ScheduledEndTime.Scan(*arg.ScheduledEndTime)
	live.LiveAppName.Scan(arg.LiveAppName)
	live.StreamName.Scan(arg.StreamName)
	live.LiveSecret = randId
	live.Encryptionkey= config.Cfg.EncryptionKey
	live.StreamBroadcastUrl = arg.LiveAppName + "/" + arg.StreamName + "/" + randId
	logger.Tf(ctx, "live %v", live)
	return s.db.CreateLive(ctx, live)
}

func (s *Service) DeleteLive(ctx context.Context, liveID int32) error {
	return s.db.DeleteLive(ctx, liveID)
}

func (s *Service) GetLiveByID(ctx context.Context, user db.User, liveId int32) (db.GetLiveByIDRow, error) {
	args:=db.GetLiveByIDParams{
		LiveID: liveId,
		Encryptionkey:config.Cfg.EncryptionKey,
	}
	live,err:= s.db.GetLiveByID(ctx, args)
	if err != nil {
		return db.GetLiveByIDRow{}, err
	}

	if(user.UserID != live.UserID ){
	return  db.GetLiveByIDRow{}, errors.New("Not allow to acess this resource")
	}else {
		return live, nil
	}
	


}

func (s *Service) GetLiveWithStatusByID(ctx context.Context, arg db.GetLiveWithStatusByIDParams) (db.GetLiveWithStatusByIDRow, error) {
	return s.db.GetLiveWithStatusByID(ctx, arg)
}

func (s *Service) GetLiveWithUserDetails(ctx context.Context, arg db.GetLiveWithUserDetailsParams) (db.GetLiveWithUserDetailsRow, error) {
	return s.db.GetLiveWithUserDetails(ctx, arg)
}

func (s *Service) GetLivesByUserID(ctx context.Context, userID pgtype.UUID) ([]db.Live, error) {
	return s.db.GetLivesByUserID(ctx, userID)
}

func (s *Service) GetOngoingLives(ctx context.Context) ([]db.Live, error) {
	return s.db.GetOngoingLives(ctx)
}

func (s *Service) UpdateLive(ctx context.Context, arg db.UpdateLiveParams) (db.Live, error) {
	return s.db.UpdateLive(ctx, arg)
}

func generateRandomID(length int) (string, error) {
	// Determine the number of random bytes needed to achieve the desired length
	numBytes := (length * 6) / 8 // 6 bits per character

	// Generate random bytes
	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	// Encode random bytes to a base64 string
	randomID := base64.RawURLEncoding.EncodeToString(randomBytes)

	// Trim to the desired length
	if len(randomID) > length {
		randomID = randomID[:length]
	}

	return randomID, nil
}