package data

import (
	"github.com/Killian264/YTLocker/golocker/models"
	"gorm.io/gorm"
)

var (
	ErrorNotFound = gorm.ErrRecordNotFound
)

func (d *Data) GetThumbnails(ID uint64, ownerType string) ([]models.Thumbnail, error) {
	thumbnails := []models.Thumbnail{}

	result := d.db.Where(models.Thumbnail{OwnerID: ID, OwnerType: ownerType}).Find(&thumbnails)

	if removeNotFound(result.Error) != nil {
		return nil, result.Error
	}

	return thumbnails, nil;
}