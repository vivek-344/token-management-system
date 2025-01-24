package models

import "time"

// Token represents a token in the system
type Token struct {
	TokenID     string
	UsageCount  int
	LastUpdated time.Time
}
