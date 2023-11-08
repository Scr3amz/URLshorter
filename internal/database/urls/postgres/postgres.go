package urls

import (
	"context"
	"database/sql"

	"github.com/Scr3amz/URLshorter/internal/database/urls"
	"github.com/jmoiron/sqlx"
)

type Client interface {
	// NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	// QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	// BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error)
	Exec(query string, args ...any) (sql.Result, error)
}

type repository struct {
	client Client
}

func (r *repository) Create(ctx context.Context, u *urls.URLs) error {
	q := `
	INSERT INTO urls
 	(longurl, shorturl) 
 	VALUES 
 	($1, $2)
	`

	if _, err := r.client.Exec(q, u.LongURL, u.ShortURL); err != nil {
		return err
	}

	return nil
}

func (r *repository) FindLong(ctx context.Context,  u urls.URLs) (urls.URLs, error) {
	q := `
	SELECT shorturl 
	FROM urls
 	WHERE longurl =  $1
	`
	var url urls.URLs
	if err := r.client.QueryRowxContext(ctx, q, u.LongURL).Scan(&url.ShortURL); err != nil {
		return urls.URLs{}, err
	}
	return url, nil
}

func (r *repository) FindShort(ctx context.Context,  u urls.URLs) (urls.URLs, error) {
	q := `
	SELECT longurl 
	FROM urls
 	WHERE shorturl =  $1
	`
	var url urls.URLs
	if err := r.client.QueryRowxContext(ctx, q, u.ShortURL).Scan(&url.LongURL); err != nil {
		return urls.URLs{}, err
	}
	return url, nil
}

func NewRepository(client Client) urls.Repository {
	return &repository{
		client: client,
	}
}
