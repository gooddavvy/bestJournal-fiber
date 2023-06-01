package fs

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gooddavvy/bestJournal-fiber/model"
)

func GetJson() ([]model.JournalPage, error) {
	// Read the JSON file
	file, err := ioutil.ReadFile("json/pages.json")
	if err != nil {
		return nil, err
	}

	// Create a slice to store the pages
	var pages []model.JournalPage

	// Unmarshal the JSON into the slice
	err = json.Unmarshal(file, &pages)
	if err != nil {
		return nil, err
	}

	return pages, nil
}
