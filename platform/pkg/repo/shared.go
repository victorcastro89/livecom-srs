package repo

import (
	"time"
)
type CreateUserPayload struct {
	FirstName   *string `json:"first_name,omitempty" validate:"omitempty,min=3,max=25"`
	LastName    *string `json:"last_name,omitempty" validate:"omitempty,min=3,max=25"`
	PhoneNumber *string `json:"phone_number,omitempty" validate:"omitempty,numeric,min=9,max=15"` // Assuming min 9 and max 15 digits for global phone numbers
	PhotoUrl    *string `json:"photo_url,omitempty" validate:"omitempty,url"`
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
}

type RequestError struct {
	Err string
}

type VerifyResponse struct{
	Code   int
	Data   *string
}