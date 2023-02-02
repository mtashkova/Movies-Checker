package database

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type Movie struct {
	ID    uuid.UUID
	Title string
	Year  int
	Rate  int
}

func NewMovieDAO(ctx context.Context, db *sql.DB) (*MovieDAO, error) {
	_, err := db.ExecContext(ctx, createMovieTableSQL)
	if err != nil {
		return nil, err
	}
	insertStmt, err := db.PrepareContext(ctx, insertMovieSQL)
	if err != nil {
		return nil, err
	}
	readByTitleStmt, err := db.PrepareContext(ctx, readMovieByTitleSQL)
	if err != nil {
		return nil, err
	}
	deleteByRateStmt, err := db.PrepareContext(ctx, deleteMovieByRateSQL)
	if err != nil {
		return nil, err
	}

	return &MovieDAO{
		insertStmt:       insertStmt,
		readByTitleStmt:  readByTitleStmt,
		deleteByRateStmt: deleteByRateStmt,
	}, nil
}

type MovieDAO struct {
	insertStmt       *sql.Stmt
	readByTitleStmt  *sql.Stmt
	deleteByRateStmt *sql.Stmt
}

func (d *MovieDAO) Insert(ctx context.Context, movie Movie) error {
	_, err := d.insertStmt.ExecContext(ctx,
		movie.ID,
		movie.Title,
		movie.Year,
		movie.Rate,
	)

	if err != nil {
		return err
	}
	return nil
}

func (d *MovieDAO) Read(ctx context.Context, title string) (Movie, bool, error) {
	rows, err := d.readByTitleStmt.QueryContext(ctx, title)
	if err != nil {
		return Movie{}, false, err
	}
	defer rows.Close()

	movies, err := scanMovies(rows)
	if err != nil {
		return Movie{}, false, err
	}

	switch len(movies) {
	case 0:
		return Movie{}, false, nil
	default:
		return movies[0], true, nil
	}

}
func (d *MovieDAO) Delete(ctx context.Context, rate int) error {
	_, err := d.deleteByRateStmt.QueryContext(ctx, rate)

	if err != nil {
		return err
	}
	return nil
}

func scanMovies(rows *sql.Rows) ([]Movie, error) {

	movies := []Movie{}
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Year,
			&movie.Rate,
		); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}
	return movies, nil
}
