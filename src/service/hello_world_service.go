package service

type HelloWorldService interface {
}

// helloWorldService service layer
type helloWorldService struct {
}

// NewHelloWorldService creates a new HelloWorldService
// Constructs type HelloWorldService, depends on []
func NewHelloWorldService() HelloWorldService {
	return &helloWorldService{}
}

func (h *helloWorldService) ToDo() (err error) {
	return err
}
