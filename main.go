package main

import (
	"net/http"
	"os"

	"github.com/codepnw/hexagonal/handler"
	"github.com/codepnw/hexagonal/repository"
	"github.com/codepnw/hexagonal/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db, err := sqlx.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}

	custRepoDB := repository.NewCustomerRepositoryDB(db)
	_ = custRepoDB
	custRepoMock := repository.NewCustomerRepositoryMock()
	custService := service.NewCustomerService(custRepoMock)
	custHandler := handler.NewCustomerHandler(custService)

	router := mux.NewRouter()

	router.HandleFunc("/customers", custHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", custHandler.GetCustomer).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)
}
