package api

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"problem3_irdeto/internal/error"
	"problem3_irdeto/internal/model"
	"strings"
	"testing"
)

type fields struct {
	client *client
}
type args struct {
	w http.ResponseWriter
	r *http.Request
}
func TestHandler(t *testing.T){
	var tests1 = []struct {
		name                 string
		url				     string
		client               func() Client
		expectedResponseCode int
		expectedBody         *model.Result

	}{
		{
			name: "GetNotes1",
			url: "http://localhost:8000/notes",
			client: func() Client {
				response := &model.HttpResponse{
					Result: model.Result{
						Notes: []model.Note{
							{0, "Note to self", "Take out the garbage", 1624631009},
							{1, "Note to self 1", "Take out the garbage 1", 1624631009},
							{2, "Note to self 2", "Take out the garbage 2", 1624631009},
							{3, "Note to self 3", "Take out the garbage 3", 1624631009},
						},
					},
					MyError: error.MyError{},
				}
				mockClient := MockClient{}
				mockClient.On("GetNotes").Return(response).Once()
				return mockClient
			},
			expectedResponseCode: 200,
			expectedBody: &model.Result{
				Notes: []model.Note{
					{0, "Note to self", "Take out the garbage", 1624631009},
					{1, "Note to self 1", "Take out the garbage 1", 1624631009},
					{2, "Note to self 2", "Take out the garbage 2", 1624631009},
					{3, "Note to self 3", "Take out the garbage 3", 1624631009},
				},
			},
		},
	}
	for _, tt := range tests1 {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.client()
			api := NewAPIHandler(client)
			r, _ := http.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()
			api.GetNotes(w, r)
			require.Equal(t, tt.expectedResponseCode, w.Code)
			if tt.expectedBody != nil {
				expectedJSON, err := json.Marshal(tt.expectedBody)
				require.NoError(t, err)
				require.Equal(t, string(expectedJSON), strings.Trim(string(w.Body.Bytes()), "\n"))
			}
		})
	}
	var tests2 = []struct {
		name                 string
		url                  string
		client              func() Client
		expectedResponseCode int
		expectedBody         *error.MyError
	}{
		{
			name: "GetNotes2",
			url: "http://localhost:8000/notes",
			client: func() Client {
				response := &model.HttpResponse{
					Result: model.Result{
						Notes: []model.Note{},
					},
					MyError: error.MyError{
						errors.New("error"), "Can't find the record", 400,
					},
				}
				mockClient := MockClient{}
				mockClient.On("GetNotes").Return(response).Once()
				return mockClient
			},
			expectedResponseCode: 400,
			expectedBody: &error.MyError{
				errors.New("error"), "Can't find the record", 400,
			},
		},
	}
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			client := tt.client()
			api := NewAPIHandler(client)
			r, _ := http.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()
			api.GetNotes(w, r)
			require.Equal(t, tt.expectedResponseCode, w.Code)
			if tt.expectedBody != nil {
				expectedJSON, err := json.Marshal(tt.expectedBody)
				require.NoError(t, err)
				require.Equal(t, string(expectedJSON), strings.Trim(string(w.Body.Bytes()), "\n"))
			}
		})
	}
}



