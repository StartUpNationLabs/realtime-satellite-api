package spacetrack

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"path/filepath"
)

func readDataFromFile(filepath string) ([]TLE, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error(err)

		}
	}(file)

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var data []TLE
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// saveDataToFile saves data to a file
func saveDataToFile(filepath2 string, data []TLE) error {

	// create the folder if it doesn't exist
	err := os.MkdirAll(filepath.Dir(filepath2), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(filepath2)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Error(err)
		}
	}(file)

	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
