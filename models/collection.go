package models

import "github.com/google/uuid"

type Collection struct {
	Name string                    `json:"name"`
	Hash uuid.UUID                 `json:"hash"`
	Data map[string]CollectionData `json:"data"`
}
type CollectionData struct {
	Hash uuid.UUID `json:"hash"`
}

func (c *Collection) Insert(key string, data CollectionData) {
	if c.Data == nil {
		c.Data = make(map[string]CollectionData)
	}
	c.Data[key] = data
}
