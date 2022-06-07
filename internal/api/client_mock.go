package api

import (
	"github.com/stretchr/testify/mock"
	"problem3_irdeto/internal/model"
)

type MockClient struct {
	mock.Mock
}

func (m MockClient) GetNotes() *model.HttpResponse {
	args := m.Called()
	res, _ := args.Get(0).(*model.HttpResponse)
	return res
}

func (m MockClient) GetNoteByID(id int) *model.HttpResponse {
	args := m.Called(id)
	res, _ := args.Get(0).(*model.HttpResponse)
	return res
}

func (m MockClient) Creat(n model.Note) *model.HttpResponse {
	args := m.Called(n)
	res, _ := args.Get(0).(*model.HttpResponse)
	return res
}