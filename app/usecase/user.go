package usecase

import (
	"github.com/yamazaki-ko/go-architecture/app/domain/entity"
	"github.com/yamazaki-ko/go-architecture/app/domain/repository"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

func (u *userUseCase) Save(id int64, arg entity.User) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		// 上書き
		user.Name = arg.Name
		user.MailAddress = arg.MailAddress
		user.Password = arg.Password
		if err := u.userRepo.Update(user); err != nil {
			return nil, err
		}
	} else {
		// 新規作成
		user = &entity.User{
			ID:          id,
			Name:        arg.Name,
			MailAddress: arg.MailAddress,
			Password:    arg.Password,
		}
		if err := u.userRepo.Create(user); err != nil {
			return nil, err
		}
	}
	// 再取得
	ret, err := u.userRepo.FindByID(id)
	return ret, err
}
