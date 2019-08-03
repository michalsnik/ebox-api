package boxes

import (
	"ebox-api/internal/db"
)

type Service struct {
	db *db.DB
}

func NewService(db *db.DB) *Service {
	return &Service{db: db}
}

func (svc *Service) GetBoxById (boxID int) Box {
	return Box{Id: boxID, Name: "Lorem ipsum"}
}

func (svc *Service) CreateBox (payload CreateBoxRequest) (*Box, error) {
	query := `
		INSERT INTO ebox.boxes (name)
		VALUES ($1)
		RETURNING id
	`

	id := 0
	err := svc.db.QueryRow(query, payload.Name).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &Box{Id: id, Name: payload.Name }, nil
}
