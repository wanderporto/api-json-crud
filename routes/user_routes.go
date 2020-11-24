package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wanderporto/api-json-crud/models"

	"io/ioutil"

	"github.com/gorilla/mux"
)

var user models.User

func addUserHandler(r *mux.Router) {

	r.HandleFunc("/users", usersGetHandler).Methods("GET")
	r.HandleFunc("/users", usersPostHanlder).Methods("POST")
	r.HandleFunc("/users/{id}", userGetHandler).Methods("GET")
	r.HandleFunc("/users/{id}", userPutHandler).Methods("PUT")
	r.HandleFunc("/users/{id}", userDeleteHandler).Methods("DELETE")
	r.HandleFunc("/login", loginPostHandler).Methods("POST")

}

func usersGetHandler(w http.ResponseWriter, r *http.Request) {
	httpInfo(r)
	users, err := models.GetUsers()

	if err != nil {
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "BAD REQUEST",
			Status: 400,
		})

		return
	}

	json.NewEncoder(w).Encode(users)
}

func userGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	user, err := models.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "NOT FOUND",
			Status: 404,
		})

		return
	}

	json.NewEncoder(w).Encode(user)

}

func usersPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var user models.User

	err = json.Unmarshal(body, &user)

	_, errorCreated := models.NewUser(user)

	if errorCreated != nil {

		ResponseError("UNPROCESSABLE ENTITY", w)

		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(user)
}

func userPutHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var user models.User

	user.Id = id

	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &user)

	if err != nil {
		ResponseError("UNPROCESSABLE ENTITY", w)

		return
	}

	rows, err := models.UpdateUser(user)

	if err != nil {
		ResponseError("UNPROCESSABLE ENTITY", w)
		return
	}

	json.NewEncoder(w).Encode(struct{ RowsAffected int64 }{RowsAffected: rows})
}

func userDeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	rows, err := models.DeleteUser(id)

	if err != nil {
		ResponseError("UNPROCESSABLE ENTITY", w)
		return
	}

	json.NewEncoder(w).Encode(struct{ RowsAffected int64 }{RowsAffected: rows})
}

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	var user models.User

	err := json.Unmarshal(body, &user)

	if err != nil {
		ResponseError("UNPROCESSABLE ENTITY", w)

		return
	}

	user, err = models.Signin(user.Email.String, user.Password.String)

	if err != nil {
		ResponseError("UNAUTHORIZED", w)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func validateLogin(user models.User) bool {

	return true
}

func ResponseError(msg string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(struct {
		Error  string
		Status int
	}{
		Error:  msg,
		Status: 422,
	})
}
