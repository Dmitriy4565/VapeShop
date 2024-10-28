package services //заменить s.db на фактическое подключение к бд, но это в конце после маина

import (
	"context"
	"errors"
	"time"

	"database/sql"
)

type Store struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type StoreService interface {
	GetAllStores() ([]Store, error)
	GetStoreByID(id string) (*Store, error)
	CreateStore(store Store) (*Store, error)
	UpdateStore(store Store) error
	DeleteStore(id string) error
}

type StoreServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

func NewStoreService(db *sql.DB) *StoreServiceImpl {
	return &StoreServiceImpl{
		db: db,
	}
}

func (s *StoreServiceImpl) GetAllStores() ([]Store, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM stores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stores []Store
	for rows.Next() {
		var store Store
		if err := rows.Scan(&store.ID, &store.Name, &store.Address, &store.CreatedAt, &store.UpdatedAt); err != nil {
			return nil, err
		}
		stores = append(stores, store)
	}

	return stores, nil
}

func (s *StoreServiceImpl) GetStoreByID(id string) (*Store, error) {
	var store Store
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM stores WHERE id = $1", id).Scan(&store.ID, &store.Name, &store.Address, &store.CreatedAt, &store.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("магазин не найден")
		}
		return nil, err
	}
	return &store, nil
}

func (s *StoreServiceImpl) CreateStore(store Store) (*Store, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO stores (name, address) VALUES ($1, $2)", store.Name, store.Address)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	store.ID = lastInsertID
	return &store, nil
}

func (s *StoreServiceImpl) UpdateStore(store Store) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE stores SET name = $1, address = $2 WHERE id = $3", store.Name, store.Address, store.ID)
	return err
}

func (s *StoreServiceImpl) DeleteStore(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM stores WHERE id = $1", id)
	return err
}
