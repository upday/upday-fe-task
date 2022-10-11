package control

import (
	"errors"

	"upday.com/upday-task-fe/internal/storage"
	"upday.com/upday-task-fe/pkg/model"
)

var ErrBoardNotFound = errors.New("id: board not found")

type Boardcontrol interface {
	List() []*model.Board
	BoardNews(boardId string) (*model.BoardNews, error)
}

type boardControlImpl struct {
	boardStorage storage.BoardStorage
	newsStorage  storage.NewsStorage
}

func (inst *boardControlImpl) List() []*model.Board {
	return inst.boardStorage.List()
}

func (inst *boardControlImpl) BoardNews(boardId string) (*model.BoardNews, error) {
	board := inst.boardStorage.Find(boardId)
	if board == nil {
		return nil, ErrBoardNotFound
	}

	return inst.newsStorage.List(boardId), nil
}

func NewBoardControl() Boardcontrol {
	boardStorage := storage.NewBoardStorage()
	newsStorage := storage.NewNewsStorage()
	return &boardControlImpl{boardStorage, newsStorage}
}
