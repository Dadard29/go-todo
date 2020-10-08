package models

import "time"

type OneTimeEntity struct {
	Id       int       `gorm:"type:int;index:id;primary_key;auto_increment"`
	CreatedAt time.Time `gorm:"type:datetime;index:created_at"`
	DueAt    time.Time `gorm:"type:datetime;index:due_at"`
	Title    string    `gorm:"type:string;index:title"`
	Category string    `gorm:"type:string;index:category"`
}

func (OneTimeEntity) TableName() string {
	return "one_time"
}

func (o OneTimeEntity) ToDto() OneTimeDto {
	return OneTimeDto{
		Id:       o.Id,
		CreatedAt: o.CreatedAt,
		DueAt:    o.DueAt,
		Title:    o.Title,
		Category: o.Category,
	}
}

type OneTimeDto struct {
	Id       int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	DueAt    time.Time `json:"due_at"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
}

type OneTimeInput struct {
	DueAt    time.Time `json:"due_at"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
}

func (o OneTimeInput) ToEntity() OneTimeEntity {
	return OneTimeEntity{
		DueAt:    o.DueAt,
		Title:    o.Title,
		Category: o.Category,
	}
}
