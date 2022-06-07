package api

import (
	"errors"
	"problem3_irdeto/internal/error"
	"problem3_irdeto/internal/model"
)

type Client interface {
	GetNotes() *model.HttpResponse
	GetNoteByID(id int) *model.HttpResponse
	Creat(n model.Note) *model.HttpResponse
}

type client struct {}

func NewAPIClient() *client {
	return &client{}
}

func (c client) GetNotes() *model.HttpResponse {
	notes := []model.Note{
			{0, "Note to self", "Take out the garbage", 1624631009},
			{1, "Note to self 1", "Take out the garbage 1", 1624631009},
	}
	if len(notes) != 0 {
		response := &model.HttpResponse{
			Result: model.Result{
				Notes: notes,
			},
			MyError: error.MyError{},
		}
		return response
	}else{
		response := &model.HttpResponse{
			Result: model.Result{
				Notes: notes,
			},
			MyError: error.MyError{
				errors.New("error"), "Can't find the record", 400,
			},
		}
		return response
	}
}

func (c client) GetNoteByID(id int) *model.HttpResponse {
	notes := []model.Note{
		{0, "Note to self", "Take out the garbage", 1624631009},
		{1, "Note to self 1", "Take out the garbage 1", 1624631009},
	}
	for _, item := range notes{
		if item.ID == id{
			response := &model.HttpResponse{
				Result: model.Result{
					Notes: []model.Note{item},
				},
				MyError: error.MyError{},
			}
			return response
		}
	}
	response := &model.HttpResponse{
		Result: model.Result{
			Notes: []model.Note{},
		},
		MyError: error.MyError{
			errors.New("error"), "Can't find the record", 400,
		},
	}
	return response
}

func (c client) Creat (n model.Note) *model.HttpResponse {
	response := &model.HttpResponse{
		Result: model.Result{
			Notes: []model.Note{},
		},
		MyError: error.MyError{},
	}
	return response
}