package db

import (
	"github.com/yamazaki-ko/go-architecture/app/domain/entity"
	"github.com/yamazaki-ko/go-architecture/app/domain/repository"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) FindByID(id int64) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.Where("id=?", id).Find(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Create(user *entity.User) error {
	return u.db.Create(&user).Find(user).Error
}

func (u *UserRepository) Update(user *entity.User) error {
	return u.db.Model(&user).Updates(user).Error
}
