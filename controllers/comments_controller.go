package controllers

import (
	"time"
	"encoding/json"
	"sample-comment-server/models"
)

// Comments Controller extends our BaseController.
// This allows to use Iris reflection and controllers methods.
// Also to access to our database through gORM.
type CommentsController struct {
	BaseController
}

// CommentResponse is the basic data structure for a comment Response
// Then parsing it to JSON data through json.Marshall
type CommentData struct {
	Id            int        `json:"id,omitempty"`
	Content       string     `json:"content"`
	Author        string     `json:"author"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
	DeletedAt     *time.Time `json:"deleted_at,omitempty"`
}

// This handle the GET requests through the Comments Controller
// TODO : Implement data retrieval through database!
func (c *CommentsController) Get() []byte {

	var comments models.Comment

	// Mock Comment Response
	users := c.DB.Find(&comments).Value

	// Formatting our Response Structure to JSON.
	response, err := json.Marshal(users)

	// In case of Marshal errors, please panic.
	if err != nil {
		panic(err)
	}

	return response
}

// This allows to post a comment
func (c *CommentsController) Post() []byte {
	var jsonRequest CommentData
	c.Ctx.ReadJSON(&jsonRequest)

	comment := models.Comment{Content: jsonRequest.Content, Author: jsonRequest.Author, CreatedAt: time.Now()}
	result, _ := json.Marshal(c.DB.Create(&comment).Value)

	return result
}

// This method allows to delete a comment using its id
func (c *CommentsController) DeleteBy(id int64) {
	// We need to refer to a model to use the ORM and the request data.
	var jsonRequest CommentData
	var model models.Comment

	// Reading JSON Request
	c.Ctx.ReadJSON(&jsonRequest)

	// Deleting the entry if its possible
	c.DB.First(&model, id).Delete(model)
}