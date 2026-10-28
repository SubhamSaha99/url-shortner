package dto

type CreateURLRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

type DeleteURL struct {
	Id uint `json:"id" binding:"required"`
}
