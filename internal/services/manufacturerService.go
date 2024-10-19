package services

import (
	"context"
	"errors"
	"time"

	// Импортируйте пакеты для работы с базой данных
	"database/sql"
)

// Manufacturer - структура, представляющая данные о производителе
type Manufacturer struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ManufacturerService - интерфейс сервиса производителей
type ManufacturerService interface {
	// Методы для работы с производителями:
	GetAllManufacturers() ([]Manufacturer, error)
	GetManufacturerByID(id string) (*Manufacturer, error)
	CreateManufacturer(manufacturer Manufacturer) (*Manufacturer, error)
	UpdateManufacturer(manufacturer Manufacturer) error
	DeleteManufacturer(id string) error
}

// ManufacturerServiceImpl - реализация сервиса производителей
type ManufacturerServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

// NewManufacturerService - конструктор сервиса производителей
func NewManufacturerService(db *sql.DB) *ManufacturerServiceImpl {
	return &ManufacturerServiceImpl{
		db: db,
	}
}

// GetAllManufacturers - получение всех производителей
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

// GetManufacturerByID - получение производителя по ID
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

// CreateManufacturer - создание нового производителя
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

// UpdateManufacturer - обновление производителя
func (s *ManufacturerServiceImpl) UpdateManufacturer(manufacturer Manufacturer) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE manufacturers SET name = $1, country = $2 WHERE id = $3", manufacturer.Name, manufacturer.Country, manufacturer.ID)
	return err
}

// DeleteManufacturer - удаление производителя
func (s *ManufacturerServiceImpl) DeleteManufacturer(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM manufacturers WHERE id = $1", id)
	return err
}
