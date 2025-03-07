package service

// MessageService provides a method to return a message.
type ChartService struct{}

// GetMessage returns a simple string message.
func (s *ChartService) GetMessage() string {
    return "Hello from ChartService!"
}
