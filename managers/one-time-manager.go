package managers

import (
	"github.com/Dadard29/go-todo/models"
	"github.com/Dadard29/go-todo/repositories"
)

func OneTimeManagerGet(id int) (models.OneTimeDto, error) {
	var f models.OneTimeDto

	o, err := repositories.OneTimeGet(id)
	if err != nil {
		return f, err
	}

	return o.ToDto(), nil
}

func OneTimeManagerList() ([]models.OneTimeDto, error) {
	var f []models.OneTimeDto

	le, err := repositories.OneTimeList()
	if err != nil {
		return f, nil
	}

	var ld = make([]models.OneTimeDto, 0)
	for _, v := range le {
		ld = append(ld, v.ToDto())
	}

	return ld, nil
}

func OneTimeManagerCreate(o models.OneTimeInput) (models.OneTimeDto, error) {
	var f models.OneTimeDto

	e, err := repositories.OneTimeCreate(o.ToEntity())
	if err != nil {
		return f, err
	}

	return e.ToDto(), nil
}

func OneTimeManagerDelete(id int) (models.OneTimeDto, error) {
	var f models.OneTimeDto

	e, err := repositories.OneTimeDelete(id)
	if err != nil {
		return f, err
	}

	return e.ToDto(), nil
}

func OneTimeManagerUpdate(o models.OneTimeInput, id int) (models.OneTimeDto, error) {
	var f models.OneTimeDto

	e, err := repositories.OneTimePut(o.ToEntity(), id)
	if err != nil {
		return f, err
	}

	return e.ToDto(), nil
}
