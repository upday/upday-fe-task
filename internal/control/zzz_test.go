package control

import (
	"errors"

	"upday.com/upday-task-fe/pkg/model"
)

type boardStorageMock struct {
	empty bool
}
type newsStorageMock struct {
	empty bool
	error bool
	news  *model.News
}

func validDraftNews() model.News {
	return model.News{
		Id:          "id",
		BoardId:     "en",
		Author:      "test@upday.com",
		Title:       "news test title",
		Description: "description",
		ImageURL:    "http://upday.com/image.jpg",
		Status:      model.NewsDraftStatus,
	}
}

func validPublishedNews() model.News {
	var news = validDraftNews()
	news.Status = model.NewsPublishedStatus
	return news
}

func validArchivedNews() model.News {
	var news = validDraftNews()
	news.Status = model.NewsArchivedStatus
	return news
}

var newsWithoutAuthor = model.News{
	Id:          "id",
	BoardId:     "en",
	Title:       "news test title",
	Description: "description",
	ImageURL:    "http://upday.com/image.jpg",
	Status:      model.NewsDraftStatus,
}

func (inst *boardStorageMock) List() []*model.Board {
	if inst.empty {
		return make([]*model.Board, 0)
	}
	return []*model.Board{{Id: "en", Name: "England"}}
}
func (inst *boardStorageMock) Find(boardId string) *model.Board {
	if inst.empty {
		return nil
	}
	return &model.Board{Id: "en", Name: "England"}
}

func (inst *newsStorageMock) List(boardId string) *model.BoardNews {
	var news = validDraftNews()
	return &model.BoardNews{
		Drafts:    []*model.News{&news},
		Published: make([]*model.News, 0),
		Archives:  make([]*model.News, 0),
	}
}

func (inst *newsStorageMock) Save(news *model.News) (*model.News, error) {
	if inst.error {
		return nil, errors.New("internal error")
	}
	news.Id = news.Title + news.BoardId
	return news, nil
}
func (inst *newsStorageMock) Find(newsId string) (*model.News, error) {
	if inst.error {
		return nil, errors.New("to archive error")
	} else if inst.empty {
		return nil, nil
	}

	if inst.news != nil {
		return inst.news, nil
	}

	news := validDraftNews()
	return &news, nil
}
func (inst *newsStorageMock) Delete(news *model.News) error {
	if inst.error {
		return errors.New("delete error")
	} else {
		return nil
	}
}
