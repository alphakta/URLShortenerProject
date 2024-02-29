package url

import (
	"time"
)

type Url struct {
	ID        string    `json:"id"`
	ShortURL  string    `json:"short_url"`
	LongURL   string    `json:"long_url"`
	CreatedAt time.Time `json:"created_at"`
}
