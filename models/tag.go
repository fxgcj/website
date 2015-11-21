package models

import (
	"gopkg.in/mgo.v2/bson"
	"sort"
)

type Tag struct {
	ID      bson.ObjectId   `bson:"_id"`
	Name    string          `bson:"name"`
	BlogIDs []bson.ObjectId `bson:"blogs"`
}

type Tags []*Tag

func (t Tags) Len() int {
	return len(t)
}

func (t Tags) Swap(a, b int) {
	t[a], t[b] = t[b], t[a]
}

func (t Tags) Less(a, b int) bool {
	if len(t[a].BlogIDs) < len(t[b].BlogIDs) {
		return false
	}
	return true
}

func (t Tags) Sort() {
	sort.Sort(t)
}
