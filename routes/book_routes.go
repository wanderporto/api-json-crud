package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/wanderporto/api-json-crud/models"

	"io/ioutil"

	"github.com/gorilla/mux"
)

func addBookHandler(r *mux.Router) {

	r.HandleFunc("/ping", pingHandler).Methods("GET")
	r.HandleFunc("/books", booksGetHandler).Methods("GET")
	r.HandleFunc("/books", booksPostHanlder).Methods("POST")
	r.HandleFunc("/books/{id}", bookGetHandler).Methods("GET")
	r.HandleFunc("/books/{id}", bookPutHandler).Methods("PUT")
	r.HandleFunc("/books/{id}", bookDeleteHandler).Methods("DELETE")

}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(struct {
		Error  string
		Status int
	}{
		Error:  "PING",
		Status: 200,
	})
}

func booksGetHandler(w http.ResponseWriter, r *http.Request) {
	httpInfo(r)
	books, err := models.GetBooks()

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

	json.NewEncoder(w).Encode(books)
}

func bookGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, _ := strconv.Atoi(vars["id"])

	book, err := models.GetBook(id)

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

	encoder := json.NewEncoder(w)
	encoder.Encode(book)

}

func booksPostHanlder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}

	var book models.Book

	err = json.Unmarshal(body, &book)

	_, errorCreated := models.NewBook(book)

	if errorCreated != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})

		return
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(book)
}

func bookPutHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var book models.Book

	book.Id = id

	body, _ := ioutil.ReadAll(r.Body)

	err := json.Unmarshal(body, &book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})

		return
	}

	rows, err := models.UpdateBook(book)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})

		return
	}

	json.NewEncoder(w).Encode(struct{ RowsAffected int64 }{RowsAffected: rows})
}

func bookDeleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	rows, err := models.DeleteBook(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct {
			Error  string
			Status int
		}{
			Error:  "UNPROCESSABLE ENTITY",
			Status: 422,
		})

		return
	}

	json.NewEncoder(w).Encode(struct{ RowsAffected int64 }{RowsAffected: rows})
}
