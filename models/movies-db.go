package models

import (
	"context"
	"database/sql"
	"time"
)

type DBModel struct {
	DB *sql.DB
}

func (m *DBModel) GetMovie(id int) (*Movie, error) {
	ctx, cancle := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancle()
	query := `select id, title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at from movies WHERE id = $1`
	row := m.DB.QueryRowContext(ctx, query, id)

	var movie Movie

	err := row.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.Year, &movie.ReleaseDate, &movie.Runtime, &movie.Rating, &movie.MPAARating, &movie.CreatedAt, &movie.UpdatedAt)
	if err != nil {
		return nil, err
	}

	// Get the genres for the movie
	query = `select 
				mg.id, mg.movie_id, mg.genre_id, g.genre_name
			from
				movie_genre mg
				left join genres g on (mg.genre_id = g.id)
			where 
				mg.movie_id = $1`
	rows, _ := m.DB.QueryContext(ctx, query, id)
	defer rows.Close()

	genres := make(map[int]string)
	for rows.Next() {
		var mg MovieGenre
		err := rows.Scan(
			&mg.ID,
			&mg.MovieID,
			&mg.GenreID,
			&mg.Genre.GenreName,
		)
		if err != nil {
			return nil, err
		}
		genres[mg.ID] = mg.Genre.GenreName
	}

	movie.MovieGenre = genres

	return &movie, nil
}

func (m *DBModel) All(id int) ([]*Movie, error) {
	return nil, nil
}
