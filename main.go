package main

import (
	"github.com/kataras/iris"
)

// CommentBook is responsible of the webserver application framework container.
// While Iris is responsible to manage the state of the application.
// It contains and handles all the necessary parts to create a fast web server.
type CommentBook struct {
	app *iris.Application
}

// New creates and returns a fresh empty *CommentBook instance.
func New() *CommentBook {
	return &CommentBook{
		app: iris.New(),
	}
}

// Main is responsible for the application running instance.
func main() {
	// Requesting a new instance of our application
	book := New()

	// Registering a GET Route at HTTP root saying "Hello World" through anonymous function
	book.app.Get("/", func(context iris.Context) {
		context.WriteString("Hello World")
	})

	// Running the web server instance at the port 8080
	book.app.Run(iris.Addr(":8080"))
}
