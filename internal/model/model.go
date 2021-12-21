package model

type RespondData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Status string `json:"status"`
}
type R struct {
	storeMap map[string]string

}
type RespondStatus struct {
	Status string `json:"status"`
}

type RespondMap struct {
	Pair []RespondData `json:"pair"`
}
type TempDataModelForSave struct {
	Key   string `json:"key"`
	Value string `json:"value"`

}
type TempDataMapForSave struct {
	Pairs []TempDataModelForSave `json:"pair"`
}
