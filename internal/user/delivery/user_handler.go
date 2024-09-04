package delivery

import (
	user "basiccrud/internal/user/entity"
	"basiccrud/internal/user/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUseCase usecase.UserUseCase
}

type UserUserCase interface {
	RegisterUserHandler(c echo.Context) error 
	LoginUserHandler(c echo.Context) error 
}

func (u *UserHandler) RegisterUserHandler(c echo.Context) error {
	var user user.UserRegister
	if err := c.Bind(&user); err != nil {

		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "bad reques",
		})

	}
	ctx := c.Request().Context()
	if err := u.userUseCase.RegisterUser(ctx, &user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "error in the data",
		})

	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "created",
	})

}

func (u *UserHandler) LoginUserHandler(c echo.Context) error {
	var userLogin user.UserLogin
	if err := c.Bind(&userLogin); err != nil {

		return c.JSON(400, map[string]string{"Error": "binding error"})

	}

	user, err := u.userUseCase.Login(&userLogin)
	if err != nil {
		return c.JSON(500, map[string]string{"Error": err.Error()})

	}
	response := map[string]interface{}{
		"data": user,
	}
	return c.JSON(200, response)
}

func UserDelivery(userUseCase usecase.UserUseCase) UserUserCase {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}
