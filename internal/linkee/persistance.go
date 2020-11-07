package linkee

var inMemoryDatabase = map[string]Page{
	"lorem-ipsum": {
		Id:    "lorem-ipsum",
		Links: []Link{{
				ID:          "zaprogramowani",
				URL:         "https://zaprogramowani.dev",
				Description: "Najlepszy blog o programowaniu",
				Count:       0,
			},
		},
	},
}

type InMemoryRepository struct {
	db map[string]Page
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{
		db: inMemoryDatabase,
	}
}

func (r *InMemoryRepository)FindPage(id string) Page {
	return inMemoryDatabase[id]
}

func (r *InMemoryRepository) UpdateCounter(pageID, linkID string) {
	for _, link := range inMemoryDatabase[pageID].Links {
		if link.ID == linkID {
			link.Count = link.Count + 1
		}
	}
}
