package controllers

import (
	"net/http"
	"strconv"

	"github.com/rcarvalho-pb/mottu-broker_service/internal/helper"
	"github.com/rcarvalho-pb/mottu-broker_service/internal/model"
)

type userController struct{}

func newUserController() *userController {
	return &userController{}
}

func (uc *userController) GetUserById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	user, err := srv.UserSvc.GetUserById(int64(id))
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, user)
}

func (uc *userController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser *model.UserDTO
	if err := helper.ReadJson(w, r, &newUser); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if err := srv.UserSvc.CreateUser(newUser); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusCreated, nil)
}

func (uc *userController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user *model.UserDTO
	if err := helper.ReadJson(w, r, user); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	if err := srv.UserSvc.UpdateUser(user); err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, nil)
}

func (uc *userController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := srv.UserSvc.GetAllUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
}

func (uc *userController) GetAllActiveUsers(w http.ResponseWriter, r *http.Request) {
	users, err := srv.UserSvc.GetAllActiveUsers()
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, users)
}

func (uc *userController) GetUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	user, err := srv.UserSvc.GetUserByUsername(username)
	if err != nil {
		helper.ErrorJson(w, err)
		return
	}
	helper.WriteJson(w, http.StatusOK, user)
}

func (uc *userController) DeleteUserById(w http.ResponseWriter, r *http.Request) {

}
