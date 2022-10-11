package control

import (
	"reflect"
	"testing"

	"upday.com/upday-task-fe/internal/storage"
	"upday.com/upday-task-fe/pkg/model"
)

func Test_newsControlImpl_Add(t *testing.T) {
	var validNews = validDraftNews()
	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		news model.News
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.News
		wantErr bool
	}{
		{
			"shold create a news with draft state",
			fields{&boardStorageMock{}, &newsStorageMock{}},
			args{validNews},
			&model.News{
				Id:          validNews.Title + validNews.BoardId,
				BoardId:     validNews.BoardId,
				Author:      validNews.Author,
				Title:       validNews.Title,
				Description: validNews.Description,
				ImageURL:    validNews.ImageURL,
				Status:      model.NewsDraftStatus,
			},
			false,
		},
		{
			"should return an error if the news storage fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{validNews},
			nil,
			true,
		},
		{
			"should return an error if the news's author is blank",
			fields{&boardStorageMock{}, &newsStorageMock{}},
			args{newsWithoutAuthor},
			nil,
			true,
		},
		{
			"should return an error if the board is not founded",
			fields{&boardStorageMock{empty: true}, &newsStorageMock{}},
			args{validNews},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			got, err := inst.Add(tt.args.news)
			if (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newsControlImpl.Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newsControlImpl_ToArchive(t *testing.T) {
	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		newsId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"should change the status", fields{&boardStorageMock{}, &newsStorageMock{}}, args{"news-uuid"}, false},
		{
			"should return an error if the news storage fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{"error-news-uuid"},
			true,
		},
		{
			"should return an error if not found the news",
			fields{&boardStorageMock{}, &newsStorageMock{empty: true}},
			args{"invalid-news-uuid"}, true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			if err := inst.ToArchive(tt.args.newsId); (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.ToArchive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newsControlImpl_ToDraft(t *testing.T) {
	draftNews := validDraftNews()
	archivedNews := validArchivedNews()
	publishedNews := validPublishedNews()

	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		newsId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should move darft to draft",
			fields{&boardStorageMock{}, &newsStorageMock{news: &draftNews}},
			args{"news-id"},
			false,
		},
		{
			"should move published to draft",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{"published-id"},
			false,
		},
		{
			"should not move to draft if state is archived",
			fields{&boardStorageMock{}, &newsStorageMock{news: &archivedNews}},
			args{"archived-id"},
			true,
		},
		{
			"should not move if news storage fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{"archived-id"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			if err := inst.ToDraft(tt.args.newsId); (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.ToDraft() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newsControlImpl_ToPublished(t *testing.T) {
	draftNews := validDraftNews()
	archivedNews := validArchivedNews()
	publishedNews := validPublishedNews()
	invalidDraftNews := validDraftNews()
	invalidDraftNews.Title = ""

	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		newsId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should move darft to published",
			fields{&boardStorageMock{}, &newsStorageMock{news: &draftNews}},
			args{"news-id"},
			false,
		},
		{
			"should move published to published",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{"published-id"},
			false,
		},
		{
			"should not move to published if state is archived",
			fields{&boardStorageMock{}, &newsStorageMock{news: &archivedNews}},
			args{"archived-id"},
			true,
		},
		{
			"should not move to published if the news is invalid",
			fields{&boardStorageMock{}, &newsStorageMock{news: &invalidDraftNews}},
			args{"archived-id"},
			true,
		},
		{
			"should not move if news storage fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{"archived-id"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			if err := inst.ToPublished(tt.args.newsId); (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.ToPublish() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newsControlImpl_Delete(t *testing.T) {
	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		newsId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"should delete a news by id",
			fields{&boardStorageMock{}, &newsStorageMock{}},
			args{"news-id"},
			false,
		},
		{
			"should return an error if the newsId is invalid",
			fields{&boardStorageMock{}, &newsStorageMock{empty: true}},
			args{"invalid-news-id"},
			true,
		},
		{
			"should return an error if the news storage fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{"invalid-news-id"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			if err := inst.Delete(tt.args.newsId); (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_newsControlImpl_Update(t *testing.T) {
	draftNews := validDraftNews()
	draftNews.Title = "updated title"
	publishedNews := validPublishedNews()
	draftNews.Title = "updated title"
	invalidTitleNews := validDraftNews()
	invalidTitleNews.Title = ""
	invalidDescriptionNews := validDraftNews()
	invalidDescriptionNews.Description = ""
	invalidImageUrlNews := validDraftNews()
	invalidImageUrlNews.ImageURL = ""
	invalidAuthorNews := validDraftNews()
	invalidAuthorNews.Author = "email"

	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		news model.News
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.News
		wantErr bool
	}{
		{
			"should update the news",
			fields{&boardStorageMock{}, &newsStorageMock{}},
			args{draftNews},
			&draftNews,
			false,
		},
		{
			"should not update the status",
			fields{&boardStorageMock{}, &newsStorageMock{news: &draftNews}},
			args{publishedNews},
			&draftNews,
			false,
		},
		{
			"should update an invalid news",
			fields{&boardStorageMock{}, &newsStorageMock{news: &draftNews}},
			args{invalidTitleNews},
			&draftNews,
			false,
		},
		{
			"should return an error if update a published news without title",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{invalidTitleNews},
			nil,
			true,
		},
		{
			"should return an error if update a published news without description",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{invalidDescriptionNews},
			nil,
			true,
		},
		{
			"should return an error if update a published news without imageUrl",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{invalidImageUrlNews},
			nil,
			true,
		},
		{
			"should return an error if update a published news with invalid author",
			fields{&boardStorageMock{}, &newsStorageMock{news: &publishedNews}},
			args{invalidAuthorNews},
			nil,
			true,
		},
		{
			"should return an error if news store fails",
			fields{&boardStorageMock{}, &newsStorageMock{error: true}},
			args{draftNews},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &newsControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			got, err := inst.Update(tt.args.news)
			if (err != nil) != tt.wantErr {
				t.Errorf("newsControlImpl.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newsControlImpl.Update() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewNewsControl(t *testing.T) {
	if NewNewsControl() == nil {
		t.Errorf("NewNewsControl() is nil")
	}
}
