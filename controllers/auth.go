package controllers

import (
	"gin-framework/services"
	"gin-framework/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{
	authService *services.AuthService
}

func (a *AuthController) InitAuthController(authService *services.AuthService)*AuthController{
	a.authService=authService
	return a
}

func (a *AuthController)InitRoutes(router *gin.Engine){
	auth :=router.Group("/auth")
	auth.POST("/register",a.Register())
	auth.POST("/login",a.Login())
	// auth.POST("/logout",a.Logout())
}

func (a *AuthController) Register() gin.HandlerFunc {
	return func(c *gin.Context){
		var user struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		newUser, err := a.authService.RegisterUser(user.Email, user.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to register user", "details": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "User registered successfully", "user": newUser})
	}
}

func (a *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": "Invalid input"})
			return
		}

		existingUser, err := a.authService.LoginUser(user.Email, user.Password)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid email or password", "details": err.Error()})
			return
		}

		token,err :=  utils.GenerateToken(existingUser.Email,existingUser.Id)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to generate token", "details": err.Error()})
			return
		}
		c.Header("Authorization", "Bearer "+token)

		c.JSON(200, gin.H{"message": "Login successful", "user": existingUser,"token": token})
	}
}