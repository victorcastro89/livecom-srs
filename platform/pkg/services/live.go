package services

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"livecom/logger"

	"livecom/pkg/crypto"
	"livecom/pkg/db"
	"livecom/pkg/repo"

	"github.com/jackc/pgx/v5/pgtype"
)
var (
	ErrNotAllowed = errors.New("Not allow to acess this resource")
)
func (s *Service) CreateLive(ctx context.Context, arg repo.CreateLivePayload) (repo.LiveDecrypted, error) {
	randId,err := generateRandomID(20);
	
	if err != nil {
		return repo.LiveDecrypted{}, err
	}
	hashedLiveSecret := crypto.HashMD5(randId);
	if err != nil {
		return repo.LiveDecrypted{}, err
	}
	encryptedLiveSecret,err := crypto.EncryptString(randId)
	if err != nil {
		return repo.LiveDecrypted{}, err
	}
	BroadCastUrl := arg.LiveAppName + "/" + arg.StreamName + "?secret=" + randId;
	encryptedBroadcastUrl,err := crypto.EncryptString(BroadCastUrl)
	if err != nil {
		return repo.LiveDecrypted{}, err
	}
	live :=db.CreateLiveParams{
	
	}
	
	live.AccountID = arg.AccountId
	live.Title = arg.Title
	live.Description.Scan(*arg.Description)
	live.ScheduledStartTime.Scan(*arg.ScheduledStartTime)
	live.ScheduledEndTime.Scan(*arg.ScheduledEndTime)
	live.LiveAppName.Scan(arg.LiveAppName)
	live.StreamName.Scan(arg.StreamName)
	live.LiveSecretHash.Scan(hashedLiveSecret)
	live.LiveSecretEncrypted.Scan(encryptedLiveSecret)
	live.StreamBroadcastUrlEncrypted.Scan(encryptedBroadcastUrl)
	logger.Tf(ctx, "live %v", live)
	createdLive,err:= s.db.CreateLive(ctx, live)
	if err != nil {
		return repo.LiveDecrypted{}, err
	}

	decryptedSecret,err:=crypto.DecryptString(live.LiveSecretEncrypted.String);
	if err != nil {
		return  repo.LiveDecrypted{}, err
	}
	decryptedStreamName,err:=crypto.DecryptString(live.StreamBroadcastUrlEncrypted.String);
	if err != nil {
		return  repo.LiveDecrypted{}, err
	}

	return repo.LiveDecrypted{
		Live: createdLive,
		Decrypted_secret: decryptedSecret,
		Decrypted_stream_name: decryptedStreamName,
	}, nil
	
}

func (s *Service) GetLiveByEncryptedSecretStreamAppName(ctx context.Context, secretHash string, streamName string, appName string) (repo.LiveWithDecryptedSecretAndStreamName, error) {
	var param db.GetLiveBySecretHashAppAndStreamParams
	
	param.StreamName.Scan(streamName);
	param.LiveSecretHash.Scan(secretHash);
	param.LiveAppName.Scan(appName);
	logger.Tf(ctx, "GetLiveBySecretHashAppAndStream %v", param)
	live ,err:=s.db.GetLiveBySecretHashAppAndStream(ctx, param)
	logger.Tf(ctx, "Live %v", live)

	if err != nil {
		logger.Tf(ctx, "Err %s", err.Error())
		return repo.LiveWithDecryptedSecretAndStreamName{}, ErrNotAllowed
	}
	decryptedSecret,err:=crypto.DecryptString(live.LiveSecretEncrypted.String);
	if err != nil {
		return repo.LiveWithDecryptedSecretAndStreamName{}, err
	}
	decryptedStreamName,err:=crypto.DecryptString(live.StreamBroadcastUrlEncrypted.String);
	if err != nil {
		return repo.LiveWithDecryptedSecretAndStreamName{}, err
	}
	return repo.LiveWithDecryptedSecretAndStreamName{
		GetLiveBySecretHashAppAndStreamRow: live,
		Decrypted_secret:decryptedSecret ,
		Decrypted_stream_name: decryptedStreamName,
	},nil
}

func (s *Service) DeleteLive(ctx context.Context, liveID int32) error {
	return s.db.DeleteLive(ctx, liveID)
}



func (s *Service) GetLiveWithStatusByID(ctx context.Context,  liveId int32) (db.GetLiveWithStatusByIDRow, error) {
	return s.db.GetLiveWithStatusByID(ctx, liveId)
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