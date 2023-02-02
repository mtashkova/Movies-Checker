package movie_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	. "github.tools.sap/distribution-store/moovies/cmd/movies/internal/movie"
	"github.tools.sap/distribution-store/moovies/cmd/movies/internal/movie/mocks"
	"github.tools.sap/distribution-store/moovies/pkg/api"
	"github.tools.sap/distribution-store/moovies/pkg/database"
)

var _ = Describe("Controller", func() {
	const title = "title"

	var gomockCtrl *gomock.Controller
	var gomockCtx context.Context
	var controller *Controller
	var movieDAO *mocks.MockMovieDAO
	var generator *mocks.MockUUIDGenerator

	BeforeEach(func() {
		gomockCtrl, gomockCtx = gomock.WithContext(context.Background(), GinkgoT())
		generator = mocks.NewMockUUIDGenerator(gomockCtrl)
		movieDAO = mocks.NewMockMovieDAO(gomockCtrl)
		controller = NewController(movieDAO, generator)
	})

	AfterEach(func() {
		gomockCtrl.Finish()
	})

	Describe("Insert", func() {
		var apiMovie api.Movie

		It("inserts in db successfully", func() {
			id := uuid.NewV4()
			movie := database.Movie{
				ID:    id,
				Title: title,
				Year:  2020,
				Rate:  9,
			}
			apiMovie = api.Movie{
				Title: title,
				Year:  "2020",
				Rate:  "9",
			}

			generator.EXPECT().Generate().Return(id)
			movieDAO.EXPECT().Insert(gomockCtx, movie).Return(nil)
			Expect(controller.Insert(gomockCtx, apiMovie)).To(Succeed())
		})

		It("returns error when year is not a number", func() {
			apiMovie = api.Movie{
				Title: title,
				Year:  "2k20",
				Rate:  "9",
			}

			Expect(controller.Insert(gomockCtx, apiMovie)).ToNot(Succeed())
		})

		It("returns error when rate is not a number", func() {
			apiMovie = api.Movie{
				Title: title,
				Year:  "2020",
				Rate:  "nine",
			}
			Expect(controller.Insert(gomockCtx, apiMovie)).ToNot(Succeed())
		})

		It("returns error when insert in db fails", func() {
			id := uuid.NewV4()
			movie := database.Movie{
				ID:    id,
				Title: title,
				Year:  2020,
				Rate:  9,
			}
			apiMovie = api.Movie{
				Title: title,
				Year:  "2020",
				Rate:  "9",
			}
			generator.EXPECT().Generate().Return(id)
			movieDAO.EXPECT().Insert(gomockCtx, movie).Return(errors.New("db error"))
			Expect(controller.Insert(gomockCtx, apiMovie)).ToNot(Succeed())
		})
	})

	Describe("Read", func() {
		It("reads successfully from db", func() {
			dbMovie := database.Movie{
				Title: title,
				Year:  2020,
				Rate:  8,
			}
			expectedAPIMovie := api.Movie{
				Title: title,
				Year:  "2020",
				Rate:  "8",
			}

			movieDAO.EXPECT().Read(gomockCtx, title).Return(dbMovie, true, nil)
			apiMovie, found, err := controller.Read(gomockCtx, title)

			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeTrue())
			Expect(apiMovie).To(Equal(expectedAPIMovie))
		})

		It("returns error when read in db fails", func() {
			movieDAO.EXPECT().Read(gomockCtx, title).Return(database.Movie{}, false, errors.New("db error"))
			_, _, err := controller.Read(gomockCtx, title)

			Expect(err).To(HaveOccurred())
		})

		It("indicates that there is no such movie in db", func() {
			movieDAO.EXPECT().Read(gomockCtx, title).Return(database.Movie{}, false, nil)
			_, found, err := controller.Read(gomockCtx, title)

			Expect(err).ToNot(HaveOccurred())
			Expect(found).To(BeFalse())
		})
	})
})
