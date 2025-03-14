package storage

import (
	"encoding/json"
	"os"
	. "todo-cli/models"
)

var dbname = "tasks.json"

func SaveToFile(list TaskList) error {
	encodeVal, encodeErr := json.MarshalIndent(list, "", " ")
	if encodeErr != nil {
		return encodeErr
	}

	return os.WriteFile(dbname, encodeVal, 0644)
}

func LoadList() (TaskList, error) {
	data, err := os.ReadFile(dbname)
	// List not found
	if err != nil {
		return TaskList{}, nil
	}

	var list TaskList
	err = json.Unmarshal(data, &list)
	if err != nil {
		return TaskList{}, err
	}
	return list, nil
}
