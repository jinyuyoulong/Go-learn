package datasource

// import "github.com/kataras/iris/v12/_examples/mvc/overview/datamodels"

import (
	"github.com/jinyuyoulong/Go-learn/src/iris/mini_demo/MVC/datamodels"
)

// Movies is our imaginary data source.
var Movies = map[int64]datamodels.Movie{
	1: {
		ID:     1,
		Name:   "Caseblance",
		Year:   1942,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/1.jpg",
	},
	2: {
		ID:     2,
		Name:   "Gone with the Wind",
		Year:   1939,
		Genre:  "Romance",
		Poster: "https://iris-go.com/images/examples/mvc-movies/2.jpg",
	},
	3: {
		ID:     3,
		Name:   "Citizen Kane",
		Year:   1941,
		Genre:  "Mystery",
		Poster: "https://iris-go.com/images/examples/mvc-movies/3.jpg",
	},
	4: {
		ID:     4,
		Name:   "The Wizard of Oz",
		Year:   1939,
		Genre:  "Fantasy",
		Poster: "https://iris-go.com/images/examples/mvc-movies/4.jpg",
	},
	5: {
		ID:     5,
		Name:   "North by Northwest",
		Year:   1959,
		Genre:  "Thriller",
		Poster: "https://iris-go.com/images/examples/mvc-movies/5.jpg",
	},
}
