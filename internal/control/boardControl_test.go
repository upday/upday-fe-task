package control

import (
	"reflect"
	"testing"

	"upday.com/upday-task-fe/internal/storage"
	"upday.com/upday-task-fe/pkg/model"
)

func Test_boardControlImpl_List(t *testing.T) {
	type fields struct {
		boardStorage storage.BoardStorage
	}

	tests := []struct {
		name   string
		fields fields
		want   []*model.Board
	}{
		{"should return all boards", fields{&boardStorageMock{}}, []*model.Board{{Id: "en", Name: "England"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &boardControlImpl{
				boardStorage: tt.fields.boardStorage,
			}
			if got := inst.List(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardControlImpl.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_boardControlImpl_BoardNews(t *testing.T) {
	var newsResult = validDraftNews()
	type fields struct {
		boardStorage storage.BoardStorage
		newsStorage  storage.NewsStorage
	}
	type args struct {
		boardId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.BoardNews
		wantErr bool
	}{
		{
			"should return an error if the boardId is invalid",
			fields{&boardStorageMock{empty: true}, &newsStorageMock{}},
			args{"it"}, nil, true,
		},
		{
			"should return a valid BoardNews",
			fields{&boardStorageMock{}, &newsStorageMock{}},
			args{"en"},
			&model.BoardNews{
				Drafts:    []*model.News{&newsResult},
				Published: make([]*model.News, 0),
				Archives:  make([]*model.News, 0),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inst := &boardControlImpl{
				boardStorage: tt.fields.boardStorage,
				newsStorage:  tt.fields.newsStorage,
			}
			got, err := inst.BoardNews(tt.args.boardId)
			if (err != nil) != tt.wantErr {
				t.Errorf("boardControlImpl.BoardNews() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("boardControlImpl.BoardNews() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoardControl(t *testing.T) {
	if NewBoardControl() == nil {
		t.Errorf("NewBoardControl() is nil")
	}
}
