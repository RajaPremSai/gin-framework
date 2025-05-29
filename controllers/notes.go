package controllers

import (
	"gin-framework/services"

	"github.com/gin-gonic/gin"
)


type NotesController struct{
    notesService *services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine,notesService *services.NotesService ) {
    n.notesService = notesService 
    notes := router.Group("/notes")
    notes.GET("/", n.GetNotes)
    notes.POST("/", n.CreateNote)
}

func (n *NotesController) GetNotes(c *gin.Context) {
    c.JSON(200, gin.H{
        "Notes": n.notesService.GetNotesService(),
		// "Notes":"Checking",
    })
}

func (n *NotesController) CreateNote(c *gin.Context) {
    c.JSON(200, gin.H{
        "Notes": n.notesService.CreateNoteService(),
		// "sdjfb":"sf",
    })
}