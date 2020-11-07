package linkee

type Page struct {
	Id string
	Links []Link
}

type Link struct {
	ID string
	URL string
	Description string
	Count int
}

type Repository interface {
	FindPage(id string) Page
	UpdateCounter(pageID, linkID string)
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

func (s *Service) UpdateCounter(PageID, LinkID string) {
	s.r.UpdateCounter(PageID, LinkID)
}
