package templates

type TemplateService struct {
}

func NewService() *TemplateService {
	return &TemplateService{}
}

func (s *TemplateService) Get() error {
	return nil
}
