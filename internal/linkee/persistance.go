package linkee

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

type MySQLRepository struct {
	db *sqlx.DB
}

func NewMySQLRepository(db *sqlx.DB) *MySQLRepository {
	return &MySQLRepository{
		db: db,
	}
}

func (r *MySQLRepository) FindPage(slug string) Page {
	var page Page
	var links []Link

	sql := `SELECT *
			FROM pages
			WHERE slug = ?`
	sqlPages := `SELECT * FROM links WHERE page_id = ?`

	err := r.db.Get(&page, sql, slug)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't get page")
	}

	err = r.db.Select(&links, sqlPages, page.Id)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't get links")
	}

	page.Links = links

	return page
}

func (r *MySQLRepository) UpdateCounter(pageSlug, linkSlug string) {
	sql := `UPDATE links SET counter = counter + 1 WHERE slug = ? AND page_id = ?`
	pageSql := `SELECT id FROM pages WHERE slug = ?`

	var pageId string

	err := r.db.Get(&pageId, pageSql, pageSlug)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't update link counter")
	}

	_, err = r.db.Exec(sql, linkSlug, pageId)
	if err != nil {
		log.Error().Err(err).Msg("Couldn't update link counter")
	}
}

func (r *MySQLRepository) CreatePage() {}
func (r *MySQLRepository) AddLink() {}
