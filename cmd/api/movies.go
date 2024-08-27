package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight/internal/data"
	"greenlight/internal/validator"
)

func (app *application) createMovie(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	if err := readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()

	if data.ValidateMovie(v, movie); !v.IsValid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	_, _ = fmt.Fprintf(w, "%v\n", input)
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:        id,
		Title:     "Casablanca",
		Year:      1998,
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		CreatedAt: time.Now(),
		Version:   1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
