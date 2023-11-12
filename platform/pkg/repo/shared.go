package repo

import (
	"livecom/pkg/db"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type UserResponse struct {
	db.GetUserWithRoleAndAccountByIDRow
	Roles []db.GetAccountsAndRolesByUserIDRow
}
type CreateUserPayload struct {
	FirstName   *string `json:"first_name,omitempty" validate:"omitempty,min=3,max=25"`
	LastName    *string `json:"last_name,omitempty" validate:"omitempty,min=3,max=25"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,numeric,min=9,max=15"` // Assuming min 9 and max 15 digits for global phone numbers
	PhotoUrl    *string `json:"photo_url,omitempty" validate:"omitempty,url"`
	AccountName string	`json:"account_name" validate:"omitempty,min=3,max=50"`
	Role        string	`json:"role" validate:"min=3,max=50"`
}
func (p *CreateUserPayload) AnyFieldNil() bool {
    return p.FirstName == nil || p.LastName == nil || p.PhoneNumber == nil || p.PhotoUrl == nil
}
type CreateLivePayload struct  {
Title              string `json:"title" validate:"min=3,max=50"`
Description        *string `json:"description" validate:"omitempty min=10"`
ScheduledStartTime *time.Time `json:"scheduled_start_time,omitempty" validate:"omitempty datetime=2006-01-02T15:04:05Z07:00"`
ScheduledEndTime   *time.Time	`json:"scheduled_end_time,omitempty" validate:"omitempty datetime=2006-01-02T15:04:05Z07:00"`
LiveAppName        string `json:"live_app_name" validate:"min=3,max=50"`
StreamName         string `json:"stream_name" validate:"min=3,max=50"`
AccountId          pgtype.UUID `json:"account_id" validate:"required"`
}

type RequestError struct {
	Err string
}

type VerifyResponse struct{
	Code   int `json:"code"`
	Data   *string `json:"data"`
}

type BroadCastVerify struct {
	ServerID  *string `json:"server_id,omitempty"`
	ServiceID *string `json:"service_id,omitempty"`
	Action    string `json:"action"  binding:"required"`
	ClientID  *string `json:"client_id,omitempty"`
	IP        *string `json:"ip,omitempty"`
	Vhost     *string `json:"vhost,omitempty"`
	App       string `json:"app"`
	TcURL     *string `json:"tcUrl,omitempty"`
	Stream    string `json:"stream"  binding:"required"`
	Param     string `json:"param,omitempty"  `
	StreamURL *string `json:"stream_url,omitempty"`
	StreamID  *string `json:"stream_id,omitempty"`
}

type LiveWithDecryptedSecretAndStreamName struct {
	db.GetLiveBySecretHashAppAndStreamRow
	Decrypted_secret string
	Decrypted_stream_name string

}
type LiveDecrypted struct {
	db.Live
	Decrypted_secret string
	Decrypted_stream_name string
}