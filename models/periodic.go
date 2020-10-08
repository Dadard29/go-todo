package models

import "time"

const (
	PeriodicTimeDaily  = "PeriodicTimeDaily"
	PeriodicTimeWeekly = "PeriodicTimeWeek"
)

var PeriodicTimeValue = map[string]time.Duration{
	PeriodicTimeDaily:  time.Hour * 24,
	PeriodicTimeWeekly: time.Hour * 24 * 7,
}

type PeriodicEntity struct {
	Id           int           `gorm:"type:int;index:id;primary_key;auto_increment"`
	OccursAt     time.Time     `gorm:"type:date;index:occurs_at"`
	Title        string        `gorm:"type:string;index:title"`
	Category     string        `gorm:"type:string;index:category"`
	PeriodicTime time.Duration `gorm:"type:int;index:periodic_time"`
}

type PeriodicDto struct {
	Id           int           `json:"id"`
	OccursAt     time.Time     `json:"occurs_at"`
	Title        string        `json:"title"`
	Category     string        `json:"category"`
	PeriodicTime time.Duration `json:"periodic_time"`
}

func (p PeriodicEntity) ToDto() PeriodicDto {
	return PeriodicDto{
		Id:           p.Id,
		OccursAt:     p.OccursAt,
		Title:        p.Title,
		Category:     p.Category,
		PeriodicTime: p.PeriodicTime,
	}
}
