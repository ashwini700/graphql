package services

import (
	"context"

	gormmodel "graphql/graph/gorm"
	"graphql/graph/model"
)

func (s *DbConnStruct) CreateUser(ctx context.Context, nu model.NewUser) (*model.User, error) {

	u := gormmodel.NewUser{
		Name:  nu.Name,
		Email: nu.Email,
	}
	err := s.db.Create(&u).Error
	if err != nil {
		return &model.User{}, err
	}
	return &model.User{
		Name:  u.Name,
		Email: u.Email,
	}, nil
}

func (s *DbConnStruct) ViewUser(ctx context.Context) ([]*model.User, error) {
	var listUser []*gormmodel.NewUser
	tx := s.db.WithContext(ctx)
	err := tx.Find(&listUser).Error
	if err != nil {
		return nil, err
	}
	var mUser []*model.User
	for _, user := range listUser {
		newUser := model.User{
			Name:  user.Name,
			Email: user.Email,
		}
		mUser = append(mUser, &newUser)
	}

	return mUser, nil
}
