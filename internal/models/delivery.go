package models

import (
	"time"
)

type Delivery struct {
	ID           int       `json:"id" db:"id"`
	DeliveryType string    `json:"delivery_type" db:"delivery_type"`
	Price        float64   `json:"price" db:"price"`
	Description  string    `json:"description" db:"description"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

func NewDelivery(deliveryType string, price float64, description string) *Delivery {
	return &Delivery{
		DeliveryType: deliveryType,
		Price:        price,
		Description:  description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (d *Delivery) Update(deliveryType string, price float64, description string) {
	d.DeliveryType = deliveryType
	d.Price = price
	d.Description = description
	d.UpdatedAt = time.Now()
}
