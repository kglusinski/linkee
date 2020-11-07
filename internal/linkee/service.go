package linkee

type Repository interface {
	FindPage(id string) Page
	UpdateCounter(pageSlug, linkSlug string)
	CreatePage()
	AddLink()
}

type Service struct {
	r Repository
}

func NewService(r Repository) Service {
	return Service{
		r: r,
	}
}

func (s *Service) GetPage(pageId string) Page {
	return s.r.FindPage(pageId)
}

func (s *Service) UpdateCounter(pageSlug, linkSlug string) {
	s.r.UpdateCounter(pageSlug, linkSlug)
}

func (s *Service) CreatePage(page CreatePage) {
	s.r.CreatePage()
}

func (s *Service) AddLink(link AddLink) {
	s.r.AddLink()
}
