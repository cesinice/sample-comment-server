package controllers

import (
	"time"
	"encoding/json"
)

// Comments Controller extends our BaseController.
// This allows to use Iris reflection and controllers methods.
// Also to access to our database through gORM.
type CommentsController struct {
	BaseController
}

// CommentResponse is the basic data structure for a comment Response
// Then parsing it to JSON data through json.Marshall
type CommentResponse struct {
	Id            int        `json:"id"`
	Content       string     `json:"content"`
	Author        string     `json:"author"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

// This handle the GET requests through the Comments Controller
// TODO : Implement data retrieval through database!
func (c *CommentsController) Get() []byte {

	// Mock Comment Response
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

	// Formatting our Response Structure to JSON.
	response, err := json.Marshal(users)

	// In case of Marshal errors, please panic.
	if err != nil {
		panic(err)
	}

	return response
}

