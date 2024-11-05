package cleanup

import (
	"fmt"
	"go-server/bootstrap"
	"go-server/domain"
	"log"
	"os"
	"path/filepath"
	"time"
)

func cleanOldFiles(app *bootstrap.Application) error {
	thresholdTime := time.Now().Add(-time.Duration(app.Env.ThresholdHours) * time.Hour)
	var oldFiles []domain.Image

	if err := app.Database.Where("created_at < ?", thresholdTime).Find(&oldFiles).Error; err != nil {
		return fmt.Errorf("failed to query old files: %v", err)
	}

	for _, file := range oldFiles {
		filePath := filepath.Join(app.Env.StoragePath, file.Filename)
		if err := os.Remove(filePath); err != nil {
			fmt.Printf("Не удалось удалить файл %s: %v\n", filePath, err)
		} else {
			fmt.Printf("Файл %s удалён\n", filePath)
		}

		if err := app.Database.Delete(&file).Error; err != nil {
			fmt.Printf("Не удалось удалить запись о файле %s из БД: %v\n", file.Filename, err)
		}
	}

	fmt.Println("Очистка завершена")
	return nil
}

func CleanOldFiles(app *bootstrap.Application) {
	interval := time.Duration(app.Env.IntervalHours) * time.Hour
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		fmt.Println("Запуск фоновой задачи очистки файлов...")
		if err := cleanOldFiles(app); err != nil {
			log.Printf("Ошибка во время очистки файлов: %v", err)
		} else {
			fmt.Println("Очистка файлов завершена успешно.")
		}
	}
}
