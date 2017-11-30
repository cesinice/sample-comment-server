package main

import (
	"github.com/kataras/iris"
	"sample-comment-server/controllers"
	"github.com/kataras/iris/middleware/recover"
	"github.com/spf13/viper"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
	"fmt"
	"os"
	"log"
)

// CommentBook is responsible of the webserver application framework container.
// While Iris is responsible to manage the state of the application.
// It contains and handles all the necessary parts to create a fast web server.
type CommentBook struct {
	app           *iris.Application
	config        *viper.Viper
	db            *gorm.DB
}

// Initializing the configuration file from a sample generated by bindata
// It will check recursively all configuration path defined here and
// read data from config file or create a new one if it does not exists
// TODO : Handle insufficient permissions to write/read case!
func (c *CommentBook) initConfiguration() *CommentBook {

	// Setting up basic configuration information
	configName := "config"
	configType := "json"

	// Setting up the configuration file name
	c.config.SetConfigName(configName)

	// Setting up configuration type
	c.config.SetConfigType(configType)

	// Preparing file copy and deferring file close due to possible
	// memory leak if we deferred the file close in the for loop
	var to *os.File
	defer to.Close()

	// Here you can define your config paths
	// TODO : Maybe move this into the constructor through DI
	configPaths := []string{
		"./",
		"$HOME/.sample-comment-server/",
	}

	// Accessing our sample config asset and catching it as an
	// array of bytes.
	sample, _ := Asset("config.sample.json")

	// For each configuration path, we need to check the existence of a configuration
	// file.
	//
	// This will also create non-existent configuration paths and files
	// TODO : Add a config state, if this state increase to one, then there is no need to continue config search
	// TODO : Move this logic to ReadInConfig error checking instead of panicking!
	for _, element := range configPaths {
		// Adding the configuration path to analysis
		c.config.AddConfigPath(element)
		filePath := element + configName + "." + configType

		// No need to continue further if a configuration file exists!
		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			break
		}

		// Configuration file does not exist, we need to copy our sample to a valid configuration path
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			// If the configuration path itself does not exist, we need to handle that ourselves.
			if _, err := os.Stat(element); os.IsNotExist(err) {
				os.MkdirAll(element, os.ModePerm)
			}

			// We obviously need to try to open with read/write access or simply to create the file
			to, err = os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, os.ModePerm)
			// Then it could go wrong due to permissions..
			// TODO : Handle read/write permissions once for all.
			if err != nil {
				log.Fatal(err)
			}

			// Writing to the file and handle errors
			_, err = to.Write(sample)
			if err != nil {
				log.Fatal(err)
			}

			// Don't need to continue further, only a single configuration file is necessary
			// We don't need one for each paths
			break
		}
	}

	// Reading configuration file
	err := c.config.ReadInConfig()

	// If even though we managed a lot of case, permissions could generate problems
	// if it's the case, panic and run !
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	return c
}

func (c *CommentBook) useGorm() *CommentBook {
	dialect, host, user, dbname, password, sslmode := c.config.GetString("db.dialect"),
	c.config.GetString("db.host"),
	c.config.GetString("db.user"),
	c.config.GetString("db.dbname"),
	c.config.GetString("db.password"),
	c.config.GetString("db.sslmode")

	c.db, _ = gorm.Open(
		dialect,
		"host="+ host +" user="+ user +" dbname="+ dbname +" sslmode="+ sslmode +" password="+ password)

	return c
}

// Running Server
func (c *CommentBook) RunServer() *CommentBook {
	c.app.Run(iris.Addr(":" + c.config.GetString("port")), iris.WithoutInterruptHandler)
	return c
}

// New creates and returns a fresh empty *CommentBook instance.
func New() *CommentBook {
	return &CommentBook{
		app: iris.New(),
		config: viper.New(),
	}
}

// Main is responsible for the application running instance.
func main() {
	// Requesting a new instance of our application
	book := New().initConfiguration().useGorm()

	// This will recover any panic attack go may have!
	book.app.Use(recover.New())

	// Registering a GET Route at HTTP root saying "Hello World" through anonymous function
	book.app.Get("/", func(context iris.Context) {
		context.WriteString("Hello World")
	})

	// Registering a Controller Grouped Route
	// Through an instance of our CommentsController
	book.app.Controller("/comments", new(controllers.CommentsController))

	// Running the web server instance
	book.RunServer()
}
