package lib

import "github.com/hashicorp/go-memdb"

var db *memdb.MemDB

func GetDB() (*memdb.MemDB, error) {
	if db != nil {
		return db, nil
	}

	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"news": {
				Name: "news",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Id"},
					},
					"boardId": {
						Name:    "boardId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "BoardId"},
					},
					"author": {
						Name:    "author",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Author"},
					},
					"title": {
						Name:         "title",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Title"},
					},
					"description": {
						Name:         "description",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "Description"},
					},
					"imageURL": {
						Name:         "imageURL",
						Unique:       false,
						AllowMissing: true,
						Indexer:      &memdb.StringFieldIndex{Field: "ImageURL"},
					},
					"status": {
						Name:    "status",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Status"},
					},
					"createdAt": {
						Name:    "createdAt",
						Unique:  false,
						Indexer: &memdb.FieldSetIndex{Field: "CreatedAt"},
					},
				},
			},
		},
	}

	var err error
	db, err = memdb.NewMemDB(schema)
	return db, err
}
