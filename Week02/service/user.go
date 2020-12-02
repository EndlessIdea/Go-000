package service

import (
	"context"
	"github.com/EndlessIdea/Go-000/Week02/dao"
)

var (
	UserService = &userService{}
)

type userService struct {}

func (us *userService) GetUserByID(ctx context.Context, id int) (*dao.User, error) {
	return dao.UserDao.GetByID(ctx, id)
}