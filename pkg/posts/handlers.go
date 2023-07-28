package Posts

import (
	"time"

	"gorm.io/gorm"
)

func CreatedWithinPastDay(db *gorm.DB) *gorm.DB {
	currentTime := time.Now()
	dayInHours := 24 * time.Hour
	return db.Where("? - created_at < ?", currentTime, dayInHours)
}

func SortByUpVotes(db *gorm.DB) *gorm.DB {
	return db.Order("votes")
}
