
package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/SubhamSaha99/url-shortner/internal/modules/url/models"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

// Create stores a new shortened URL.
func (r *URLRepository) Create(
	ctx context.Context,
	url *models.URL,
) error {
	if err := r.db.WithContext(ctx).Create(url).Error; err != nil {
		return fmt.Errorf("create URL: %w", err)
	}

	return nil
}

// FindByShortCode retrieves a URL using its short code.
func (r *URLRepository) FindByShortCode(
	ctx context.Context,
	shortCode string,
) (*models.URL, error) {
	var url models.URL

	err := r.db.WithContext(ctx).
		Where("short_code = ?", shortCode).
		First(&url).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("find URL by short code: %w", err)
	}

	return &url, nil
}

// FindByID retrieves a URL using its database ID.
func (r *URLRepository) FindByID(
	ctx context.Context,
	id uint,
) (*models.URL, error) {
	var url models.URL

	err := r.db.WithContext(ctx).First(&url, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("find URL by ID: %w", err)
	}

	return &url, nil
}

// Delete removes a URL using its database ID.
func (r *URLRepository) Delete(
	ctx context.Context,
	id uint,
) error {
	result := r.db.WithContext(ctx).Delete(&models.URL{}, id)

	if result.Error != nil {
		return fmt.Errorf("delete URL: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
