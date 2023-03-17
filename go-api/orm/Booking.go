package orm

import(
	"time"

	"gorm.io/gorm" // farmwork ต่อกับ database ภาษา Go
)

type Booking struct {
	gorm.Model
	UserID string
	CarID string
	Start time.Time
	End time.Time
}