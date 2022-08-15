package repository

import "github.com/yamazaki-ko/go-architecture/app/domain/entity"

type UserRepository interface {
	FindByID(id int64) (*entity.User, error)
	Create(user *entity.User) error
	Update(user *entity.User) error
}
