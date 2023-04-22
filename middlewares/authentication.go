package middlewares

import (
	"myGram/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Middleware untuk autentikasi menggunakan JWT
// @Description Middleware untuk memverifikasi token JWT dan memeriksa apakah pengguna memiliki hak akses ke endpoint tertentu
// @Tags middleware
// @Security ApiKeyAuth
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {string} string "OK"
// @Failure 401 {string} string "Unauthorized"
// @Failure 500 {string} string "Internal Server Error"
// @Router /middleware/jwt [get]
func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		verifyToken, err := helpers.VerifyToken(ctx)
		_ = verifyToken

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "failed",
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		ctx.Set("userData", verifyToken)
		ctx.Next()
	}
}
