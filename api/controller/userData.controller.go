package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/NEHA20-1992/task_1/api/formattor"
	"github.com/NEHA20-1992/task_1/api/model"
	"github.com/NEHA20-1992/task_1/api/response"
	"github.com/gorilla/mux"
	// "golang.org/x/crypto/bcrypt"
)

var formattedError=errors.New("email already exist")

func (server *Server) CreateUser(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user := model.UserData{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	userCreated, err := user.SaveUser(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
		
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, userCreated.ID))
	response.JSON(w, http.StatusCreated, userCreated)
}
func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {

	user := model.UserData{}

	users, err := user.FindAllUsers(server.DB)
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	response.JSON(w, http.StatusOK, users)
}
func (server *Server) GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	user := model.UserData{}
	userGotten, err := user.FindUserByID(server.DB, uint32(uid))
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	response.JSON(w, http.StatusOK, userGotten)
}
func (server *Server) UpdateUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := model.UserData{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user.Prepare()
	err = user.Validate("update")
	if err != nil {
		response.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedUser, err := user.UpdateAUser(server.DB, uint32(uid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		response.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	response.JSON(w, http.StatusOK, updatedUser)
}
func (server *Server) DeleteUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	user := model.UserData{}

	uid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		response.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = user.DeleteAUser(server.DB, uint32(uid))
	if err != nil {
		response.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", uid))
	response.JSON(w, http.StatusNoContent, "")
}





