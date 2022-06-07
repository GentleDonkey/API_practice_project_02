package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"problem3_irdeto/internal/model"
	"strconv"

	"problem3_irdeto/internal/error"
)

type Handler interface {
	RegisterRoute(*mux.Router)
	GetNotes(http.ResponseWriter, *http.Request)
	GetNoteByID(http.ResponseWriter, *http.Request)
	Creat(http.ResponseWriter, *http.Request)
	NewHttpResponse(http.ResponseWriter, model.HttpResponse)
}

type apiHandler struct {
	client Client
}

func NewAPIHandler(c Client) Handler {
	return &apiHandler{
		c,
	}
}

func (handler *apiHandler) RegisterRoute(router *mux.Router) {
	router.HandleFunc("/notes", handler.GetNotes).Methods("GET")
	router.HandleFunc("/notes/{id}", handler.GetNoteByID).Methods("GET")
	router.HandleFunc("/notes", handler.Creat).Methods("POST")
}

func (handler *apiHandler) GetNotes (w http.ResponseWriter, r *http.Request) {
	response := model.HttpResponse{}
	response = *handler.client.GetNotes()
	handler.NewHttpResponse(w, response)
	return
}

func (handler *apiHandler) GetNoteByID (w http.ResponseWriter, r *http.Request) {
	noteId := mux.Vars(r)["id"]
	id, err := strconv.Atoi(noteId)
	if noteId == "" || err != nil {
		response := model.HttpResponse{}
		response.MyError = error.InvalidPara
		handler.NewHttpResponse(w, response)
		return
	}
	response := model.HttpResponse{}
	response = *handler.client.GetNoteByID(id)
	handler.NewHttpResponse(w, response)
	return
}

func (handler *apiHandler) Creat (w http.ResponseWriter, r *http.Request) {
	var note model.Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		response := model.HttpResponse{}
		response.MyError = *error.NewError(err, "Unable to convert a json to an object", 400)
		handler.NewHttpResponse(w, response)
		return
	}
	response := model.HttpResponse{}
	response = *handler.client.Creat(note)
	w.WriteHeader(201)
	handler.NewHttpResponse(w, response)
	return
}

func (handler *apiHandler) NewHttpResponse(w http.ResponseWriter, response model.HttpResponse) {
	if response.MyError != (error.MyError{}) {
		jsonNewResp, err := json.Marshal(response.MyError)
		if err != nil {
			temp := error.JsonMarshalError
			temp1, _ := json.Marshal(temp)
			w.WriteHeader(temp.StatusCode)
			w.Write(temp1)
			return
		}
		w.WriteHeader(response.MyError.StatusCode)
		w.Write(jsonNewResp)
		return
	}
	jsonNewResp, err := json.Marshal(response.Result)
	if err != nil {
		temp := error.JsonMarshalError
		temp1, _ := json.Marshal(temp)
		w.WriteHeader(temp.StatusCode)
		w.Write(temp1)
		return
	}
	//w.WriteHeader(http.StatusOK)
	w.Write(jsonNewResp)
	return
}