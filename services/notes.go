package services

import (
	"errors"
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

func (n *NotesService) GetNotesService(status bool) ([]*internal.Notes,error) {
	var notes []*internal.Notes

	if err:=n.db.Where("status= ?",status).Find(&notes).Error; err !=nil{
		return nil,err;
	}
	return notes,nil;

}

func (n *NotesService) CreateNoteService(title string,status bool) (*internal.Notes, error) {

	note := &internal.Notes{
		Title:title,
		Status:status,
	}

	if note.Title==""{
		return nil,errors.New("Title cannot be empty");
	}

	err := n.db.Create(note).Error;

	if err!=nil{
		fmt.Println(err);
		return nil, err
	}
	return note,nil
}


func (n *NotesService) UpdateNoteService(id int,title string,status bool) (*internal.Notes, error) {

	var note *internal.Notes
	if err := n.db.Where("id = ?",id).First(&note).Error; err != nil {
		return nil, err;
	}
	note.Title = title
	note.Status = status


	if note.Title==""{
		return nil,errors.New("title cannot be empty");
	}

	err := n.db.Save(note).Error;

	if err!=nil{
		fmt.Println(err);
		return nil, err
	}
	return note,nil
}

func (n *NotesService) DeleteNoteService(id int) (*internal.Notes, error) {

	var note *internal.Notes

	if err := n.db.Where("id = ?",id).First(&note).Error; err != nil {
		return nil, err;
	}
	if note == nil {
		return nil, errors.New("note not found");
	}
	err := n.db.Delete(note).Error;

	if err!=nil{
		fmt.Println(err);
		return nil, err
	}
	return note,nil
}

func (n *NotesService) GetNoteByIdService(id int) (*internal.Notes, error) {

	var note *internal.Notes

	if err := n.db.Where("id = ?",id).First(&note).Error; err != nil {
		return nil, err;
	}
	if note == nil {
		return nil, errors.New("note not found");
	}

	return note,nil
}