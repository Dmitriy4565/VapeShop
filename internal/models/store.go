package models

import (
	"time"
)

// Store структура, описывающая магазин
type Store struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Address   string    `json:"address" db:"address"`
	Phone     string    `json:"phone" db:"phone"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// NewStore - конструктор магазина
func NewStore(name, address, phone string) *Store {
	return &Store{
		Name:      name,
		Address:   address,
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// Update - обновление данных магазина
func (s *Store) Update(name, address, phone string) {
	s.Name = name
	s.Address = address
	s.Phone = phone
	s.UpdatedAt = time.Now()
}
