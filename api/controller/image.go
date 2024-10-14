package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-server/domain"
	"go-server/internal/logger"
	"go.uber.org/zap"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type ImageController struct {
	ImageUsecase domain.ImageUsecase
}

const imagePath = "./uploads/"

// GetImage
// @Summary Получить изображение по ID.
// @Description Возвращает изображение по ID.
// @Tags Image
// @Produce image/png
// @Produce image/jpeg
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Param id path int true "Image ID"
// @Security Bearer Authentication
// @Success 200 {file} file "Image received"
// @Failure 400 {object} domain.Error
// @Failure 404 {object} domain.Error
// @Router /image/{id} [get]
func (ic *ImageController) GetImage(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: "ID is invalid"})
		return
	}

	r, err := ic.ImageUsecase.GetByID(c, id)
	if err != nil {
		logger.Error("image not found", zap.Error(err))
		c.JSON(http.StatusNotFound, domain.Error{Message: err.Error()})
		return
	}

	c.File(imagePath + r.Filename)
}

// UploadImage
// @Summary Загрузить изображение.
// @Description Загружает выбранное изображение.
// @Tags Image
// @Security Bearer Authentication
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Image file"
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Success 200 {string} string "Image ID"
// @Failure 400 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /image/ [post]
func (ic *ImageController) UploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("failed to get image", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: err.Error()})
		return
	}

	mimeType := file.Header.Get("Content-Type")
	if mimeType != "image/png" && mimeType != "image/jpeg" {
		logger.Error("wrong file type", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: "Only image files (png, jpeg) are allowed"})
		return
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(imagePath, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		logger.Error("failed to save image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	image := domain.Image{
		Filename: filename,
	}

	r, err := ic.ImageUsecase.Create(c, &image)
	if err != nil {
		logger.Error("failed to create image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, r)
}

// DeleteImageById
// @Summary Удалить изображение по ID.
// @Description Удаляет изображение по ID.
// @Tags Image
// @Security Bearer Authentication
// @Accept mpfd
// @Produce json
// @Param id path int true "Image ID"
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Success 200 {string} string "Image deleted successfully"
// @Failure 400 {object} domain.Error
// @Failure 404 {object} domain.Error
// @Failure 500 {object} domain.Error
// @Router /image/{id} [delete]
func (ic *ImageController) DeleteImageById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		c.JSON(http.StatusBadRequest, domain.Error{Message: err.Error()})
		return
	}

	image, err := ic.ImageUsecase.GetByID(c, id)
	if err != nil {
		logger.Error("image not found", zap.Error(err))
		c.JSON(http.StatusNotFound, domain.Error{Message: err.Error()})
		return
	}

	filePath := filepath.Join(imagePath, image.Filename)
	if err = os.Remove(filePath); err != nil {
		logger.Error("failed to delete image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	if err = ic.ImageUsecase.DeleteByID(c, id); err != nil {
		logger.Error("failed to delete image", zap.Error(err))
		c.JSON(http.StatusInternalServerError, domain.Error{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Image deleted successfully"})
}
