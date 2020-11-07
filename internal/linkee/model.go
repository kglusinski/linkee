package linkee

type Page struct {
	Id    string `db:"id" json:"id"`
	Slug string `db:"slug" json:"slug"`
	UserID string `db:"user_id" json:"-"`
	Links []Link `json:"links"`
}

type Link struct {
	ID          string `db:"id" json:"id"`
	Slug string `db:"slug" json:"slug"`
	URL         string `db:"url" json:"url"`
	Description string `db:"description" json:"description"`
	Title       string `db:"title" json:"title"`
	Count       int    `db:"counter" json:"count"`
	PageID		string `db:"page_id" json:"-"`
}
