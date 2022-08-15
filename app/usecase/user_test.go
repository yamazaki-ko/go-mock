package usecase

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yamazaki-ko/go-architecture/app/domain/entity"
	"github.com/yamazaki-ko/go-architecture/app/domain/repository/mock"
)

type userUsecaseTestSet struct {
	name     string
	findUser entity.User
	user     entity.User
	m        *mock.MockUserRepository
}

func setUserUsecaseTesting(t *testing.T, ctrl *gomock.Controller) [2]userUsecaseTestSet {
	// モックセット
	m1 := mock.NewMockUserRepository(ctrl)
	m2 := mock.NewMockUserRepository(ctrl)

	// テストデータ設定
	u := [2]userUsecaseTestSet{}

	// 設定値
	u[0].name = "No1（正常テスト:Create）"
	u[0].findUser.ID = int64(123)
	u[0].user.ID = u[0].findUser.ID
	u[0].user.Name = "hoge"
	u[0].user.MailAddress = "hoge@myfuns.com"
	u[0].user.Password = "password_hoge"

	m1.EXPECT().FindByID(gomock.Eq(u[0].findUser.ID)).Return(nil, nil)
	m1.EXPECT().Create(gomock.Eq(&u[0].user)).Return(nil)
	m1.EXPECT().FindByID(gomock.Eq(u[0].user.ID)).Return(&u[0].user, nil)
	u[0].m = m1

	u[1].name = "No2（正常テスト:Update）"
	u[1].findUser.ID = int64(123)
	u[1].findUser.Name = "hoge"
	u[1].findUser.MailAddress = "hoge@myfuns.co"
	u[1].findUser.Password = "password_hoge"
	u[1].user.ID = u[0].findUser.ID
	u[1].user.Name = "fuga"
	u[1].user.MailAddress = "fuga@myfuns.com"
	u[1].user.Password = "password_fuga"

	m2.EXPECT().FindByID(gomock.Eq(u[1].findUser.ID)).Return(&u[1].findUser, nil)
	m2.EXPECT().Update(gomock.Eq(&u[1].user)).Return(nil)
	m2.EXPECT().FindByID(gomock.Eq(u[1].user.ID)).Return(&u[1].user, nil)
	u[1].m = m2

	return u
}

func Test_userUseCase_Save(t *testing.T) {
	// モックコントローラー生成
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := int64(123)
	tests := setUserUsecaseTesting(t, ctrl)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userUseCase{userRepo: tt.m}
			got, err := u.Save(id, tt.user)
			if err != nil {
				t.Errorf("予期せぬエラーが発生 %v", err)
			}
			if !reflect.DeepEqual(got, &tt.user) {
				t.Errorf("保存結果が期待値と異なる\n期待:%+v\n実際:%+v", &tt.user, got)
			}
		})
	}
}
