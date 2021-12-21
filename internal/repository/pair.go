package repository

import (
	"encoding/json"
	"fmt"
	"sync"
	"yemeksepeti-keyvalue/internal/model"
	"yemeksepeti-keyvalue/pkg/util"
)

type Repository struct {
	pairMap map[string]string
	mu      sync.Mutex
}

func NewRepository() *Repository {
	return &Repository{
		pairMap: make(map[string]string),
	}
}

/*
checks is key using or not
*/
func (r *Repository) IsKeyUsedBefore(k string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	if r.pairMap[k] != "" {
		return true
	}
	return false
}

/*
saves key - value pair to memeory
*/
func (r *Repository) SetPair(k string, v string) bool {
	defer r.mu.Unlock()
	r.mu.Lock()
	r.pairMap[k] = v
	return false
}

/*
returns the value of the key
*/
func (r *Repository) GetValue(k string) string {
	defer r.mu.Unlock()
	r.mu.Lock()
	if !(r.pairMap[k] == "") {
		return r.pairMap[k]
	}
	return ""
}

/*
returns all pairs in the memory
*/
func (r *Repository) GetAll() map[string]string {
	return r.pairMap
}

/*
deletes all pair in the memory and file
*/
func (r *Repository) Flush() error {
	defer r.mu.Unlock()
	r.mu.Lock()
	for key := range r.pairMap {
		delete(r.pairMap, key)
	}
	util.DeleteAllFiles()
	return nil
}

/*
deletes the key-value pair
*/
func (r *Repository) Delete(k string) error {
	defer r.mu.Unlock()
	r.mu.Lock()
	delete(r.pairMap, k)
	return nil
}

/*
saves json file to memory
*/
func (l *Repository) SaveToMemory(data []model.RespondData) error {
	defer l.mu.Unlock()
	l.mu.Lock()
	for i := 0; i < len(data); i++ {
		l.pairMap[data[i].Key] = data[i].Value
	}
	return nil
}

/*
saves memort to json file
*/
func (l *Repository) SaveToFile() error {
	defer l.mu.Unlock()
	l.mu.Lock()
	var temp model.TempDataMapForSave
	for key, val := range l.pairMap {
		temp.Pairs = append(temp.Pairs, model.TempDataModelForSave{Key: key, Value: val})
		fmt.Println(val)
	}
	jsondat := &model.TempDataMapForSave{Pairs: temp.Pairs}
	jsonString, _ := json.Marshal(jsondat)
	util.CreateFile(jsonString)
	return nil
}
