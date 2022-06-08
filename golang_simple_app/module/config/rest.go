package config

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"simpleapp/module/service"
)

// RestConfig will hold all configuration that we need to run this server
type RestConfig struct {
	DatabaseConfig DatabaseConfig `envconfig:"DB"`
}

// NewRestServer create REST server with registered router
func NewRestServer(rc RestConfig) (http.Handler, error) {
	router := httprouter.New()

	sqlxDB, err := sqlx.Connect(rc.DatabaseConfig.Driver, rc.DatabaseConfig.RWDataSourceName())
	if err != nil {
		return nil, err
	}

	// Register all endpoint related to book service here.
	bookService := service.NewBookService(newBookUsecase(sqlxDB))
	router.PUT("/books", bookService.Update)
	router.POST("/books", bookService.Create)
	router.GET("/books/:id", bookService.Get)
	router.DELETE("/books/:id", bookService.Delete)

	router.HandlerFunc("GET", "/metrics", promhttp.Handler().ServeHTTP)

	return router, nil
}
