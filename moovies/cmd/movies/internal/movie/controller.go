package movie

import (
	"context"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.tools.sap/distribution-store/moovies/pkg/api"
	"github.tools.sap/distribution-store/moovies/pkg/database"
)

//go:generate mockgen --source=controller.go --destination mocks/mock_controller.go --package mocks

type MovieDAO interface {
	Insert(ctx context.Context, movie database.Movie) error
	Read(ctx context.Context, title string) (database.Movie, bool, error)
	Delete(ctx context.Context, rate int) error
}

type UUIDGenerator interface {
	Generate() uuid.UUID
}

func NewController(movieDAO MovieDAO, generator UUIDGenerator) *Controller {
	return &Controller{
		generator: generator,
		movieDAO:  movieDAO,
	}
}

type Controller struct {
	generator UUIDGenerator
	movieDAO  MovieDAO
}

func (c *Controller) Insert(ctx context.Context, movie api.Movie) error {
	parsedYear, err := strconv.Atoi(movie.Year)
	if err != nil {
		return err
	}
	parsedRate, err := strconv.Atoi(movie.Rate)
	if err != nil {
		return err
	}

	return c.movieDAO.Insert(ctx, database.Movie{
		ID:    c.generator.Generate(),
		Title: movie.Title,
		Year:  parsedYear,
		Rate:  parsedRate,
	})
}

func (c *Controller) Read(ctx context.Context, title string) (api.Movie, bool, error) {
	movie, found, err := c.movieDAO.Read(ctx, title)

	if err != nil {
		return api.Movie{}, false, err
	}
	if !found {
		return api.Movie{}, false, nil
	}

	apiMovie := api.Movie{
		Title: title,
		Year:  strconv.Itoa(movie.Year),
		Rate:  strconv.Itoa(movie.Rate),
	}
	return apiMovie, true, nil
}
