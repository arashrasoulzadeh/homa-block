package main

import (
	"fmt"

	"github.com/arashrasoulzadeh/homa-block/models"
	"github.com/arashrasoulzadeh/homa-block/util"
	"github.com/google/uuid"
)

func main() {

	databases_dirs := util.ListDirectories("data")
	var databases = make([]models.Database, len(databases_dirs))
	for index, name := range databases_dirs {
		databases[index] = models.Database{
			Name: name,
		}
		ReadCollections(&databases[index])
	}

	fmt.Printf("%+v\n", databases)
	// databases[0].GetCollectionAtIndex(0).Insert("test", models.CollectionData{
	// 	Hash: uuid.New(),
	// })
	// databases[0].GetCollectionAtIndex(0).Insert("test2", models.CollectionData{
	// 	Hash: uuid.New(),
	// })

	databases[0].Save()
}

func ReadCollections(database *models.Database) {
	collections_dirs := util.ListFiles("data/" + database.Name)
	var collections = make([]models.Collection, len(collections_dirs))
	for index, name := range collections_dirs {
		file, err := util.ReadFile("data/" + database.Name + "/" + name)
		if err != nil {
			panic(err)
		}
		dataMap := file.(map[string]interface{})
		collcetionUuid, err := uuid.Parse(dataMap["hash"].(string))
		if err != nil {
			panic(err)
		}
		collections[index].Hash = collcetionUuid
		collections[index].Name = dataMap["name"].(string)
		collections[index].Data = make(map[string]models.CollectionData)
		// collections[index].Data = dataMap["data"].(map[string]models.CollectionData)
		for k, d := range dataMap["data"].(map[string]interface{}) {
			fmt.Println(k, d)
			cd := d.(map[string]interface{})
			collcetionDataUuid, err := uuid.Parse(cd["hash"].(string))
			if err != nil {
				panic(err)
			}
			collections[index].Data[k] = models.CollectionData{
				Hash: collcetionDataUuid,
			}
		}
		database.AddCollection(&collections[index])
	}

}
