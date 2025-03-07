package service

import (
	"os"

	"github.com/spf13/viper"
)

// MessageService provides a method to return a message.
type ChartService struct{}

// GetMessage returns a simple string message.
func (s *ChartService) GetMessage() string {
    message := os.Getenv("DATABASE_URL")
    db_url := viper.GetString("DATABASE_URL")
    return "Hello from ChartService! message is: " + message + "db-url: " + db_url
}
