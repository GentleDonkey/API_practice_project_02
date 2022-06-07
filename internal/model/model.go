package model

import (
	"problem3_irdeto/internal/error"
)

type Note struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	When int `json:"when"`
}

type Result struct {
	Notes []Note `json:"notes"`
}

type HttpResponse struct {
	MyError error.MyError `json:"myError"`
	Result  Result        `json:"result"`
}