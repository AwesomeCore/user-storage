package us

import (
//	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserStorage struct {
	Collection *mgo.Collection
}

type UserStorageItem struct {
	Key string
	Value string
}

func (US *UserStorage) SetItem(item UserStorageItem, uid string) {
	US.Collection.Upsert(bson.M{"userId": uid}, bson.M{"userId": uid, item.Key: item.Value})
}

func (US *UserStorage) GetItem(key string, uid string) string {
	var data []bson.M
	US.Collection.Find(bson.M{"userId": uid}).All(&data)

	for _, item := range data {
		if val, ok := item[key]; ok {
			return val.(string)
		}
	}
	return ""
}
