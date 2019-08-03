package boxes

import "database/sql"

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

func (svc *Service) GetBoxById (boxID int) Box {
	return Box{Id: boxID, Name: "Lorem ipsum"}
}

func (svc *Service) CreateBox (payload CreateBoxRequest) (Box, error) {
	return Box{Id: 12, Name: payload.Name }, nil
}
