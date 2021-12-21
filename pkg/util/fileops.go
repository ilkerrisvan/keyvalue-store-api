package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"
	"yemeksepeti-keyvalue/internal/model"
)

const basePath = "pkg/temp/"

func DeleteAllFiles() {
	dir, _ := ioutil.ReadDir(basePath)
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{basePath, d.Name()}...))
	}
}
func CreateFile(data []byte) {
	filename := fmt.Sprintf("pkg/temp/%v-data.json", time.Now().Unix())
	os.WriteFile(filename, data, 0644)

}
func IsThereAnyFile() bool {
	dir, _ := ioutil.ReadDir(basePath)
	if len(dir) > 0 {
		return true
	}
	return false
}
func OpenFileAndReturnDataFromFile() []model.RespondData {
	dir, _ := ioutil.ReadDir(basePath)
	name := dir[len(dir)-1].Name()
	filepath := basePath + name
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload model.RespondMap
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return payload.Pair
}
