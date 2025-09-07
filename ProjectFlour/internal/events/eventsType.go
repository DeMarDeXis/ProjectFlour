package events

import (
	"errors"
)

const (
	EventFileImported = "file_imported"
	//eventDataUpdated  = "data_updated"
)

type Event interface {
	Validate() error
}

type FileImportedEvent struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Count   int    `json:"count"`
}

//type DataUpdatedEvent struct {
//	Entity  string `json:"entity"`
//	Action  string `json:"action"`
//	Success bool   `json:"success"`
//}

func (e FileImportedEvent) Validate() error {
	if e.Type == "" || e.Message == "" {
		return errors.New("invalid file imported event")
	}

	return nil
}
