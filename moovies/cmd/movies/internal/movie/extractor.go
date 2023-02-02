package movie

import "context"

// type MovieDAO interface {
// 	Delete(ctx context.Context, rate int) error
// }

type Extractor struct {
	movieDAO MovieDAO
}

func NewExtractor(movieDAO MovieDAO) *Extractor {
	return &Extractor{
		movieDAO: movieDAO,
	}
}

func (e *Extractor) Delete(ctx context.Context, rate int) error {
	err := e.movieDAO.Delete(ctx, rate)
	if err != nil {
		return err
	}

	return nil
}
