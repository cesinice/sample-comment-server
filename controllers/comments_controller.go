package controllers

import (
	"time"
	"encoding/json"
)

type CommentsController struct {
	BaseController
}

// CommentResponse is the basic data structure for a comment Response
type CommentResponse struct {
	Id            int        `json:"id"`
	Content       string     `json:"content"`
	Author        string     `json:"author"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

// This handle the GET requests through the Comments Controller
func (c *CommentsController) Get() []byte {
	users := &[]CommentResponse{
		{
			Id: 1,
			Content: `Meilleure conférence de l'Exia ! J'avoue, je suis un peu biaisé..`,
			Author: "Tony BRIET",
			CreatedAt: time.Now(),
		},
		{
			Id: 2,
			Content: `Mon dieu, il est mauvais !`,
			Author: "Antoine Orfila",
			CreatedAt: time.Now(),
		},
	}

	response, err := json.Marshal(users)

	if err != nil {
		panic(err)
	}

	return response
}

