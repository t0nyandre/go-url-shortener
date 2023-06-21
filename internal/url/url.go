package url

import (
	"time"

	"github.com/gookit/validate"
)

type Url struct {
	ID       int32  `json:"id" db:"id"`
	ShortUrl string `json:"short_url" db:"short_url"`
	LongUrl  string `json:"long_url" db:"long_url"`
	Clicks   int32  `json:"clicks" db:"clicks"`
	// Future implementation with user accounts etc.
	// UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func (u *Url) Validate() error {
	v := validate.Struct(u)
	if v.Validate() {
		return nil
	}
	return v.Errors
}
