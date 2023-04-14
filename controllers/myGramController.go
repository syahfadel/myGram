package controllers

import (
	"myGram/entities"
	"myGram/helpers"
	"myGram/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type MyGramController struct {
	DB            *gorm.DB
	MyGramService *services.MyGramService
}

type RequestUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
}

func (mc *MyGramController) Register(ctx *gin.Context) {
	var requestUser RequestUser
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user := entities.User{
		Username: requestUser.Username,
		Email:    requestUser.Email,
		Password: requestUser.Password,
		Age:      requestUser.Age,
	}

	result, err := mc.MyGramService.CreateUser(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"result": result,
	})

}

func (mc *MyGramController) Login(ctx *gin.Context) {
	var requestUser RequestUser
	if err := ctx.ShouldBind(&requestUser); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	user := entities.User{
		Email: requestUser.Email,
	}

	result, err := mc.MyGramService.Login(user)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	success := helpers.ComparePass([]byte(result.Password), []byte(requestUser.Password))
	if !success {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"status":  "failed",
			"error":   "Unauthorized",
			"message": "invalid email/password",
		})
		return
	}

	token := helpers.GenerateToken(result.ID, result.Email, result.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

func (mc *MyGramController) GetAllPhoto(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) GetOnePhoto(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) CreatePhoto(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) UpdatePhoto(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) DeletePhoto(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) GetAllComment(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) GetOneComment(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) CreateComment(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) UpdateComment(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) DeleteComment(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) GetAllSocialMedia(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) GetOneSocialMedia(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) CreateSocialMedia(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) UpdateSocialMedia(ctx *gin.Context) {
	// TO DO
}

func (mc *MyGramController) DeleteSocialMedia(ctx *gin.Context) {
	// TO DO
}
