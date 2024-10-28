package services

import (
	"context"
	"errors"
	"time"

	"database/sql"
)

type Manufacturer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type ManufacturerService interface {
	GetAllManufacturers() ([]Manufacturer, error)
	GetManufacturerByID(id string) (*Manufacturer, error)
	CreateManufacturer(manufacturer Manufacturer) (*Manufacturer, error)
	UpdateManufacturer(manufacturer Manufacturer) error
	DeleteManufacturer(id string) error
}

type ManufacturerServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewManufacturerService(db *sql.DB) *ManufacturerServiceImpl {
	return &ManufacturerServiceImpl{
		db: db,
	}
}

func (s *ManufacturerServiceImpl) GetAllManufacturers() ([]Manufacturer, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM manufacturers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var manufacturers []Manufacturer
	for rows.Next() {
		var manufacturer Manufacturer
		if err := rows.Scan(&manufacturer.ID, &manufacturer.Name, &manufacturer.Country, &manufacturer.CreatedAt, &manufacturer.UpdatedAt); err != nil {
			return nil, err
		}
		manufacturers = append(manufacturers, manufacturer)
	}

	return manufacturers, nil
}

func (s *ManufacturerServiceImpl) GetManufacturerByID(id string) (*Manufacturer, error) {
	var manufacturer Manufacturer
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM manufacturers WHERE id = $1", id).Scan(&manufacturer.ID, &manufacturer.Name, &manufacturer.Country, &manufacturer.CreatedAt, &manufacturer.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("производитель не найден")
		}
		return nil, err
	}
	return &manufacturer, nil
}

func (s *ManufacturerServiceImpl) CreateManufacturer(manufacturer Manufacturer) (*Manufacturer, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO manufacturers (name, country) VALUES ($1, $2)", manufacturer.Name, manufacturer.Country)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	manufacturer.ID = lastInsertID
	return &manufacturer, nil
}

func (s *ManufacturerServiceImpl) UpdateManufacturer(manufacturer Manufacturer) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE manufacturers SET name = $1, country = $2 WHERE id = $3", manufacturer.Name, manufacturer.Country, manufacturer.ID)
	return err
}

func (s *ManufacturerServiceImpl) DeleteManufacturer(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM manufacturers WHERE id = $1", id)
	return err
}
