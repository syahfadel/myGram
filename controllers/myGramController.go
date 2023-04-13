package controllers

import (
	"myGram/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MyGramController struct {
	DB            *gorm.DB
	MyGramService *services.MyGramService
}

func (mc *MyGramController) Register(ctx *gin.Context) {

}

func (mc *MyGramController) Login(ctx *gin.Context) {

}

func (mc *MyGramController) GetAllPhoto(ctx *gin.Context) {

}

func (mc *MyGramController) GetOnePhoto(ctx *gin.Context) {

}

func (mc *MyGramController) CreatePhoto(ctx *gin.Context) {

}

func (mc *MyGramController) UpdatePhoto(ctx *gin.Context) {

}

func (mc *MyGramController) DeletePhoto(ctx *gin.Context) {

}

func (mc *MyGramController) GetAllComment(ctx *gin.Context) {

}

func (mc *MyGramController) GetOneComment(ctx *gin.Context) {

}

func (mc *MyGramController) CreateComment(ctx *gin.Context) {

}

func (mc *MyGramController) UpdateComment(ctx *gin.Context) {

}

func (mc *MyGramController) DeleteComment(ctx *gin.Context) {

}

func (mc *MyGramController) GetAllSocialMedia(ctx *gin.Context) {

}

func (mc *MyGramController) GetOneSocialMedia(ctx *gin.Context) {

}

func (mc *MyGramController) CreateSocialMedia(ctx *gin.Context) {

}

func (mc *MyGramController) UpdateSocialMedia(ctx *gin.Context) {

}

func (mc *MyGramController) DeleteSocialMedia(ctx *gin.Context) {

}
