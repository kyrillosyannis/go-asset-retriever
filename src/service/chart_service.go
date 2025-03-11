package service

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.challenge/models/domain"
	"go.challenge/repository"
)

// MessageService provides a method to return a message.
type ChartService struct{
    repo *repository.ChartRepository
}

func NewChartService(repo *repository.ChartRepository) *ChartService {
	return &ChartService{repo: repo}
}

// GetMessage returns a simple string message.
func (s *ChartService) GetMessage() string {
    message := os.Getenv("DATABASE_URL")
    db_url := viper.GetString("DATABASE_URL")
    return "Hello from ChartService! message is: " + message + "db-url: " + db_url
}

func (s *ChartService) GetAllCharts() ([]domain.Chart, error) {
	items, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

    fmt.Println(items)

	var itemDTOs []domain.Chart
	for _, item := range items {
		itemDTOs = append(itemDTOs, item)
	}
	return itemDTOs, nil
}
