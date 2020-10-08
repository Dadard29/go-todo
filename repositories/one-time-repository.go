package repositories

import (
	"errors"
	"fmt"
	"github.com/Dadard29/go-todo/api"
	"github.com/Dadard29/go-todo/models"
	"time"
)

func oneTimeExists(id int) bool {
	var o models.OneTimeEntity
	api.Api.Database.Orm.Where(&models.OneTimeEntity{
		Id: id,
	}).First(&o)

	return o.Id == id
}

func oneTimeGetFromCreationDate(createdAt time.Time) (models.OneTimeEntity, error) {
	var f models.OneTimeEntity

	ol, err := OneTimeList()
	if err != nil {
		return f, err
	}

	for _, o := range ol {
		if o.CreatedAt.Equal(createdAt) {
			return o, nil
		}
	}

	return f, errors.New("one-time task not found in list")
}

func OneTimePut(p models.OneTimeEntity, id int) (models.OneTimeEntity, error) {
	var f models.OneTimeEntity

	o, err := OneTimeGet(id)
	if err != nil {
		return f, err
	}

	o.DueAt = p.DueAt
	o.Title = p.Title
	o.Category = p.Category

	api.Api.Database.Orm.Save(&o)

	return o, nil
}

func OneTimeDelete(id int) (models.OneTimeEntity, error) {
	var f models.OneTimeEntity

	o, err := OneTimeGet(id)
	if err != nil {
		return f, nil
	}

	api.Api.Database.Orm.Delete(&o)

	if oneTimeExists(id) {
		msg := fmt.Sprintf("failed to delete OneTime task with id %d", id)
		return f, errors.New(msg)
	}

	return o, nil
}

func OneTimeGet(id int) (models.OneTimeEntity, error) {
	var f models.OneTimeEntity

	var o models.OneTimeEntity
	api.Api.Database.Orm.Where(&models.OneTimeEntity{
		Id: id,
	}).First(&o)

	if o.Id != id {
		msg := fmt.Sprintf("OneTime task with id %d not found", id)
		return f, errors.New(msg)
	}

	return o, nil
}

func OneTimeCreate(o models.OneTimeEntity) (models.OneTimeEntity, error) {
	var f models.OneTimeEntity

	createdAt := time.Now().UTC().Round(time.Second)

	api.Api.Database.Orm.Create(&models.OneTimeEntity{
		CreatedAt: createdAt,
		DueAt:    o.DueAt,
		Title:    o.Title,
		Category: o.Category,
	})

	oc, err := oneTimeGetFromCreationDate(createdAt)
	if err != nil {
		return f, err
	}

	return oc, nil
}

func OneTimeList() ([]models.OneTimeEntity, error) {
	var o []models.OneTimeEntity
	api.Api.Database.Orm.Find(&o)

	return o, nil
}
