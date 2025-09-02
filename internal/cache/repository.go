package cache

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db: db}
}

func (r *Repository) LoadSexes() (map[uint]model.Sex, error) {
	var items []model.Sex
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.Sex, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadEducations() (map[uint]model.Education, error) {
	var items []model.Education
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.Education, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadZodiacSigns() (map[uint]model.ZodiacSign, error) {
	var items []model.ZodiacSign
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.ZodiacSign, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadWorldviews() (map[uint]model.Worldview, error) {
	var items []model.Worldview
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.Worldview, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadTypeOfDating() (map[uint]model.TypeOfDating, error) {
	var items []model.TypeOfDating
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.TypeOfDating, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadAttitudeToAlcohol() (map[uint]model.AttitudeToAlcohol, error) {
	var items []model.AttitudeToAlcohol
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.AttitudeToAlcohol, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadAttitudeToSmoking() (map[uint]model.AttitudeToSmoking, error) {
	var items []model.AttitudeToSmoking
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.AttitudeToSmoking, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadStatuses() (map[uint]model.Status, error) {
	var items []model.Status
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.Status, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}

func (r *Repository) LoadInterests() (map[uint]model.Interest, error) {
	var items []model.Interest
	if err := r.db.PgDb.Find(&items).Error; err != nil {
		return nil, err
	}
	res := make(map[uint]model.Interest, len(items))
	for _, item := range items {
		res[item.ID] = item
	}
	return res, nil
}
