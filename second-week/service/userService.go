package userService

import (
	userDao "error-demo/dao"
	xerrors "github.com/pkg/errors"
)

func FindUserById(id string) (*userDao.User, error) {
	user, err := userDao.FindUserById(id)
	if err != nil {
		return nil, xerrors.Wrap(err, "[userService] FindUserById Error")
	}
	return user, nil
}
