package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthController interface is a contract what this controller can do

type AuthController interface{
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authController struct{
	// disini masukan service yang kalian butuh
	// this is where put yout service
}

// NewAuthController creates a new instance of authcontroller
func NewAuthController() AuthController{
	return &authController{}
}

func (c *authController) Login(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Login",
	})
}
func (c *authController) Register(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello Register",
	})
}