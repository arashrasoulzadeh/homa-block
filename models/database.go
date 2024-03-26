package models

import (
	"github.com/arashrasoulzadeh/homa-block/util"
)

type Database struct {
	collections []Collection
	Name        string `json:"name"`
}

func (db *Database) AddCollection(collection *Collection) {
	db.collections = append(db.collections, *collection)
}

func (db *Database) GetCollectionAtIndex(index int) *Collection {
	return &db.collections[index]
}

func (db *Database) Save() {
	for index, collection := range db.collections {
		util.Log("saving", index, collection)
		err := util.CreateDirectoryIfNotExists("data/" + db.Name)
		if err != nil {
			panic(err)
		}
		util.SyncCollection("data/"+db.Name+"/"+collection.Name+".json", collection)
	}
}
