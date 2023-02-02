package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.tools.sap/distribution-store/moovies/cmd/movies/env"
	"github.tools.sap/distribution-store/moovies/cmd/movies/internal/movie"
	"github.tools.sap/distribution-store/moovies/pkg/api"
	"github.tools.sap/distribution-store/moovies/pkg/connection"
	"github.tools.sap/distribution-store/moovies/pkg/database"
	"github.tools.sap/distribution-store/moovies/pkg/uuid"
)

// 1. Reimplement basic auth in way it could be configured
// 2. The same as host and port

func main() {
	config, err := connection.LoadConfig()
	if err != nil {
		panic(err)
	}

	db, err := connection.ConnectToDB(config)
	if err != nil {
		panic(err)
	}
	generator := uuid.NewUUIDGenerator()
	movieDAO, err := database.NewMovieDAO(context.Background(), db)
	if err != nil {
		panic("Unable to init db")
	}
	movieController := movie.NewController(movieDAO, generator)
	movieExtractor := movie.NewExtractor(movieDAO)

	r := gin.Default()

	basicAuthConfig, err := env.BasicAuthConfiguration()
	if err != nil {
		panic(err)
	}
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		basicAuthConfig.Username: basicAuthConfig.Password,
	}))

	authorized.POST("/movie", func(c *gin.Context) {
		var movie api.Movie
		if err := c.Bind(&movie); err != nil {
			c.JSON(http.StatusInternalServerError, struct{}{})
			return
		}
		err = movieController.Insert(c.Request.Context(), movie)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct{}{})
			return
		}
		c.JSON(http.StatusOK, struct{}{})
	})

	authorized.GET("/movie/:title", func(c *gin.Context) {
		title := c.Param("title")
		movie, found, err := movieController.Read(c.Request.Context(), title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct{}{})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusOK, movie)
	})

	authorized.DELETE("/movie/:rate", func(c *gin.Context) {
		rate := c.Param("rate")
		movie, found, err := movieExtractor.Delete(c.Request.Context(), rate)
		if err != nil {
			c.JSON(http.StatusInternalServerError, struct{}{})
			return
		}
		if !found {
			c.JSON(http.StatusNotFound, struct{}{})
			return
		}
		c.JSON(http.StatusOK, movie)
	})

	srv := &http.Server{
		Addr:    ":8050",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		fmt.Println("timeout of 5 seconds.")
	}

	fmt.Println("Successfully stopped!")
}
