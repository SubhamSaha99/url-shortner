package services

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"net/url"
	"strings"

	"github.com/SubhamSaha99/url-shortner/internal/modules/url/models"
	"gorm.io/gorm"
)

type URLRepository interface {
	Create(ctx context.Context, url *models.URL) error
	FindByShortCode(ctx context.Context, shortCode string) (*models.URL, error)
	Delete(ctx context.Context, id uint) error
}

type URLService struct {
	repo    URLRepository
	charset string
}

func NewURLService(repo URLRepository, charset string) *URLService {
	return &URLService{
		repo:    repo,
		charset: charset,
	}
}

func generateShortCode(length int, charset string) (string, error) {
	var result strings.Builder

	for range length {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("generate short code: %w", err)
		}

		result.WriteByte(charset[n.Int64()])
	}

	return result.String(), nil
}

func validateURL(rawURL string) error {
	parsed, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return errors.New("invalid URL format")
	}

	if (parsed.Scheme != "http" && parsed.Scheme != "https") ||
		parsed.Hostname() == "" {
		return errors.New("URL must start with http:// or https://")
	}

	return nil
}

func (s *URLService) CreateShortURL(
	ctx context.Context,
	originalURL string,
) (*models.URL, error) {
	originalURL = strings.TrimSpace(originalURL)

	if err := validateURL(originalURL); err != nil {
		return nil, err
	}

	for range 5 {
		code, err := generateShortCode(8, s.charset)
		if err != nil {
			return nil, err
		}

		record := &models.URL{
			OriginalURL: originalURL,
			ShortCode:   code,
		}

		if err := s.repo.Create(ctx, record); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				continue
			}

			return nil, fmt.Errorf("save shortened URL: %w", err)
		}

		return record, nil
	}

	return nil, errors.New("could not generate a short code")
}

func (s *URLService) GetOriginalURL(
	ctx context.Context,
	shortCode string,
) (*models.URL, error) {
	if strings.TrimSpace(shortCode) == "" {
		return nil, errors.New("short code is required")
	}

	record, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, fmt.Errorf("retrieve shortened URL: %w", err)
	}

	return record, nil
}

func (s *URLService) DeleteURL(
	ctx context.Context,
	id uint,
) error {
	if id == 0 {
		return errors.New("id should be greater than 0")
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return fmt.Errorf("error while deleting: %w", err)
	}
	return nil
}
