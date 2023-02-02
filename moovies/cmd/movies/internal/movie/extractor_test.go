package movie_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.tools.sap/distribution-store/moovies/cmd/movies/internal/movie"
	"github.tools.sap/distribution-store/moovies/cmd/movies/internal/movie/mocks"
)

var _ = Describe("Extractor", func() {
	const rate = 4

	var gomockCtrl *gomock.Controller
	var gomockCtx context.Context
	var extractor *Extractor
	var movieDAO *mocks.MockMovieDAO

	BeforeEach(func() {
		gomockCtrl, gomockCtx = gomock.WithContext(context.Background(), GinkgoT())
		movieDAO = mocks.NewMockMovieDAO(gomockCtrl)
		extractor = NewExtractor(movieDAO)
	})

	AfterEach(func() {
		gomockCtrl.Finish()
	})

	Describe("Delete", func() {

		It("deletes in db successfully", func() {
			movieDAO.EXPECT().Delete(gomockCtx, rate).Return(nil)
			err := extractor.Delete(gomockCtx, rate)

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns an error when delete in db fails", func() {
			movieDAO.EXPECT().Delete(gomockCtx, rate).Return(errors.New("db error"))
			err := extractor.Delete(gomockCtx, rate)

			Expect(err).To(HaveOccurred())
		})
	})
})
