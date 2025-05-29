package services

import (
	"fmt"
	internal "gin-framework/internal/model"

	"gorm.io/gorm"
)

type NotesService struct{
	db *gorm.DB
}

func (n *NotesService) InitService(database *gorm.DB){
	n.db=database;
	n.db.AutoMigrate(&internal.Notes{})
}

type Notes struct {
	Id   int
	Name string
}

func (n *NotesService) GetNotesService() []Notes {
	data := []Notes{
		{
			Id:   1,
			Name: "Note 1",
		},
		{
			Id:   2,
			Name: "Note 2",
		},
	}
	return data
}

func (n *NotesService) CreateNoteService() Notes {
	data := Notes{
		Id:3,
		Name:"Note 3", 
	}

	err := n.db.Create(&internal.Notes{
		Id: 1,
		Title:"Notes",
		Status:true,

	});

	if err!=nil{
		fmt.Println(err);
	}
	fmt.Println("Note created successfully")
	return data
}