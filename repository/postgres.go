package repository

import (
	"database/sql"
	"mini_search_engine/model"
	//"mini_search_engine/model"
)

type PostgresRepository struct{
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
    return &PostgresRepository{db: db}
}

func (r *PostgresRepository)InsertPage(page *model.Page)error{
	query := `
        INSERT INTO mypages.pages (url, title, content, crawled_at)
        VALUES ($1, $2, $3, NOW())
        ON CONFLICT (url) DO UPDATE
        SET title = EXCLUDED.title, content = EXCLUDED.content, crawled_at = NOW();
    `
    _, err := r.db.Exec(query, page.URL, page.Title, page.Content)
    return err
}

// repository/page_repository.go
// func (r *PostgresRepository) SearchPages(query string, limit int) ([]model.SearchResult, error) {
//     rows, err := r.db.Query(`
//         SELECT url, title 
//         FROM mypages.pages 
//         WHERE content_tsv @@ plainto_tsquery('english', $1)
//         ORDER BY ts_rank_cd(content_tsv, plainto_tsquery('english', $1)) DESC
//         LIMIT $2`, query, limit)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var pages []model.SearchResult
//     for rows.Next() {
//         var p model.SearchResult
//         if err := rows.Scan(&p.URL, &p.Title); err != nil {
//             return nil, err
//         }
//         pages = append(pages, pages)
//     }
//     return pages, nil
// }

// func (r *PostgresRepository) SearchPages(query string, limit int) ([]model.SearchResult, error) {
//     rows, err := r.db.Query(`
//         SELECT url, title
//         FROM mypages.pages
//         WHERE content_tsv @@ plainto_tsquery('english', $1)
//         ORDER BY ts_rank_cd(content_tsv, plainto_tsquery('english', $1)) DESC
//         LIMIT $2
//     `, query, limit)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var results []model.SearchResult
//     for rows.Next() {
//         var res model.SearchResult
//         if err := rows.Scan(&res.URL, &res.Title); err != nil {
//             return nil, err
//         }
//         results = append(results, res)
//     }
//     return results, nil
// }

func (r *PostgresRepository) SearchPages(query string, limit int) ([]model.SearchResult, error) {
    rows, err := r.db.Query(`
        SELECT url, title
        FROM mypages.pages
        WHERE content_tsv @@ plainto_tsquery('english', $1)
        ORDER BY ts_rank_cd(content_tsv, plainto_tsquery('english', $1)) DESC
        LIMIT $2
    `, query, limit)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var results []model.SearchResult
    for rows.Next() {
        var res model.SearchResult
        if err := rows.Scan(&res.URL, &res.Title); err != nil {
            return nil, err
        }
        results = append(results, res)
    }
    return results, nil
}