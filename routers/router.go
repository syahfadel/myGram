package routers

import (
	_ "myGram/docs"

	ginSwagger "github.com/swaggo/gin-swagger"

	swaggerFiles "github.com/swaggo/files"

	"myGram/controllers"
	"myGram/middlewares"
	"myGram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @title           MyGram API
// @version         1.0
// @description     This is a API for MyGram application.
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4000
// @BasePath  /
func StartService(db *gorm.DB) *gin.Engine {

	myGramService := services.MyGramService{
		DB: db,
	}

	myGramController := controllers.MyGramController{
		DB:            db,
		MyGramService: &myGramService,
	}

	app := gin.Default()
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	user := app.Group("/user")
	{
		// Create
		user.POST("/register", myGramController.Register)
		// Read
		user.POST("/login", myGramController.Login)
	}

	photo := app.Group("/photo")
	{
		app.Use(middlewares.Authentication())
		// Read All
		photo.GET("/getAll", myGramController.GetAllPhoto)
		// Read
		photo.GET("/getOne/:id", myGramController.GetOnePhoto)
		// Create
		photo.POST("/createPhoto", myGramController.CreatePhoto)
		// Update
		photo.PUT("/updatePhoto/:id", myGramController.UpdatePhoto)
		// DELETE
		photo.DELETE("/deletePhoto/:id", myGramController.DeletePhoto)
	}

	comment := app.Group("/comment")
	{
		app.Use(middlewares.Authentication())
		// Read All
		comment.GET("/getAll", myGramController.GetAllComment)
		// Read
		comment.GET("/getOne/:id", myGramController.GetOneComment)
		// Create
		comment.POST("/createComment", myGramController.CreateComment)
		// Update
		comment.PUT("/updateComment/:id", myGramController.UpdateComment)
		// DELETE
		comment.DELETE("/deleteComment/:id", myGramController.DeleteComment)
	}

	socialMedia := app.Group("/socialMedia")
	{
		app.Use(middlewares.Authentication())
		// Read All
		socialMedia.GET("/getAll", myGramController.GetAllSocialMedia)
		// Read
		socialMedia.GET("/getOne/:id", myGramController.GetOneSocialMedia)
		// Create
		socialMedia.POST("/createSocialMedia", myGramController.CreateSocialMedia)
		// Update
		socialMedia.PUT("/updateSocialMedia/:id", myGramController.UpdateSocialMedia)
		// DELETE
		socialMedia.DELETE("/deleteSocialMedia/:id", myGramController.DeleteSocialMedia)
	}

	return app
}
