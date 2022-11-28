package service

import (
	"net/http"

	"github.com/myApp/config"
	"github.com/myApp/domain"
	"github.com/myApp/helper"
	"github.com/myApp/model"
)

type svcUser struct {
	c    config.Config
	repo domain.UserAdapterRepository
}

func NewServiceUser(repo domain.UserAdapterRepository, c config.Config) domain.UserAdapterService {
	return &svcUser{
		repo: repo,
		c:    c,
	}
}
func (su *svcUser) CreateUserService(user model.User) (error, int) {

	return su.repo.CreateUsers(user), http.StatusOK
}

func (s *svcUser) GetAllUsersService() []model.User {
	return s.repo.GetAll()
}

func (s *svcUser) LoginUser(email, password string) (string, int) {
	if len(email) == 0 {
		return "email atau password tidak boleh kosong", http.StatusBadRequest
	}
	user, _ := s.repo.GetOneByEmail(email)

	if (user.Password != password) && (user.Email != email) {
		return "", http.StatusUnauthorized
	}

	token, err := helper.CreateToken(int(user.UserID), user.Email, user.Role, s.c.JWT_KEY)

	if err != nil {
		return "", http.StatusInternalServerError
	}
	return token, http.StatusOK
}
