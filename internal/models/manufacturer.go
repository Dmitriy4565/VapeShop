package models

import (
	"time"
)

type Manufacturer struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Country   string    `json:"country" db:"country"`
	Website   string    `json:"website" db:"website"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewManufacturer(name, country, website string) *Manufacturer {
	return &Manufacturer{
		Name:      name,
		Country:   country,
		Website:   website,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (m *Manufacturer) Update(name, country, website string) {
	m.Name = name
	m.Country = country
	m.Website = website
	m.UpdatedAt = time.Now()
}
