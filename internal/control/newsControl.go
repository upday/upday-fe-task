package control

import (
	"errors"

	"upday.com/upday-task-fe/internal/storage"
	"upday.com/upday-task-fe/pkg/model"
)

var ErrNewsNotFound = errors.New("id: news not found")

type NewsControl interface {
	Add(model.News) (*model.News, error)
	Update(model.News) (*model.News, error)
	Find(newsId string) (*model.News, error)
	Delete(newsId string) error
	ToArchive(newsId string) error
	ToDraft(newsId string) error
	ToPublished(newsId string) error
}

type newsControlImpl struct {
	boardStorage storage.BoardStorage
	newsStorage  storage.NewsStorage
}

func (inst *newsControlImpl) Add(news model.News) (*model.News, error) {
	if inst.boardStorage.Find(news.BoardId) == nil {
		return nil, ErrBoardNotFound
	}

	err := news.ValidateAuthor()
	if err != nil {
		return nil, err
	}

	news.Status = model.NewsDraftStatus
	return inst.newsStorage.Save(&news)
}

func (inst *newsControlImpl) Update(news model.News) (*model.News, error) {
	stored, err := inst.Find(news.Id)
	if err != nil {
		return nil, err
	}

	if err := news.ValidateAuthor(); err != nil {
		return nil, err
	}

	if stored.Status == model.NewsPublishedStatus {
		if err := news.Validate(); err != nil {
			return nil, err
		}
	}

	stored.Author = news.Author
	stored.Title = news.Title
	stored.Description = news.Description
	stored.ImageURL = news.ImageURL

	return stored, nil
}

func (inst *newsControlImpl) Find(newsId string) (*model.News, error) {
	news, err := inst.newsStorage.Find(newsId)
	if err != nil {
		return nil, err
	} else if news == nil {
		return nil, ErrNewsNotFound
	}
	return news, nil
}

func (inst *newsControlImpl) Delete(newsId string) error {
	news, err := inst.Find(newsId)
	if err != nil {
		return err
	}
	return inst.newsStorage.Delete(news)
}

func (inst *newsControlImpl) ToArchive(newsId string) error {
	news, err := inst.Find(newsId)
	if err != nil {
		return err
	}

	news.Status = model.NewsArchivedStatus
	return nil
}

func (inst *newsControlImpl) ToDraft(newsId string) error {
	news, err := inst.Find(newsId)
	if err != nil {
		return err
	}
	if err = news.ValidateNewStatus(model.NewsDraftStatus); err != nil {
		return err
	}

	news.Status = model.NewsDraftStatus
	return nil
}

func (inst *newsControlImpl) ToPublished(newsId string) error {
	news, err := inst.Find(newsId)
	if err != nil {
		return err
	}
	if err = news.ValidateNewStatus(model.NewsPublishedStatus); err != nil {
		return err
	}

	news.Status = model.NewsPublishedStatus
	return nil
}

func NewNewsControl() NewsControl {
	boardStorage := storage.NewBoardStorage()
	newsStorage := storage.NewNewsStorage()
	return &newsControlImpl{boardStorage, newsStorage}
}
