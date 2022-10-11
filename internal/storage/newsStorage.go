package storage

import (
	"sort"
	"time"

	"github.com/hashicorp/go-uuid"
	"upday.com/upday-task-fe/internal/lib"
	"upday.com/upday-task-fe/pkg/model"
)

type NewsStorage interface {
	List(boardId string) *model.BoardNews
	Save(news *model.News) (*model.News, error)
	Find(newsId string) (*model.News, error)
	Delete(news *model.News) error
}

type newsStorageImpl struct {
}

func (inst *newsStorageImpl) List(boardId string) *model.BoardNews {
	db, _ := lib.GetDB()
	tnx := db.Txn(false)
	defer tnx.Abort()

	it, _ := tnx.Get("news", "boardId", boardId)

	published := make([]*model.News, 0)
	archives := make([]*model.News, 0)
	drafts := make([]*model.News, 0)
	for obj := it.Next(); obj != nil; obj = it.Next() {
		news := obj.(*model.News)
		switch news.Status {
		case model.NewsArchivedStatus:
			archives = append(archives, obj.(*model.News))
		case model.NewsDraftStatus:
			drafts = append(drafts, obj.(*model.News))
		default:
			published = append(published, obj.(*model.News))
		}

	}

	sort.Sort(model.ByNews(published))
	sort.Sort(model.ByNews(archives))
	sort.Sort(model.ByNews(drafts))

	return &model.BoardNews{
		Archives:  archives,
		Drafts:    drafts,
		Published: published,
	}
}

func (inst *newsStorageImpl) Save(news *model.News) (*model.News, error) {
	db, err := lib.GetDB()
	if err != nil {
		return nil, err
	}

	id, err := uuid.GenerateUUID()
	if err != nil {
		return nil, err
	}

	news.Id = id
	news.CreatedAt = time.Now()

	txn := db.Txn(true)
	defer txn.Abort()
	txn.Insert("news", news)
	txn.Commit()

	return news, nil
}

func (inst *newsStorageImpl) Find(newsId string) (*model.News, error) {
	db, _ := lib.GetDB()

	txn := db.Txn(false)
	defer txn.Abort()
	raw, err := txn.First("news", "id", newsId)
	if err != nil || raw == nil {
		return nil, err
	}

	return raw.(*model.News), nil
}

func (inst *newsStorageImpl) Delete(news *model.News) error {
	db, _ := lib.GetDB()

	txn := db.Txn(true)
	defer txn.Abort()
	err := txn.Delete("news", news)
	if err != nil {
		return err
	}
	txn.Commit()

	return nil
}

func NewNewsStorage() NewsStorage {
	return &newsStorageImpl{}
}
