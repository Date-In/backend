package user

import (
	"dating_service/internal/model"
	"dating_service/pkg/db"
	"errors"
	"gorm.io/gorm"
)

type Repository struct {
	db *db.Db
}

func NewRepository(db *db.Db) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(user *model.User) error {
	return r.db.PgDb.Create(user).Error
}

func (r *Repository) FindByPhone(phone string) (*model.User, error) {
	var user model.User
	err := r.db.PgDb.Where("phone = ?", phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.PgDb.
		Preload("Interests").
		Preload("Photos", "is_avatar = ?", false).
		Preload("Avatar", "is_avatar = ?", true).
		Where("id = ?", id).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &user, err
}

func (r *Repository) Update(id uint, updateData *model.User) error {
	result := r.db.PgDb.
		Model(&model.User{}).
		Where("id = ?", id).
		Updates(updateData)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *Repository) ReplaceInterests(userID uint, interests []*model.Interest) error {
	var user model.User
	user.ID = userID
	err := r.db.PgDb.Model(&user).Association("Interests").Replace(interests)
	return err
}

func (r *Repository) FindUsersWithFilter(minAge, maxAge, sexID uint, location string, page, pageSize int) ([]*model.User, int64, error) {
	var users []*model.User
	query := r.db.PgDb.Model(&model.User{})
	var totalCount int64
	offset := (page - 1) * pageSize
	query = query.Where("age BETWEEN ? AND ?", minAge, maxAge).
		Where("sex_id = ?", sexID)
	if location != "" {
		query = query.Where("city = ?", location)
	}
	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}
	query = query.
		Offset(offset).
		Limit(pageSize).
		Preload("Photos").
		Preload("Interests")
	result := query.Find(&users)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return users, totalCount, nil
}

func (r *Repository) FindUserWithoutEntity(userId uint) (*model.User, error) {
	var user *model.User
	err := r.db.PgDb.First(&user, "id = ?", userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r *Repository) ChangeStatusUsers(ids []uint) error {
	return r.db.PgDb.Model(&model.User{}).Where("id IN (?) AND status_id != 3", ids).Update("status_id", 2).Error
}
func (r *Repository) ReactivateUser(userID uint) error {
	result := r.db.PgDb.Model(&model.User{}).
		Where("id = ? AND status_id = ? AND status_id != 3", userID, 2).
		Update("status_id", 1)
	return result.Error
}

func (r *Repository) GetStatusUser(id uint) (uint, error) {
	var statusID uint
	err := r.db.PgDb.Model(&model.User{}).
		Select("status_id").
		Where("id = ?", id).
		Scan(&statusID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return statusID, nil
}
