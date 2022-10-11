package storage

import "upday.com/upday-task-fe/pkg/model"

type BoardStorage interface {
	List() []*model.Board
	Find(boardId string) *model.Board
}

var enBoard = model.Board{Id: "en", Name: "England"}
var deBoard = model.Board{Id: "de", Name: "Deutsch"}
var itBoard = model.Board{Id: "it", Name: "Italiano"}

type boardStorageImpl struct {
}

func (inst *boardStorageImpl) List() []*model.Board {
	return []*model.Board{&enBoard, &deBoard, &itBoard}
}

func (inst *boardStorageImpl) Find(boardId string) *model.Board {
	if boardId == "en" {
		return &enBoard
	} else if boardId == "de" {
		return &deBoard
	} else {
		return nil
	}
}

func NewBoardStorage() BoardStorage {
	return &boardStorageImpl{}
}
