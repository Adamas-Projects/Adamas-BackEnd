package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type EventService struct {
	EventDB *database.EventDB
}

func NewEventService(eventDB *database.EventDB) *EventService {
	return &EventService{
		EventDB: eventDB,
	}
}

func (es *EventService) CreateEvent(name, address, date, description string, institutionID int) (*entity.Event, error) {
	event, err := es.EventDB.CreateEvent(name, address, date, description, institutionID)
	if err != nil {
		return nil, err 
	}
	return event,nil
}