package services //во всех сервисах пройтись по бд и проверить, заменить

import (
	"context"
	"errors"
	"time"

	// Импортируйте пакеты для работы с базой данных
	"database/sql"
)

// Category - структура, представляющая данные о категории
type Category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// CategoryService - интерфейс сервиса категорий
type CategoryService interface {
	// Методы для работы с категориями:
	GetAllCategories() ([]Category, error)
	GetCategoryByID(id string) (*Category, error)
	CreateCategory(category Category) (*Category, error)
	UpdateCategory(category Category) error
	DeleteCategory(id string) error
}

// CategoryServiceImpl - реализация сервиса категорий
type CategoryServiceImpl struct {
	db *sql.DB // Ссылка на объект базы данных
}

// NewCategoryService - конструктор сервиса категорий
func NewCategoryService(db *sql.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{
		db: db,
	}
}

// GetAllCategories - получение всех категорий
func (s *CategoryServiceImpl) GetAllCategories() ([]Category, error) {
	rows, err := s.db.QueryContext(context.Background(), "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetCategoryByID - получение категории по ID
func (s *CategoryServiceImpl) GetCategoryByID(id string) (*Category, error) {
	var category Category
	err := s.db.QueryRowContext(context.Background(), "SELECT * FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("категория не найдена")
		}
		return nil, err
	}
	return &category, nil
}

// CreateCategory - создание новой категории
func (s *CategoryServiceImpl) CreateCategory(category Category) (*Category, error) {
	ctx := context.Background()
	result, err := s.db.ExecContext(ctx, "INSERT INTO categories (name) VALUES ($1)", category.Name)
	if err != nil {
		return nil, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	category.ID = lastInsertID
	return &category, nil
}

// UpdateCategory - обновление категории
func (s *CategoryServiceImpl) UpdateCategory(category Category) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	return err
}

// DeleteCategory - удаление категории
func (s *CategoryServiceImpl) DeleteCategory(id string) error {
	ctx := context.Background()
	_, err := s.db.ExecContext(ctx, "DELETE FROM categories WHERE id = $1", id)
	return err
}
