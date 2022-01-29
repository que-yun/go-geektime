package handler

import (
	userService "error-demo/service"
	"errors"
	"github.com/gin-gonic/gin"
	xerrors "github.com/pkg/errors"
	"net/http"
)

func FindUser(c *gin.Context) error {

	id, ok := c.GetQuery("id")
	if ok == false {
		return errors.New("[userHandler] id is required")
	}
	user, err := userService.FindUserById(id)
	if err != nil {
		return xerrors.WithMessage(err, "[userHandler] FindUser Error")
	}
	c.JSON(http.StatusOK, user)
	return nil
}
