package postgres

import (
	"database/sql"
	"kino-site/internal/models"
)

type MovieStorage struct {
	DB *sql.DB
}

func (s *MovieStorage) GetAll() ([]models.Movie, error) {
	rows, err := s.DB.Query(`
	SELECT 
		id, 
		title, 
		slug, 
		image, 
		description, 
		year, 
		director 
	FROM 
		movies
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID, 
			&m.Title, 
			&m.Slug, 
			&m.Image,
			&m.Description,
			&m.Year,
			&m.Director,
		)
		if err != nil {
			return nil, err
		}
		
		movies = append(movies, m)
	}

	return movies, nil
}

func (s *MovieStorage) GetBySlug(slug string) (*models.Movie, error) {
	var m models.Movie
	var file sql.NullString

	err := s.DB.QueryRow(`
	SELECT 
		id, 
		title, 
		slug, 
		image,
		description,
		year,
		director,
		file
	FROM 
		movies 
	WHERE 
		slug=$1`, 
	slug).Scan(
		&m.ID, 
		&m.Title, 
		&m.Slug, 
		&m.Image,
		&m.Description,
		&m.Year,
		&m.Director,
		&file,
	)
	if err != nil {
		return nil, err
	}

	if file.Valid {
		m.File = file.String
	} else {
		m.File = "" 
	}

	actorRows, err := s.DB.Query(`
	SELECT 
		a.name
	FROM
		actors a
	JOIN
		movie_actors ma
	ON
		a.id = ma.actor_id
	WHERE
		ma.movie_id = $1
	`, m.ID)
	if err != nil {
		return nil, err
	}
	defer actorRows.Close()

	for actorRows.Next() {
		var name string
		if err := actorRows.Scan(&name); err != nil {
			return nil, err
		}
		m.Actors = append(m.Actors, name)
	}

	genreRows, err := s.DB.Query(`
	SELECT 
		g.name
	FROM 
		genres g
	JOIN
		movie_genres mg
	ON
		g.id = mg.genre_id
	WHERE
		mg.movie_id = $1
	`, m.ID)
	if err != nil {
		return nil, err
	}
	defer genreRows.Close()

	for genreRows.Next() {
		var name string
		if err := genreRows.Scan(&name); err != nil {
			return nil, err
		}
		m.Genres = append(m.Genres, name)
	}

	return &m, nil
}

func (s *MovieStorage) Search(query string) ([]models.Movie, error) {
	rows, err := s.DB.Query(`
	SELECT
		id, 
		title,
		slug,
		image,
		description, 
		year, 
		director
	FROM
		movies
	WHERE
		LOWER(title) LIKE LOWER($1)
	`, "%" + query + "%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
            &m.Title,
            &m.Slug,
            &m.Image,
            &m.Description,
            &m.Year,
            &m.Director,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (s *MovieStorage) GetByGenre(genre string) ([]models.Movie, error) {
	rows, err := s.DB.Query(`
	SELECT 
		m.id, m.title, m.slug, m.image, m.description, m.year, m.director
	FROM
		movies m
	JOIN
		movie_genres mg
	ON
		m.id = mg.movie_id
	JOIN
		genres g
	ON
		g.id = mg.genre_id
	WHERE
		LOWER(g.name) = LOWER($1)
	`, genre)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Slug,
			&m.Image,
			&m.Description,
			&m.Year,
			&m.Director,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}

func (s *MovieStorage) GetByDirector(director string) ([]models.Movie, error) {
	rows, err := s.DB.Query(`
	SELECT 
		id, title, slug, image, description, year, director
	FROM 
		movies
	WHERE 
		LOWER(director) = LOWER($1)
	`, director)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Slug,
			&m.Image,
			&m.Description,
			&m.Year,
			&m.Director,
		)
		if err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}

	return movies, nil
}