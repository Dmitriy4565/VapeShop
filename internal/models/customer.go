package models

import (
	"time"
)

type Customer struct {
	ID        int       `json:"id"`         // Уникальный идентификатор покупателя
	FirstName string    `json:"first_name"` // Имя покупателя
	LastName  string    `json:"last_name"`  // Фамилия покупателя
	Email     string    `json:"email"`      // Электронная почта покупателя
	Phone     string    `json:"phone"`      // Телефон покупателя
	Address   string    `json:"address"`    // Адрес покупателя
	CreatedAt time.Time `json:"created_at"` // Время создания профиля покупателя
	UpdatedAt time.Time `json:"updated_at"` // Время последнего обновления профиля покупателя
}

func NewCustomer(firstName, lastName, email, phone, address string) *Customer {
	return &Customer{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		Address:   address,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Customer) Update(firstName, lastName, email, phone, address string) {
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	c.Phone = phone
	c.Address = address
	c.UpdatedAt = time.Now()
}
