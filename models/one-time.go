package models

import "time"

type OneTimeEntity struct {
	Id       int       `gorm:"type:int;index:id;primary_key;auto_increment"`
	DueAt    time.Time `gorm:"type:date;index:due_at"`
	Title    string    `gorm:"type:string;index:title"`
	Category string    `gorm:"type:string;index:category"`
}

func (o OneTimeEntity) ToDto() OneTimeDto {
	return OneTimeDto{
		Id:       o.Id,
		DueAt:    o.DueAt,
		Title:    o.Title,
		Category: o.Category,
	}
}

type OneTimeDto struct {
	Id       int       `json:"id"`
	DueAt    time.Time `json:"due_at"`
	Title    string    `json:"title"`
	Category string    `json:"category"`
}

func (o OneTimeDto) ToEntity() OneTimeEntity {
	return OneTimeEntity{
		Id:       o.Id,
		DueAt:    o.DueAt,
		Title:    o.Title,
		Category: o.Category,
	}
}
