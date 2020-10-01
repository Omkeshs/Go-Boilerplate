package services

import (
	"context"
	"encoding/json"
	"net/http"
	"practice/whetherinfo/workers/models"
	repository "practice/whetherinfo/workers/repositories"
	"sync"
)

// Service describes the service.
type Service interface {
	CreateCity(context.Context, *sync.WaitGroup) error
	GetCity(context.Context, string) (*models.Cities, error)
}

//ServiceImpl **
type ServiceImpl struct {
	repo repository.Repository
}

//NewServiceImpl inject depedancies user repositiory
func NewServiceImpl(repo repository.Repository) Service {
	return &ServiceImpl{repo: repo}
}

const (
	APPID = "869d6e65fd1a807d3a53ff1ad6b5a458"
)

func (serviceImpl ServiceImpl) CreateCity(ctx context.Context, wg *sync.WaitGroup) error {

	wg.Add(10)

	var routine = len(Cities)
	citiesCh := make(chan string)

	for i := 0; i < routine; i++ {
		go func(cityCh <-chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			for {
				cityname, ok := <-cityCh
				if !ok {
					return
				}

				response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + cityname + "&appid=" + APPID)

				if err != nil {
					return
				}
				cityInfo := models.CityInfo{}
				decoder := json.NewDecoder(response.Body)
				err = decoder.Decode(&cityInfo)

				if err != nil {
					return
				}

				if err != nil {
					return
				}

				b, err := json.Marshal(cityInfo)

				if err != nil {
					return
				}

				city := models.Cities{}
				city.Name = cityname
				city.Lat = cityInfo.Coord.Lat
				city.Lon = cityInfo.Coord.Lon
				city.Temp = cityInfo.Main.Temp
				city.Pressure = cityInfo.Main.Pressure
				city.SeaLevel = cityInfo.Main.SeaLevel
				city.MoreInfo = string(b)

				err = serviceImpl.repo.CreateCity(ctx, city)

				if err != nil {
					return
				}
			}
		}(citiesCh, wg)
	}

	for _, citi := range Cities {
		citiesCh <- citi
	}

	wg.Wait()

	close(citiesCh)
	return nil
}

func (serviceImpl ServiceImpl) GetCity(ctx context.Context, cityName string) (resp *models.Cities, err error) {
	resp, err = serviceImpl.repo.GetCity(ctx, cityName)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
