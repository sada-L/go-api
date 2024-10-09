package handlers

/*import (
	"fmt"
	"path/filepath"
	"time"

	"gorm.io/gorm"
)

var DB *gorm.DB

// Путь для хранения изображений
const imagePath = "./uploads/images/"

// Загрузка изображения
func UploadImage(c *fiber.Ctx) error {
	// Получаем файл из запроса
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Не удалось получить изображение",
		})
	}

	// Генерация уникального имени файла
	filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	filePath := filepath.Join(imagePath, filename)

	// Сохраняем файл на диск
	if err := c.SaveFile(file, filePath); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить изображение",
		})
	}

	// Сохраняем метаданные изображения в базе данных
	image := Image{
		Filename:   filename,
		Filepath:   filePath,
		UploadTime: time.Now(),
	}
	if result := DB.Create(&image); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Не удалось сохранить информацию об изображении в базе данных",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(image)
}
*/
