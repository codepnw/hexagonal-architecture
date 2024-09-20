package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/codepnw/hexagonal/handler"
	"github.com/codepnw/hexagonal/logs"
	"github.com/codepnw/hexagonal/repository"
	"github.com/codepnw/hexagonal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimezone()
	initConfig()
	db := initDatabase()

	custRepoDB := repository.NewCustomerRepositoryDB(db)
	custRepoMock := repository.NewCustomerRepositoryMock()
	_ = custRepoMock
	custService := service.NewCustomerService(custRepoDB)
	custHandler := handler.NewCustomerHandler(custService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", custHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", custHandler.GetCustomer).Methods(http.MethodGet)

	port := fmt.Sprintf(":%d", viper.GetInt("app.port"))

	logs.Info("server starting at " + port)
	http.ListenAndServe(port, router)
}

func initConfig() {
	viper.SetConfigName("example_config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()
	// example: APP_PORT=5000 go run .
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func initTimezone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(2 * time.Minute)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
