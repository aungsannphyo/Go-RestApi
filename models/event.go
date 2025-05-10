package models

import (
	"time"

	"github.com/aungsannphyo/go-restapi/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

func (event *Event) CreateEvent() error {
	query := `
	INSERT INTO events(name,description,location,dateTime,user_id)
	VALUES(?, ?, ?, ?, ?)`

	statement, err := db.DBInstance.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()
	result, err := statement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	event.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `SELECT * FROM events`

	rows, error := db.DBInstance.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description,
			&event.Location, &event.DateTime, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)

	}

	if len(events) == 0 {
		return []Event{}, nil
	}

	return events, nil
}

func GetEventById(eventId int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`

	row := db.DBInstance.QueryRow(query, eventId)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description,
		&event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) UpdateEvent() error {
	query := `UPDATE events 
	SET
	name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?`

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)

	return err

}

func (event Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DBInstance.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(event.ID)

	return err
}
