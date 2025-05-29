package main

import (
	"gin-framework/controllers"
	"gin-framework/services"

	// "net/http"
	"fmt"

	internal "gin-framework/internal/database"

	"github.com/gin-gonic/gin"
)

func main(){
	router:=gin.Default()
	db:= internal.InitDB()
	if  db == nil {
		fmt.Println("Failed to connect to the database")
	}
	notesService := &services.NotesService{}
	notesService.InitService(db)

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(router, notesService)

	router.Run(":8080") // listen and serve on
}










	// router.GET("/hello", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Hello, World!",
	// 	})
	// })
	// router.GET("/me/:id/:new_id", func(c *gin.Context) {
	// 	var id=c.Param("id")
	// 	var newId=c.Param("new_id")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Goodbye, World!",
	// 		"check":    "This is a test",
	// 		"user_id": id,
	// 		"new_user_id": newId,
	// 	})
	// })

	// router.POST("/post", func(c *gin.Context) {
	// 	type Post struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	var post Post

	// 	if err :=c.BindJSON(&post);err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Post received",
	// 		"email": post.Email,
	// 		"password": post.Password,
	// 	})
	// })

	// router.PUT("/post", func(c *gin.Context) {
	// 	type Post struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	var post Post

	// 	if err :=c.BindJSON(&post);err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Post received",
	// 		"email": post.Email,
	// 		"password": post.Password,
	// 	})
	// })
	// //totally replaces the data in the database

	// // PATCH is used to update part of the data
	// router.PATCH("/post", func(c *gin.Context) {
	// 	type Post struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	var post Post

	// 	if err :=c.BindJSON(&post);err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Post received",
	// 		"email": post.Email,
	// 		"password": post.Password,
	// 	})
	// })

	// router.DELETE("/post", func(c *gin.Context) {
	// 	type Post struct {
	// 		Email string `json:"email" binding:"required"`
	// 		Password string `json:"password"`
	// 	}

	// 	var post Post

	// 	if err :=c.BindJSON(&post);err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Post received",
	// 		"email": post.Email,
	// 		"password": post.Password,
	// 	})
	// })