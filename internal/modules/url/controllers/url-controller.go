package controllers

import (
	"net/http"
	"strings"

	"github.com/SubhamSaha99/url-shortner/internal/modules/url/dto"
	"github.com/SubhamSaha99/url-shortner/internal/modules/url/services"
	"github.com/gin-gonic/gin"
)

type URLController struct {
	service *services.URLService
	baseURL string
}

func NewURLController(
	service *services.URLService,
	baseURL string,
) *URLController {
	return &URLController{
		service: service,
		baseURL: strings.TrimRight(baseURL, "/"),
	}
}

func (ctrl *URLController) Create(c *gin.Context) {
	var req dto.CreateURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "original_url is required",
		})
		return
	}

	// Controller calls the service here.
	record, err := ctrl.service.CreateShortURL(
		c.Request.Context(),
		req.OriginalURL,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"short_code": record.ShortCode,
		"short_url":  ctrl.baseURL + "/" + record.ShortCode,
	})
}

func (ctrl *URLController) Redirect(c *gin.Context) {
	// Controller asks the service to find the original URL.
	record, err := ctrl.service.GetOriginalURL(
		c.Request.Context(),
		c.Param("code"),
	)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "short URL not found",
		})
		return
	}

	c.Redirect(http.StatusFound, record.OriginalURL)
}

func (ctrl *URLController) DeleteURL(c *gin.Context) {
	var req dto.DeleteURL

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "id is required",
		})
		return
	}

	if err := ctrl.service.DeleteURL(c.Request.Context(), req.Id); err != nil {
		c.JSON(http.StatusExpectationFailed, gin.H{
			"error": "error while deleting url",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "url deleted successfully!",
	})
}
