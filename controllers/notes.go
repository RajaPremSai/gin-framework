package controllers

import (
	"gin-framework/services"
	"strconv"

	"github.com/gin-gonic/gin"
)


type NotesController struct{
    notesService *services.NotesService
}

func (n *NotesController) InitController(notesService *services.NotesService ) *NotesController{
	n.notesService=notesService
	return n
}

func (n *NotesController) InitRoutes(router *gin.Engine ) {
    // n.notesService = notesService 
    notes := router.Group("/notes")
    notes.GET("/", n.GetNotes())
    notes.POST("/", n.CreateNote())
	notes.PUT("/", n.UpdateNotes())
	notes.DELETE("/:id",n.DeleteNote())
	notes.GET("/:id",n.GetNoteById())
}


func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {

		status := c.Query("status")
		actualStatus,err :=strconv.ParseBool(status)

		if err!=nil{
			c.JSON(400,gin.H{
				"message": "Invalid status value, must be true or false",
				"error": err.Error(),
			})
			return
		}

		notes,err :=n.notesService.GetNotesService(actualStatus)
		if err!=nil{
			c.JSON(400, gin.H{"message": err.Error(),})
			return
		}
		c.JSON(200,gin.H{
			"notes":notes,
			"message": "Notes retrieved successfully",
		})
	}
}

func (n *NotesController) CreateNote() gin.HandlerFunc {
	type NoteBody struct{
		Title string `json:"title" binding:"required`
		Status bool `json:"status"`
	}

    return func(c *gin.Context){
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{"message": err.Error(),})
			return
		}
		note, err := n.notesService.CreateNoteService(noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to create note",})
			return
		}
		c.JSON(201, gin.H{
			"message": "Note created successfully",
			"note": note,
		})
		// c.JSON(200, gin.H{
        // 	"notes": n.notesService.CreateNoteService(),
    	// })
	}
}

func (n *NotesController) UpdateNotes() gin.HandlerFunc {
	type NoteBody struct{
		Title string `json:"title"`
		Status bool `json:"status"`
		Id int `json:"id" binding:"required"`
	}

    return func(c *gin.Context){
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{"message": err.Error(),})
			return
		}
		note, err := n.notesService.UpdateNoteService(noteBody.Id,noteBody.Title, noteBody.Status)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to update note",})
			return
		}
		c.JSON(201, gin.H{
			"message": "Note updated successfully",
			"note": note,
		})
	}
}

func (n *NotesController) DeleteNote() gin.HandlerFunc {
	type NoteBody struct{
		Title string `json:"title"`
		Status bool `json:"status"`
		Id int `json:"id" binding:"required"`
	}

    return func(c *gin.Context){
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{"message": err.Error(),})
			return
		}
		// Extract the ID from the URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid ID format",})
			return
		}
		// Set the ID in the noteBody struct
		noteBody.Id = id
		// Call the service to delete the note
		note, err := n.notesService.DeleteNoteService(noteBody.Id)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to delete note",})
			return
		}
		c.JSON(201, gin.H{
			"message": "Note deleted successfully",
			"note": note,
		})
	}
}

func (n *NotesController) GetNoteById() gin.HandlerFunc {
	type NoteBody struct{
		Title string `json:"title"`
		Status bool `json:"status"`
		Id int `json:"id" binding:"required"`
	}

    return func(c *gin.Context){
		var noteBody NoteBody
		if err := c.BindJSON(&noteBody); err != nil {
			c.JSON(400, gin.H{"message": err.Error(),})
			return
		}
		// Extract the ID from the URL parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid ID format",})
			return
		}
		// Set the ID in the noteBody struct
		noteBody.Id = id
		// Call the service to delete the note
		note, err := n.notesService.GetNoteByIdService(noteBody.Id)
		if err != nil {
			c.JSON(500, gin.H{"message": "Failed to fetch note",})
			return
		}
		c.JSON(201, gin.H{
			"message": "Note deleted successfully",
			"note": note,
		})
	}
}