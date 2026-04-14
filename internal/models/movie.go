package models

type Movie struct {
	ID	  		int
	Title 		string
	Slug  		string
	Image 		string
	Description string
	Year		int
	Director 	string
	Actors		[]string
	Genres      []string
	File 		string
}

type MovieListPage struct {
	Movies []Movie
	Title string
}

