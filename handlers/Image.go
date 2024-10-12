package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api/logger"
	. "go-api/models"
	"go-api/services"
	"go-api/utils"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type (
	ImageHandler interface {
		GetImage(c echo.Context) error
		UploadImage(c echo.Context) error
		DeleteImageById(c echo.Context) error
	}

	imageHandler struct {
		services.ImageService
	}
)

const imagePath = "./uploads/"

// GetImage
// @Summary Получить изображение по ID.
// @Description Возвращает изображение по ID.
// @Tags Image
// @Produce application/octet-stream
// @Param Authorization header string true "'Bearer _YOUR_TOKEN_'"
// @Param id path int true "Image ID"
// @Security Bearer Authentication
// @Success 200 {file} file "Image received"
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Router /api/v1/image/{id} [get]
func (h *imageHandler) GetImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: "ID is invalid"})
	}

	r, err := h.ImageService.GetImageById(id)
	if err != nil {
		logger.Error("image not found", zap.Error(err))
		return c.JSON(http.StatusNotFound, utils.Error{Message: err.Error()})
	}

	path := fmt.Sprintf("%s%s", imagePath, r.Filename)
	return c.File(path)
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
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/image/ [post]
func (h *imageHandler) UploadImage(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		logger.Error("failed to get image", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	mimeType := file.Header.Get("Content-Type")
	if mimeType != "image/png" && mimeType != "image/jpeg" {
		logger.Error("wrong file type", zap.Error(err))
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "Only image files (png, jpeg) are allowed",
		})
	}

	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(imagePath, filename)

	src, err := file.Open()
	if err != nil {
		logger.Error("Unable to open file", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		logger.Error("Unable to save file", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		logger.Error("Unable to copy file", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	image := Image{
		Filename: filename,
	}

	r, err := h.ImageService.CreateImage(&image)
	if err != nil {
		logger.Error("failed to create image", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, r)
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
// @Failure 400 {object} utils.Error
// @Failure 404 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /api/v1/image/{id} [delete]
func (h *imageHandler) DeleteImageById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		logger.Error("failed to parse id", zap.Error(err))
		return c.JSON(http.StatusBadRequest, utils.Error{Message: err.Error()})
	}

	image, err := h.ImageService.GetImageById(id)
	if err != nil {
		logger.Error("image not found", zap.Error(err))
		return c.JSON(http.StatusNotFound, utils.Error{Message: err.Error()})
	}

	filePath := filepath.Join(imagePath, image.Filename)
	if err = os.Remove(filePath); err != nil {
		logger.Error("failed to delete image", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	if err = h.ImageService.DeleteImage(id); err != nil {
		logger.Error("failed to delete image", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, utils.Error{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Image deleted successfully"})
}
