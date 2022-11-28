package controller

import (
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/maabodi/goApp/domain"
	"github.com/maabodi/goApp/model"
)

type UserController struct {
	SVC domain.UserAdapterService
}

func (uc *UserController) CreateUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	user.Role = "user"

	err, statusCode := uc.SVC.CreateUserService(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
			"status":   statusCode,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success " + string(statusCode),
		"users":    user,
	})
}

func (uc *UserController) GetUsersController(c echo.Context) error {
	users := uc.SVC.GetAllUsersService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users":    users,
	}, "  ")
}

func (uc *UserController) LoginUserController(c echo.Context) error {
	userLogin := make(map[string]interface{})

	c.Bind(&userLogin)

	token, statusCode := uc.SVC.LoginUser(userLogin["email"].(string), userLogin["password"].(string))

	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusBadRequest, map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "email atau password salah",
		}, " ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal error",
		}, " ")

	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"token":   token,
	}, " ")
}

func (uc *UserController) AdminRoute(c echo.Context) error {
	bearer := c.Get("user").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	if string(claim["role"].(string)) != "operator" {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "gagal, ini khusus operator",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Halo berhasil, ini rute operator !",
	})

}
