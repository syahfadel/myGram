package controllers

import (
	"fmt"
	"myGram/entities"
	"myGram/helpers"
	"myGram/services"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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

type RequestPhoto struct {
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
}

type RequestComment struct {
	PhotoID uint   `json:"photo_id"`
	Message string `json:"message"`
}

type RequestSocialMedia struct {
	SocialMediaUrl string `json:"social_media_url"`
	Name           string `json:"name"`
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
	if err := ctx.ShouldBindJSON(&requestUser); err != nil {
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
	result, err := mc.MyGramService.GetAllPhoto()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) GetOnePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	photo, err := mc.MyGramService.GetOnePhoto(uint(photoID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   photo,
	})
}

func (mc *MyGramController) CreatePhoto(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestPhoto RequestPhoto
	if err := ctx.ShouldBindJSON(&requestPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	photo := entities.Photo{
		Title:    requestPhoto.Title,
		Caption:  requestPhoto.Caption,
		PhotoUrl: requestPhoto.PhotoUrl,
		UserID:   uint(userId),
	}

	result, err := mc.MyGramService.CreatePhoto(photo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) UpdatePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestPhoto RequestPhoto
	if err := ctx.ShouldBindJSON(&requestPhoto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	photo := entities.Photo{
		ID:       uint(photoID),
		Title:    requestPhoto.Title,
		Caption:  requestPhoto.Caption,
		PhotoUrl: requestPhoto.PhotoUrl,
		UserID:   uint(userId),
	}

	result, err := mc.MyGramService.UpdatePhoto(photo)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) DeletePhoto(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	err = mc.MyGramService.DeletePhoto(uint(photoID), uint(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (mc *MyGramController) GetAllComment(ctx *gin.Context) {
	id := ctx.Param("id")
	photoID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	result, err := mc.MyGramService.GetAllComment(uint(photoID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) GetOneComment(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	comment, err := mc.MyGramService.GetOneComment(uint(commentID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   comment,
	})
}

func (mc *MyGramController) CreateComment(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestComment RequestComment
	if err := ctx.ShouldBindJSON(&requestComment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	comment := entities.Comment{
		PhotoID: requestComment.PhotoID,
		Message: requestComment.Message,
		UserID:  uint(userId),
	}

	result, err := mc.MyGramService.CreateComment(comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestComment RequestComment
	if err := ctx.ShouldBindJSON(&requestComment); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	comment := entities.Comment{
		ID:      uint(commentID),
		Message: requestComment.Message,
		PhotoID: requestComment.PhotoID,
		UserID:  uint(userId),
	}

	result, err := mc.MyGramService.UpdateComment(comment)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	err = mc.MyGramService.DeleteComment(uint(commentID), uint(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (mc *MyGramController) GetAllSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := userData["id"].(float64)

	result, err := mc.MyGramService.GetAllSocialMedia(uint(userID))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) GetOneSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	smID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)

	result, err := mc.MyGramService.GetOneSocialMedia(uint(smID), uint(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) CreateSocialMedia(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestSocialMedia RequestSocialMedia
	if err := ctx.ShouldBindJSON(&requestSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	socialMedia := entities.Socialmedia{
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl,
		Name:           requestSocialMedia.Name,
		UserID:         uint(userId),
	}

	result, err := mc.MyGramService.CreateSocialMedia(socialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) UpdateSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	smID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	var requestSocialMedia RequestSocialMedia
	if err := ctx.ShouldBindJSON(&requestSocialMedia); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	SocialMedia := entities.Socialmedia{
		ID:             uint(smID),
		Name:           requestSocialMedia.Name,
		SocialMediaUrl: requestSocialMedia.SocialMediaUrl,
		UserID:         uint(userId),
	}

	result, err := mc.MyGramService.UpdateSocialMedia(SocialMedia)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   result,
	})
}

func (mc *MyGramController) DeleteSocialMedia(ctx *gin.Context) {
	id := ctx.Param("id")
	smID, err := strconv.Atoi(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":        "failed",
			"error_status":  "wrong parameter",
			"error_message": fmt.Sprintf("%v not an integer", id),
		})
		return
	}

	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userId := userData["id"].(float64)
	err = mc.MyGramService.DeleteSocialMedia(uint(smID), uint(userId))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "failed",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}
