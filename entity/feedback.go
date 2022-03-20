package entity

import (
	"time"
)

type Feedback struct {
	Content       string    `json:"content,omitempty" validate:"required"`
	CreatedDate   time.Time `json:"createdDate,omitempty"`
	ApplicationId string    `json:"applicationId,omitempty"`
}
