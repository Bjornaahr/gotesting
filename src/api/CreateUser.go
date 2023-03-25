package api

import (
	"VekterBackend/src/controllers"
	"VekterBackend/src/initializers"
	"VekterBackend/src/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func CreateUser(context *gin.Context) {

	var createUserStruct struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
		//DoB      time.Time `form:"DoB" binding:"required"`
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(createUserStruct.Password), 10)
	if err != nil {
		context.AbortWithStatus(http.StatusInternalServerError)
	}

	if err := context.ShouldBind(&createUserStruct); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.Users{DoB: time.Now(), Email: createUserStruct.Email, Password: string(hash), Name: createUserStruct.Name}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		context.JSONP(400, gin.H{
			"message": "User already created",
		})
		return
	}
	context.JSONP(200, gin.H{
		"message": "pong",
	})
}

func Login(context *gin.Context) {

	var loginForm struct {
		Email    string `form:"email" binding:"required"`
		Password string `form:"password" binding:"required"`
	}

	if err := context.ShouldBind(&loginForm); err != nil {
		fmt.Printf("BindFailed")
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.Users

	err := initializers.DB.Where("email = ?", loginForm.Email).First(&user).Error
	if err != nil {
		fmt.Printf("DB failed")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if user.CheckPassword(loginForm.Password) {
		fmt.Printf("CheckPasswordFailed")
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := controllers.GenerateToken(loginForm.Email)
	if err != nil {
		fmt.Printf("TokenFailed")
		context.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": token})
}

func Pong(context *gin.Context) {

	context.JSONP(200, gin.H{
		"message": "pong",
	})
}
