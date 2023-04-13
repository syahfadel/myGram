package routers

import (
	"myGram/controllers"
	"myGram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *gin.Engine {

	myGramService := services.MyGramService{
		DB: db,
	}

	myGramController := controllers.MyGramController{
		DB:            db,
		MyGramService: &myGramService,
	}

	app := gin.Default()
	user := app.Group("/user")
	{
		user.POST("/register", myGramController.Register)
		user.POST("/login", myGramController.Login)
	}

	photo := app.Group("/photo")
	{
		photo.GET("/getAll", myGramController.GetAllPhoto)
		photo.GET("/getOne", myGramController.GetOnePhoto)
		photo.POST("/createPhoto", myGramController.CreatePhoto)
		photo.PUT("/updatePhoto", myGramController.UpdatePhoto)
		photo.DELETE("/deletePhoto", myGramController.DeletePhoto)
	}

	comment := app.Group("/comment")
	{
		comment.GET("/getAll", myGramController.GetAllComment)
		comment.GET("/getOne", myGramController.GetOneComment)
		comment.POST("/createComment", myGramController.CreateComment)
		comment.PUT("/updateComment", myGramController.UpdateComment)
		comment.DELETE("/deleteComment", myGramController.DeleteComment)
	}

	socialMedia := app.Group("/socialMedia")
	{
		socialMedia.GET("/getAll", myGramController.GetAllSocialMedia)
		socialMedia.GET("/getOne", myGramController.GetOneSocialMedia)
		socialMedia.POST("/createSocialMedia", myGramController.CreateSocialMedia)
		socialMedia.PUT("/updateSocialMedia", myGramController.UpdateSocialMedia)
		socialMedia.DELETE("/deleteSocialMedia", myGramController.DeleteSocialMedia)
	}

	return app
}
