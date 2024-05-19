// internal/model/model.go
package model

import "time"

type URL struct {
	ID        string
	Original  string
	Short     string
	CreatedAt time.Time
	ExpiresAt time.Time
	Clicks    int
}
