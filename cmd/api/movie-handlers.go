package main

import (
	"backend/models"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("Fetching movie with ID:", id)

	movie := models.Movie{
		ID:          id,
		Title:       "Inception",
		Description: "A mind-bending thriller",
		Year:        2010,
		ReleaseDate: time.Date(2010, 7, 16, 0, 0, 0, 0, time.Local),
		Runtime:     148,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {

}
