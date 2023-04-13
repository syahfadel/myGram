package routers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartService(db *gorm.DB) *gin.Engine {
	app := gin.Default()
	user := app.Group("/user")
	{
		user.POST("/register", mygramController.Register)
		user.POST("/login", mygramController.Login)
	}

	photo := app.Group("/photo")
	{
		photo.GET("/getAll", mygramController.GetAllPhoto)
		photo.GET("/getOne", mygramController.GetOnePhoto)
		photo.POST("/createPhoto", mygramController.CreatePhoto)
		photo.PUT("/updatePhoto", mygramController.UpdatePhoto)
		photo.DELETE("/deletePhoto", mygramController.DeletePhoto)
	}

	comment := app.Group("/comment")
	{
		comment.GET("/getAll", mygramController.GetAllComment)
		comment.GET("/getOne", mygramController.GetOneComment)
		comment.POST("/createComment", mygramController.CreateComment)
		comment.PUT("/updateComment", mygramController.UpdateComment)
		comment.DELETE("/deleteComment", mygramController.DeleteComment)
	}

	socialMedia := app.Group("/socialMedia")
	{
		socialMedia.GET("/getAll", mygramController.GetAllSocialMedia)
		socialMedia.GET("/getOne", mygramController.GetOneSocialMedia)
		socialMedia.POST("/createSocialMedia", mygramController.CreateSocialMedia)
		socialMedia.PUT("/updateSocialMedia", mygramController.UpdateSocialMedia)
		socialMedia.DELETE("/deleteSocialMedia", mygramController.DeleteSocialMedia)
	}

	return app
}
