package repositories

import (
	"context"
	"practice/whetherinfo/workers/models"

	"github.com/jinzhu/gorm"
)

//Repository implimets all methods in Repository
type Repository interface {
	CreateCity(context.Context, models.Cities) error
	GetCity(context.Context, string) (*models.Cities, error)
}

//RepositoryImpl **
type RepositoryImpl struct {
	dbConn *gorm.DB
}

//NewRepositoryImpl inject dependancies of DataStore
func NewRepositoryImpl(dbConn *gorm.DB) Repository {
	return &RepositoryImpl{dbConn: dbConn}
}

func (repositoryImpl RepositoryImpl) CreateCity(ctx context.Context, city models.Cities) (err error) {
	dbConn := repositoryImpl.dbConn
	err = dbConn.Table("cities").Create(&city).Error
	if err != nil {
		return
	}
	return nil
}

func (repositoryImpl RepositoryImpl) GetCity(ctx context.Context, cityName string) (*models.Cities, error) {
	dbConn := repositoryImpl.dbConn
	city := models.Cities{}
	err := dbConn.Where("name=?", cityName).First(&city).Error
	if err != nil {
		return nil, err
	}
	return &city, nil
}
