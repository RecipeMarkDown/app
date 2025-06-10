// internal/models/user.go
package models

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID        uint
	GoogleID  string
	Email     string
	Username  string
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
