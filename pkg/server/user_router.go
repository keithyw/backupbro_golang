package server

import (
	"backupBro/pkg"
	"github.com/gorilla/mux"
	"net/http"
	"errors"
	"encoding/json"
	"log"
	"gopkg.in/go-playground/validator.v9"
	"regexp"
)

type userRouter struct {
	userService root.UserService
}

func NewUserRouter(u root.UserService, router *mux.Router) *mux.Router {
	userRouter := userRouter{u}
	router.HandleFunc("/register", userRouter.createUserHandler).Methods("POST")
	router.HandleFunc("/email", userRouter.getUserHandler).Methods("POST")
	router.HandleFunc("/login", userRouter.loginHandler).Methods("POST")
	return router
}

func (ur *userRouter) createUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := decodeUser(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	validate := validator.New()
	validate.RegisterValidation("username", username)
	err = validate.Struct(user)
	if err != nil {
		Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	olduser, _ := ur.userService.GetByEmail(user.Email)
	if len(olduser.Email) > 0 {
		Error(w, http.StatusUnprocessableEntity, "Email already exist " + olduser.Email + " old email")
		return
	}
	oldusername, _ := ur.userService.GetByUsername(user.Username)
	if len(oldusername.Username) > 0 {
		Error(w, http.StatusUnprocessableEntity, "Username already exist")
		return
	}

	err = ur.userService.CreateUser(&user)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	newuser, _ := ur.userService.GetByEmail(user.Email)
	Json(w, http.StatusOK, newuser)
}

func (ur *userRouter) getUserHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	if email == "" {
		Error(w, http.StatusBadRequest, "email not set")
		return
	}

	user, err := ur.userService.GetByEmail(email)
	if err != nil {
		Error(w, http.StatusNotFound, err.Error())
		return
	}
	Json(w, http.StatusOK, user)
}

func (ur *userRouter) loginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("loginHandler")
	credentials, err := decodeCredentials(r)
	if err != nil {
		Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	var user root.User
	user, err = ur.userService.Login(credentials)
	if err == nil {
		Json(w, http.StatusOK, user)
	} else {
		Error(w, http.StatusInternalServerError, "Incorrect password")
	}
}

func decodeCredentials(r *http.Request) (root.Credentials, error) {
	var c root.Credentials
	if r.Body == nil {
		return c, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)
	return c, err
}

func decodeUser(r *http.Request) (root.User, error) {
	var u root.User
	if r.Body == nil {
		return u, errors.New("no request body")
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&u)
	return u, err
}

func username(fl validator.FieldLevel) bool {
	if m, _ := regexp.MatchString(`^\w+$`, fl.Field().String()); !m {
		return false
	}
	return true
}